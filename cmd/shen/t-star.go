package main

import . "github.com/tiancaiamao/shen-go/kl"

var TStarMain = MakeNative(func(__e *ControlFlow) {
tmp12362 := MakeNative(func(__e *ControlFlow) {
V4853 := __e.Get(1)
_ = V4853
V4854 := __e.Get(2)
_ = V4854
tmp12363 := MakeNative(func(__e *ControlFlow) {
W4855 := __e.Get(1)
_ = W4855
tmp12364 := MakeNative(func(__e *ControlFlow) {
W4856 := __e.Get(1)
_ = W4856
tmp12365 := MakeNative(func(__e *ControlFlow) {
W4857 := __e.Get(1)
_ = W4857
tmp12366 := MakeNative(func(__e *ControlFlow) {
Z4858 := __e.Get(1)
_ = Z4858
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4859 := __e.Get(1)
_ = Z4859
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4860 := __e.Get(1)
_ = Z4860
__e.Return(MakeNative(func(__e *ControlFlow) {
Z4861 := __e.Get(1)
_ = Z4861
tmp12367 := MakeNative(func(__e *ControlFlow) {
W4862 := __e.Get(1)
_ = W4862
tmp12368 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12368

tmp12369 := Call(__e, PrimFunc(symshen_4deref), W4855, Z4858)


tmp12370 := Call(__e, PrimFunc(symreceive), tmp12369)


tmp12371 := Call(__e, PrimFunc(symshen_4deref), W4856, Z4858)


tmp12372 := Call(__e, PrimFunc(symreceive), tmp12371)


tmp12373 := MakeNative(func(__e *ControlFlow) {
tmp12374 := Call(__e, PrimFunc(symshen_4deref), W4857, Z4858)


tmp12375 := Call(__e, PrimFunc(symreceive), tmp12374)


tmp12376 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symreturn), W4862, Z4858, Z4859, Z4860, Z4861)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4toplevel_1forms), tmp12375, W4862, Z4858, Z4859, Z4860, tmp12376)
return


}, 0)

tmp12377 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), tmp12370, tmp12372, W4862, Z4858, Z4859, Z4860, tmp12373)


__e.TailApply(PrimFunc(symshen_4gc), Z4858, tmp12377)
return


}, 1)

tmp12378 := Call(__e, PrimFunc(symshen_4newpv), Z4858)


__e.TailApply(tmp12367, tmp12378)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp12379 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp12380 := Call(__e, tmp12366, tmp12379)


tmp12381 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp12382 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp12381)


tmp12383 := Call(__e, PrimFunc(sym_8v), True, tmp12382)


tmp12384 := Call(__e, tmp12380, tmp12383)


tmp12385 := Call(__e, tmp12384, MakeNumber(0))


tmp12386 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

__e.TailApply(tmp12385, tmp12386)
return


}, 1)

tmp12387 := Call(__e, PrimFunc(symshen_4curry), V4853)


__e.TailApply(tmp12365, tmp12387)
return


}, 1)

tmp12388 := Call(__e, PrimFunc(symshen_4rectify_1type), V4854)


__e.TailApply(tmp12364, tmp12388)
return


}, 1)

tmp12389 := Call(__e, PrimFunc(symshen_4extract_1vars), V4854)


__e.TailApply(tmp12363, tmp12389)
return


}, 2)

tmp12390 := Call(__e, ns2_1set, symshen_4typecheck, tmp12362)


_ = tmp12390

tmp12391 := MakeNative(func(__e *ControlFlow) {
V4863 := __e.Get(1)
_ = V4863
V4864 := __e.Get(2)
_ = V4864
V4865 := __e.Get(3)
_ = V4865
V4866 := __e.Get(4)
_ = V4866
V4867 := __e.Get(5)
_ = V4867
V4868 := __e.Get(6)
_ = V4868
V4869 := __e.Get(7)
_ = V4869
tmp12392 := MakeNative(func(__e *ControlFlow) {
W4870 := __e.Get(1)
_ = W4870
tmp12410 := PrimEqual(W4870, False)

if True == tmp12410 {
tmp12408 := Call(__e, PrimFunc(symshen_4unlocked_2), V4867)


if True == tmp12408 {
tmp12393 := MakeNative(func(__e *ControlFlow) {
W4872 := __e.Get(1)
_ = W4872
tmp12405 := PrimIsPair(W4872)

if True == tmp12405 {
tmp12394 := MakeNative(func(__e *ControlFlow) {
W4873 := __e.Get(1)
_ = W4873
tmp12395 := MakeNative(func(__e *ControlFlow) {
W4874 := __e.Get(1)
_ = W4874
tmp12396 := MakeNative(func(__e *ControlFlow) {
W4875 := __e.Get(1)
_ = W4875
tmp12397 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12397

tmp12398 := Call(__e, PrimFunc(symshen_4deref), W4875, V4866)


tmp12399 := Call(__e, PrimFunc(symsubst), tmp12398, W4873, V4864)


tmp12400 := Call(__e, PrimFunc(symshen_4insert_1prolog_1variables), W4874, tmp12399, V4865, V4866, V4867, V4868, V4869)


__e.TailApply(PrimFunc(symshen_4gc), V4866, tmp12400)
return


}, 1)

tmp12401 := Call(__e, PrimFunc(symshen_4newpv), V4866)


__e.TailApply(tmp12396, tmp12401)
return


}, 1)

tmp12402 := PrimTail(W4872)

__e.TailApply(tmp12395, tmp12402)
return


}, 1)

tmp12403 := PrimHead(W4872)

__e.TailApply(tmp12394, tmp12403)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12406 := Call(__e, PrimFunc(symshen_4lazyderef), V4863, V4866)


__e.TailApply(tmp12393, tmp12406)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4870)
return
}


}, 1)

tmp12418 := Call(__e, PrimFunc(symshen_4unlocked_2), V4867)


var ifres12411 Obj

if True == tmp12418 {
tmp12412 := MakeNative(func(__e *ControlFlow) {
W4871 := __e.Get(1)
_ = W4871
tmp12415 := PrimEqual(W4871, Nil)

if True == tmp12415 {
tmp12413 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12413

__e.TailApply(PrimFunc(symis_b), V4864, V4865, V4866, V4867, V4868, V4869)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12416 := Call(__e, PrimFunc(symshen_4lazyderef), V4863, V4866)


tmp12417 := Call(__e, tmp12412, tmp12416)


ifres12411 = tmp12417


} else {
ifres12411 = False


}

__e.TailApply(tmp12392, ifres12411)
return


}, 7)

tmp12419 := Call(__e, ns2_1set, symshen_4insert_1prolog_1variables, tmp12391)


_ = tmp12419

tmp12420 := MakeNative(func(__e *ControlFlow) {
V4876 := __e.Get(1)
_ = V4876
V4877 := __e.Get(2)
_ = V4877
V4878 := __e.Get(3)
_ = V4878
V4879 := __e.Get(4)
_ = V4879
V4880 := __e.Get(5)
_ = V4880
V4881 := __e.Get(6)
_ = V4881
tmp12421 := MakeNative(func(__e *ControlFlow) {
W4882 := __e.Get(1)
_ = W4882
tmp12422 := MakeNative(func(__e *ControlFlow) {
W4883 := __e.Get(1)
_ = W4883
tmp12435 := PrimEqual(W4883, False)

if True == tmp12435 {
tmp12423 := MakeNative(func(__e *ControlFlow) {
W4889 := __e.Get(1)
_ = W4889
tmp12425 := PrimEqual(W4889, False)

if True == tmp12425 {
__e.TailApply(PrimFunc(symshen_4unlock), V4879, W4882)
return
} else {
__e.Return(W4889)
return
}


}, 1)

tmp12433 := Call(__e, PrimFunc(symshen_4unlocked_2), V4879)


var ifres12426 Obj

if True == tmp12433 {
tmp12427 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12427

tmp12428 := PrimIntern(MakeString(":"))

tmp12429 := PrimCons(V4877, Nil)

tmp12430 := PrimCons(tmp12428, tmp12429)

tmp12431 := PrimCons(V4876, tmp12430)

tmp12432 := Call(__e, PrimFunc(symshen_4system_1S), tmp12431, Nil, V4878, V4879, W4882, V4881)


ifres12426 = tmp12432


} else {
ifres12426 = False


}

__e.TailApply(tmp12423, ifres12426)
return


} else {
__e.Return(W4883)
return
}


}, 1)

tmp12464 := Call(__e, PrimFunc(symshen_4unlocked_2), V4879)


var ifres12436 Obj

if True == tmp12464 {
tmp12437 := MakeNative(func(__e *ControlFlow) {
W4884 := __e.Get(1)
_ = W4884
tmp12461 := PrimIsPair(W4884)

if True == tmp12461 {
tmp12438 := MakeNative(func(__e *ControlFlow) {
W4885 := __e.Get(1)
_ = W4885
tmp12457 := PrimEqual(W4885, symdefine)

if True == tmp12457 {
tmp12439 := MakeNative(func(__e *ControlFlow) {
W4886 := __e.Get(1)
_ = W4886
tmp12453 := PrimIsPair(W4886)

if True == tmp12453 {
tmp12440 := MakeNative(func(__e *ControlFlow) {
W4887 := __e.Get(1)
_ = W4887
tmp12441 := MakeNative(func(__e *ControlFlow) {
W4888 := __e.Get(1)
_ = W4888
tmp12442 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12442

tmp12443 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp12444 := MakeNative(func(__e *ControlFlow) {
tmp12445 := MakeNative(func(__e *ControlFlow) {
tmp12446 := PrimValue(symshen_4_dspy_d)

tmp12447 := MakeNative(func(__e *ControlFlow) {
tmp12448 := PrimCons(W4887, W4888)

tmp12449 := PrimCons(symdefine, tmp12448)

__e.TailApply(PrimFunc(symshen_4t_d), tmp12449, V4877, V4878, V4879, W4882, V4881)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4signal_1def), tmp12446, W4887, V4878, V4879, W4882, tmp12447)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4878, V4879, W4882, tmp12445)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp12443, V4878, V4879, W4882, tmp12444)
return


}, 1)

tmp12450 := PrimTail(W4886)

__e.TailApply(tmp12441, tmp12450)
return


}, 1)

tmp12451 := PrimHead(W4886)

__e.TailApply(tmp12440, tmp12451)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12454 := PrimTail(W4884)

tmp12455 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12454, V4878)


__e.TailApply(tmp12439, tmp12455)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12458 := PrimHead(W4884)

tmp12459 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12458, V4878)


__e.TailApply(tmp12438, tmp12459)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12462 := Call(__e, PrimFunc(symshen_4lazyderef), V4876, V4878)


tmp12463 := Call(__e, tmp12437, tmp12462)


ifres12436 = tmp12463


} else {
ifres12436 = False


}

__e.TailApply(tmp12422, ifres12436)
return


}, 1)

tmp12465 := PrimNumberAdd(V4880, MakeNumber(1))

__e.TailApply(tmp12421, tmp12465)
return


}, 6)

tmp12466 := Call(__e, ns2_1set, symshen_4toplevel_1forms, tmp12420)


_ = tmp12466

tmp12467 := MakeNative(func(__e *ControlFlow) {
V4890 := __e.Get(1)
_ = V4890
V4891 := __e.Get(2)
_ = V4891
V4892 := __e.Get(3)
_ = V4892
V4893 := __e.Get(4)
_ = V4893
V4894 := __e.Get(5)
_ = V4894
V4895 := __e.Get(6)
_ = V4895
tmp12468 := MakeNative(func(__e *ControlFlow) {
W4896 := __e.Get(1)
_ = W4896
tmp12485 := PrimEqual(W4896, False)

if True == tmp12485 {
tmp12483 := Call(__e, PrimFunc(symshen_4unlocked_2), V4893)


if True == tmp12483 {
tmp12469 := MakeNative(func(__e *ControlFlow) {
W4898 := __e.Get(1)
_ = W4898
tmp12480 := PrimEqual(W4898, True)

if True == tmp12480 {
tmp12470 := MakeNative(func(__e *ControlFlow) {
W4899 := __e.Get(1)
_ = W4899
tmp12471 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12471

tmp12472 := Call(__e, PrimFunc(symshen_4deref), V4891, V4892)


tmp12473 := Call(__e, PrimFunc(symshen_4app), tmp12472, MakeString(")\n"), symshen_4a)


tmp12474 := PrimStringConcat(MakeString("\ntypechecking (fn "), tmp12473)

tmp12475 := Call(__e, PrimFunc(symstoutput))


tmp12476 := Call(__e, PrimFunc(sympr), tmp12474, tmp12475)


tmp12477 := Call(__e, PrimFunc(symis), W4899, tmp12476, V4892, V4893, V4894, V4895)


__e.TailApply(PrimFunc(symshen_4gc), V4892, tmp12477)
return


}, 1)

tmp12478 := Call(__e, PrimFunc(symshen_4newpv), V4892)


__e.TailApply(tmp12470, tmp12478)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12481 := Call(__e, PrimFunc(symshen_4lazyderef), V4890, V4892)


__e.TailApply(tmp12469, tmp12481)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W4896)
return
}


}, 1)

tmp12493 := Call(__e, PrimFunc(symshen_4unlocked_2), V4893)


var ifres12486 Obj

if True == tmp12493 {
tmp12487 := MakeNative(func(__e *ControlFlow) {
W4897 := __e.Get(1)
_ = W4897
tmp12490 := PrimEqual(W4897, False)

if True == tmp12490 {
tmp12488 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12488

__e.TailApply(PrimFunc(symthaw), V4895)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12491 := Call(__e, PrimFunc(symshen_4lazyderef), V4890, V4892)


tmp12492 := Call(__e, tmp12487, tmp12491)


ifres12486 = tmp12492


} else {
ifres12486 = False


}

__e.TailApply(tmp12468, ifres12486)
return


}, 6)

tmp12494 := Call(__e, ns2_1set, symshen_4signal_1def, tmp12467)


_ = tmp12494

tmp12495 := MakeNative(func(__e *ControlFlow) {
V4900 := __e.Get(1)
_ = V4900
tmp12496 := Call(__e, PrimFunc(symshen_4curry_1type), V4900)


__e.TailApply(PrimFunc(symshen_4demodulate), tmp12496)
return


}, 1)

tmp12497 := Call(__e, ns2_1set, symshen_4rectify_1type, tmp12495)


_ = tmp12497

tmp12498 := MakeNative(func(__e *ControlFlow) {
V4901 := __e.Get(1)
_ = V4901
tmp12499 := MakeNative(func(__e *ControlFlow) {
tmp12500 := MakeNative(func(__e *ControlFlow) {
W4902 := __e.Get(1)
_ = W4902
tmp12502 := PrimEqual(W4902, V4901)

if True == tmp12502 {
__e.Return(V4901)
return
} else {
__e.TailApply(PrimFunc(symshen_4demodulate), W4902)
return
}


}, 1)

tmp12503 := MakeNative(func(__e *ControlFlow) {
Z4903 := __e.Get(1)
_ = Z4903
__e.TailApply(PrimFunc(symshen_4demod), Z4903)
return
}, 1)

tmp12504 := Call(__e, PrimFunc(symshen_4walk), tmp12503, V4901)


__e.TailApply(tmp12500, tmp12504)
return


}, 0)

tmp12505 := MakeNative(func(__e *ControlFlow) {
Z4904 := __e.Get(1)
_ = Z4904
__e.Return(V4901)
return
}, 1)

__e.TailApply(try_1catch, tmp12499, tmp12505)
return


}, 1)

tmp12506 := Call(__e, ns2_1set, symshen_4demodulate, tmp12498)


_ = tmp12506

tmp12507 := MakeNative(func(__e *ControlFlow) {
V4905 := __e.Get(1)
_ = V4905
tmp12631 := PrimIsPair(V4905)

var ifres12604 Obj

if True == tmp12631 {
tmp12629 := PrimTail(V4905)

tmp12630 := PrimIsPair(tmp12629)

var ifres12606 Obj

if True == tmp12630 {
tmp12626 := PrimTail(V4905)

tmp12627 := PrimHead(tmp12626)

tmp12628 := PrimEqual(sym_1_1_6, tmp12627)

var ifres12608 Obj

if True == tmp12628 {
tmp12623 := PrimTail(V4905)

tmp12624 := PrimTail(tmp12623)

tmp12625 := PrimIsPair(tmp12624)

var ifres12610 Obj

if True == tmp12625 {
tmp12619 := PrimTail(V4905)

tmp12620 := PrimTail(tmp12619)

tmp12621 := PrimTail(tmp12620)

tmp12622 := PrimIsPair(tmp12621)

var ifres12612 Obj

if True == tmp12622 {
tmp12614 := PrimTail(V4905)

tmp12615 := PrimTail(tmp12614)

tmp12616 := PrimTail(tmp12615)

tmp12617 := PrimHead(tmp12616)

tmp12618 := PrimEqual(sym_1_1_6, tmp12617)

var ifres12613 Obj

if True == tmp12618 {
ifres12613 = True


} else {
ifres12613 = False


}

ifres12612 = ifres12613


} else {
ifres12612 = False


}

var ifres12611 Obj

if True == ifres12612 {
ifres12611 = True


} else {
ifres12611 = False


}

ifres12610 = ifres12611


} else {
ifres12610 = False


}

var ifres12609 Obj

if True == ifres12610 {
ifres12609 = True


} else {
ifres12609 = False


}

ifres12608 = ifres12609


} else {
ifres12608 = False


}

var ifres12607 Obj

if True == ifres12608 {
ifres12607 = True


} else {
ifres12607 = False


}

ifres12606 = ifres12607


} else {
ifres12606 = False


}

var ifres12605 Obj

if True == ifres12606 {
ifres12605 = True


} else {
ifres12605 = False


}

ifres12604 = ifres12605


} else {
ifres12604 = False


}

if True == ifres12604 {
tmp12508 := PrimHead(V4905)

tmp12509 := PrimTail(V4905)

tmp12510 := PrimTail(tmp12509)

tmp12511 := PrimCons(tmp12510, Nil)

tmp12512 := PrimCons(sym_1_1_6, tmp12511)

tmp12513 := PrimCons(tmp12508, tmp12512)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12513)
return


} else {
tmp12602 := PrimIsPair(V4905)

var ifres12562 Obj

if True == tmp12602 {
tmp12600 := PrimHead(V4905)

tmp12601 := PrimIsPair(tmp12600)

var ifres12564 Obj

if True == tmp12601 {
tmp12597 := PrimHead(V4905)

tmp12598 := PrimHead(tmp12597)

tmp12599 := PrimEqual(symlist, tmp12598)

var ifres12566 Obj

if True == tmp12599 {
tmp12594 := PrimHead(V4905)

tmp12595 := PrimTail(tmp12594)

tmp12596 := PrimIsPair(tmp12595)

var ifres12568 Obj

if True == tmp12596 {
tmp12590 := PrimHead(V4905)

tmp12591 := PrimTail(tmp12590)

tmp12592 := PrimTail(tmp12591)

tmp12593 := PrimEqual(Nil, tmp12592)

var ifres12570 Obj

if True == tmp12593 {
tmp12588 := PrimTail(V4905)

tmp12589 := PrimIsPair(tmp12588)

var ifres12572 Obj

if True == tmp12589 {
tmp12585 := PrimTail(V4905)

tmp12586 := PrimHead(tmp12585)

tmp12587 := PrimEqual(sym_a_a_6, tmp12586)

var ifres12574 Obj

if True == tmp12587 {
tmp12582 := PrimTail(V4905)

tmp12583 := PrimTail(tmp12582)

tmp12584 := PrimIsPair(tmp12583)

var ifres12576 Obj

if True == tmp12584 {
tmp12578 := PrimTail(V4905)

tmp12579 := PrimTail(tmp12578)

tmp12580 := PrimTail(tmp12579)

tmp12581 := PrimEqual(Nil, tmp12580)

var ifres12577 Obj

if True == tmp12581 {
ifres12577 = True


} else {
ifres12577 = False


}

ifres12576 = ifres12577


} else {
ifres12576 = False


}

var ifres12575 Obj

if True == ifres12576 {
ifres12575 = True


} else {
ifres12575 = False


}

ifres12574 = ifres12575


} else {
ifres12574 = False


}

var ifres12573 Obj

if True == ifres12574 {
ifres12573 = True


} else {
ifres12573 = False


}

ifres12572 = ifres12573


} else {
ifres12572 = False


}

var ifres12571 Obj

if True == ifres12572 {
ifres12571 = True


} else {
ifres12571 = False


}

ifres12570 = ifres12571


} else {
ifres12570 = False


}

var ifres12569 Obj

if True == ifres12570 {
ifres12569 = True


} else {
ifres12569 = False


}

ifres12568 = ifres12569


} else {
ifres12568 = False


}

var ifres12567 Obj

if True == ifres12568 {
ifres12567 = True


} else {
ifres12567 = False


}

ifres12566 = ifres12567


} else {
ifres12566 = False


}

var ifres12565 Obj

if True == ifres12566 {
ifres12565 = True


} else {
ifres12565 = False


}

ifres12564 = ifres12565


} else {
ifres12564 = False


}

var ifres12563 Obj

if True == ifres12564 {
ifres12563 = True


} else {
ifres12563 = False


}

ifres12562 = ifres12563


} else {
ifres12562 = False


}

if True == ifres12562 {
tmp12514 := PrimHead(V4905)

tmp12515 := PrimHead(V4905)

tmp12516 := PrimTail(V4905)

tmp12517 := PrimTail(tmp12516)

tmp12518 := PrimCons(tmp12515, tmp12517)

tmp12519 := PrimCons(symstr, tmp12518)

tmp12520 := PrimCons(tmp12519, Nil)

tmp12521 := PrimCons(sym_1_1_6, tmp12520)

tmp12522 := PrimCons(tmp12514, tmp12521)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12522)
return


} else {
tmp12560 := PrimIsPair(V4905)

var ifres12533 Obj

if True == tmp12560 {
tmp12558 := PrimTail(V4905)

tmp12559 := PrimIsPair(tmp12558)

var ifres12535 Obj

if True == tmp12559 {
tmp12555 := PrimTail(V4905)

tmp12556 := PrimHead(tmp12555)

tmp12557 := PrimEqual(sym_d, tmp12556)

var ifres12537 Obj

if True == tmp12557 {
tmp12552 := PrimTail(V4905)

tmp12553 := PrimTail(tmp12552)

tmp12554 := PrimIsPair(tmp12553)

var ifres12539 Obj

if True == tmp12554 {
tmp12548 := PrimTail(V4905)

tmp12549 := PrimTail(tmp12548)

tmp12550 := PrimTail(tmp12549)

tmp12551 := PrimIsPair(tmp12550)

var ifres12541 Obj

if True == tmp12551 {
tmp12543 := PrimTail(V4905)

tmp12544 := PrimTail(tmp12543)

tmp12545 := PrimTail(tmp12544)

tmp12546 := PrimHead(tmp12545)

tmp12547 := PrimEqual(sym_d, tmp12546)

var ifres12542 Obj

if True == tmp12547 {
ifres12542 = True


} else {
ifres12542 = False


}

ifres12541 = ifres12542


} else {
ifres12541 = False


}

var ifres12540 Obj

if True == ifres12541 {
ifres12540 = True


} else {
ifres12540 = False


}

ifres12539 = ifres12540


} else {
ifres12539 = False


}

var ifres12538 Obj

if True == ifres12539 {
ifres12538 = True


} else {
ifres12538 = False


}

ifres12537 = ifres12538


} else {
ifres12537 = False


}

var ifres12536 Obj

if True == ifres12537 {
ifres12536 = True


} else {
ifres12536 = False


}

ifres12535 = ifres12536


} else {
ifres12535 = False


}

var ifres12534 Obj

if True == ifres12535 {
ifres12534 = True


} else {
ifres12534 = False


}

ifres12533 = ifres12534


} else {
ifres12533 = False


}

if True == ifres12533 {
tmp12523 := PrimHead(V4905)

tmp12524 := PrimTail(V4905)

tmp12525 := PrimTail(tmp12524)

tmp12526 := PrimCons(tmp12525, Nil)

tmp12527 := PrimCons(sym_d, tmp12526)

tmp12528 := PrimCons(tmp12523, tmp12527)

__e.TailApply(PrimFunc(symshen_4curry_1type), tmp12528)
return


} else {
tmp12531 := PrimIsPair(V4905)

if True == tmp12531 {
tmp12529 := MakeNative(func(__e *ControlFlow) {
Z4906 := __e.Get(1)
_ = Z4906
__e.TailApply(PrimFunc(symshen_4curry_1type), Z4906)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12529, V4905)
return


} else {
__e.Return(V4905)
return
}


}


}


}


}, 1)

tmp12632 := Call(__e, ns2_1set, symshen_4curry_1type, tmp12507)


_ = tmp12632

tmp12633 := MakeNative(func(__e *ControlFlow) {
V4907 := __e.Get(1)
_ = V4907
tmp12751 := PrimIsPair(V4907)

var ifres12743 Obj

if True == tmp12751 {
tmp12749 := PrimHead(V4907)

tmp12750 := PrimEqual(symdefine, tmp12749)

var ifres12745 Obj

if True == tmp12750 {
tmp12747 := PrimTail(V4907)

tmp12748 := PrimIsPair(tmp12747)

var ifres12746 Obj

if True == tmp12748 {
ifres12746 = True


} else {
ifres12746 = False


}

ifres12745 = ifres12746


} else {
ifres12745 = False


}

var ifres12744 Obj

if True == ifres12745 {
ifres12744 = True


} else {
ifres12744 = False


}

ifres12743 = ifres12744


} else {
ifres12743 = False


}

if True == ifres12743 {
__e.Return(V4907)
return
} else {
tmp12741 := PrimIsPair(V4907)

var ifres12722 Obj

if True == tmp12741 {
tmp12739 := PrimHead(V4907)

tmp12740 := PrimEqual(symtype, tmp12739)

var ifres12724 Obj

if True == tmp12740 {
tmp12737 := PrimTail(V4907)

tmp12738 := PrimIsPair(tmp12737)

var ifres12726 Obj

if True == tmp12738 {
tmp12734 := PrimTail(V4907)

tmp12735 := PrimTail(tmp12734)

tmp12736 := PrimIsPair(tmp12735)

var ifres12728 Obj

if True == tmp12736 {
tmp12730 := PrimTail(V4907)

tmp12731 := PrimTail(tmp12730)

tmp12732 := PrimTail(tmp12731)

tmp12733 := PrimEqual(Nil, tmp12732)

var ifres12729 Obj

if True == tmp12733 {
ifres12729 = True


} else {
ifres12729 = False


}

ifres12728 = ifres12729


} else {
ifres12728 = False


}

var ifres12727 Obj

if True == ifres12728 {
ifres12727 = True


} else {
ifres12727 = False


}

ifres12726 = ifres12727


} else {
ifres12726 = False


}

var ifres12725 Obj

if True == ifres12726 {
ifres12725 = True


} else {
ifres12725 = False


}

ifres12724 = ifres12725


} else {
ifres12724 = False


}

var ifres12723 Obj

if True == ifres12724 {
ifres12723 = True


} else {
ifres12723 = False


}

ifres12722 = ifres12723


} else {
ifres12722 = False


}

if True == ifres12722 {
tmp12634 := PrimTail(V4907)

tmp12635 := PrimHead(tmp12634)

tmp12636 := Call(__e, PrimFunc(symshen_4curry), tmp12635)


tmp12637 := PrimTail(V4907)

tmp12638 := PrimTail(tmp12637)

tmp12639 := PrimCons(tmp12636, tmp12638)

__e.Return(PrimCons(symtype, tmp12639))
return


} else {
tmp12720 := PrimIsPair(V4907)

var ifres12701 Obj

if True == tmp12720 {
tmp12718 := PrimHead(V4907)

tmp12719 := PrimEqual(syminput_7, tmp12718)

var ifres12703 Obj

if True == tmp12719 {
tmp12716 := PrimTail(V4907)

tmp12717 := PrimIsPair(tmp12716)

var ifres12705 Obj

if True == tmp12717 {
tmp12713 := PrimTail(V4907)

tmp12714 := PrimTail(tmp12713)

tmp12715 := PrimIsPair(tmp12714)

var ifres12707 Obj

if True == tmp12715 {
tmp12709 := PrimTail(V4907)

tmp12710 := PrimTail(tmp12709)

tmp12711 := PrimTail(tmp12710)

tmp12712 := PrimEqual(Nil, tmp12711)

var ifres12708 Obj

if True == tmp12712 {
ifres12708 = True


} else {
ifres12708 = False


}

ifres12707 = ifres12708


} else {
ifres12707 = False


}

var ifres12706 Obj

if True == ifres12707 {
ifres12706 = True


} else {
ifres12706 = False


}

ifres12705 = ifres12706


} else {
ifres12705 = False


}

var ifres12704 Obj

if True == ifres12705 {
ifres12704 = True


} else {
ifres12704 = False


}

ifres12703 = ifres12704


} else {
ifres12703 = False


}

var ifres12702 Obj

if True == ifres12703 {
ifres12702 = True


} else {
ifres12702 = False


}

ifres12701 = ifres12702


} else {
ifres12701 = False


}

if True == ifres12701 {
tmp12640 := PrimTail(V4907)

tmp12641 := PrimHead(tmp12640)

tmp12642 := PrimTail(V4907)

tmp12643 := PrimTail(tmp12642)

tmp12644 := PrimHead(tmp12643)

tmp12645 := Call(__e, PrimFunc(symshen_4curry), tmp12644)


tmp12646 := PrimCons(tmp12645, Nil)

tmp12647 := PrimCons(tmp12641, tmp12646)

__e.Return(PrimCons(syminput_7, tmp12647))
return


} else {
tmp12699 := PrimIsPair(V4907)

var ifres12695 Obj

if True == tmp12699 {
tmp12697 := PrimHead(V4907)

tmp12698 := Call(__e, PrimFunc(symshen_4special_2), tmp12697)


var ifres12696 Obj

if True == tmp12698 {
ifres12696 = True


} else {
ifres12696 = False


}

ifres12695 = ifres12696


} else {
ifres12695 = False


}

if True == ifres12695 {
tmp12648 := PrimHead(V4907)

tmp12649 := MakeNative(func(__e *ControlFlow) {
Z4908 := __e.Get(1)
_ = Z4908
__e.TailApply(PrimFunc(symshen_4curry), Z4908)
return
}, 1)

tmp12650 := PrimTail(V4907)

tmp12651 := Call(__e, PrimFunc(symmap), tmp12649, tmp12650)


__e.Return(PrimCons(tmp12648, tmp12651))
return


} else {
tmp12693 := PrimIsPair(V4907)

var ifres12689 Obj

if True == tmp12693 {
tmp12691 := PrimHead(V4907)

tmp12692 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp12691)


var ifres12690 Obj

if True == tmp12692 {
ifres12690 = True


} else {
ifres12690 = False


}

ifres12689 = ifres12690


} else {
ifres12689 = False


}

if True == ifres12689 {
__e.Return(V4907)
return
} else {
tmp12687 := PrimIsPair(V4907)

var ifres12678 Obj

if True == tmp12687 {
tmp12685 := PrimTail(V4907)

tmp12686 := PrimIsPair(tmp12685)

var ifres12680 Obj

if True == tmp12686 {
tmp12682 := PrimTail(V4907)

tmp12683 := PrimTail(tmp12682)

tmp12684 := PrimIsPair(tmp12683)

var ifres12681 Obj

if True == tmp12684 {
ifres12681 = True


} else {
ifres12681 = False


}

ifres12680 = ifres12681


} else {
ifres12680 = False


}

var ifres12679 Obj

if True == ifres12680 {
ifres12679 = True


} else {
ifres12679 = False


}

ifres12678 = ifres12679


} else {
ifres12678 = False


}

if True == ifres12678 {
tmp12652 := PrimHead(V4907)

tmp12653 := PrimTail(V4907)

tmp12654 := PrimHead(tmp12653)

tmp12655 := PrimCons(tmp12654, Nil)

tmp12656 := PrimCons(tmp12652, tmp12655)

tmp12657 := PrimTail(V4907)

tmp12658 := PrimTail(tmp12657)

tmp12659 := PrimCons(tmp12656, tmp12658)

__e.TailApply(PrimFunc(symshen_4curry), tmp12659)
return


} else {
tmp12676 := PrimIsPair(V4907)

var ifres12667 Obj

if True == tmp12676 {
tmp12674 := PrimTail(V4907)

tmp12675 := PrimIsPair(tmp12674)

var ifres12669 Obj

if True == tmp12675 {
tmp12671 := PrimTail(V4907)

tmp12672 := PrimTail(tmp12671)

tmp12673 := PrimEqual(Nil, tmp12672)

var ifres12670 Obj

if True == tmp12673 {
ifres12670 = True


} else {
ifres12670 = False


}

ifres12669 = ifres12670


} else {
ifres12669 = False


}

var ifres12668 Obj

if True == ifres12669 {
ifres12668 = True


} else {
ifres12668 = False


}

ifres12667 = ifres12668


} else {
ifres12667 = False


}

if True == ifres12667 {
tmp12660 := PrimHead(V4907)

tmp12661 := Call(__e, PrimFunc(symshen_4curry), tmp12660)


tmp12662 := PrimTail(V4907)

tmp12663 := PrimHead(tmp12662)

tmp12664 := Call(__e, PrimFunc(symshen_4curry), tmp12663)


tmp12665 := PrimCons(tmp12664, Nil)

__e.Return(PrimCons(tmp12661, tmp12665))
return


} else {
__e.Return(V4907)
return
}


}


}


}


}


}


}


}, 1)

tmp12752 := Call(__e, ns2_1set, symshen_4curry, tmp12633)


_ = tmp12752

tmp12753 := MakeNative(func(__e *ControlFlow) {
V4909 := __e.Get(1)
_ = V4909
tmp12754 := PrimValue(symshen_4_dspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4909, tmp12754)
return


}, 1)

tmp12755 := Call(__e, ns2_1set, symshen_4special_2, tmp12753)


_ = tmp12755

tmp12756 := MakeNative(func(__e *ControlFlow) {
V4910 := __e.Get(1)
_ = V4910
tmp12757 := PrimValue(symshen_4_dextraspecial_d)

__e.TailApply(PrimFunc(symelement_2), V4910, tmp12757)
return


}, 1)

tmp12758 := Call(__e, ns2_1set, symshen_4extraspecial_2, tmp12756)


_ = tmp12758

tmp12759 := MakeNative(func(__e *ControlFlow) {
V4911 := __e.Get(1)
_ = V4911
V4912 := __e.Get(2)
_ = V4912
V4913 := __e.Get(3)
_ = V4913
V4914 := __e.Get(4)
_ = V4914
V4915 := __e.Get(5)
_ = V4915
V4916 := __e.Get(6)
_ = V4916
tmp12760 := MakeNative(func(__e *ControlFlow) {
W4917 := __e.Get(1)
_ = W4917
tmp12761 := MakeNative(func(__e *ControlFlow) {
W4918 := __e.Get(1)
_ = W4918
tmp12819 := PrimEqual(W4918, False)

if True == tmp12819 {
tmp12762 := MakeNative(func(__e *ControlFlow) {
W4919 := __e.Get(1)
_ = W4919
tmp12781 := PrimEqual(W4919, False)

if True == tmp12781 {
tmp12763 := MakeNative(func(__e *ControlFlow) {
W4927 := __e.Get(1)
_ = W4927
tmp12773 := PrimEqual(W4927, False)

if True == tmp12773 {
tmp12764 := MakeNative(func(__e *ControlFlow) {
W4928 := __e.Get(1)
_ = W4928
tmp12766 := PrimEqual(W4928, False)

if True == tmp12766 {
__e.TailApply(PrimFunc(symshen_4unlock), V4914, W4917)
return
} else {
__e.Return(W4928)
return
}


}, 1)

tmp12771 := Call(__e, PrimFunc(symshen_4unlocked_2), V4914)


var ifres12767 Obj

if True == tmp12771 {
tmp12768 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12768

tmp12769 := PrimValue(symshen_4_ddatatypes_d)

tmp12770 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), V4911, V4912, tmp12769, V4913, V4914, W4917, V4916)


ifres12767 = tmp12770


} else {
ifres12767 = False


}

__e.TailApply(tmp12764, ifres12767)
return


} else {
__e.Return(W4927)
return
}


}, 1)

tmp12779 := Call(__e, PrimFunc(symshen_4unlocked_2), V4914)


var ifres12774 Obj

if True == tmp12779 {
tmp12775 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12775

tmp12776 := PrimValue(symshen_4_dspy_d)

tmp12777 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4show), V4911, V4912, V4913, V4914, W4917, V4916)
return
}, 0)

tmp12778 := Call(__e, PrimFunc(symwhen), tmp12776, V4913, V4914, W4917, tmp12777)


ifres12774 = tmp12778


} else {
ifres12774 = False


}

__e.TailApply(tmp12763, ifres12774)
return


} else {
__e.Return(W4919)
return
}


}, 1)

tmp12817 := Call(__e, PrimFunc(symshen_4unlocked_2), V4914)


var ifres12782 Obj

if True == tmp12817 {
tmp12783 := MakeNative(func(__e *ControlFlow) {
W4920 := __e.Get(1)
_ = W4920
tmp12814 := PrimIsPair(W4920)

if True == tmp12814 {
tmp12784 := MakeNative(func(__e *ControlFlow) {
W4921 := __e.Get(1)
_ = W4921
tmp12785 := MakeNative(func(__e *ControlFlow) {
W4922 := __e.Get(1)
_ = W4922
tmp12809 := PrimIsPair(W4922)

if True == tmp12809 {
tmp12786 := MakeNative(func(__e *ControlFlow) {
W4923 := __e.Get(1)
_ = W4923
tmp12787 := MakeNative(func(__e *ControlFlow) {
W4924 := __e.Get(1)
_ = W4924
tmp12804 := PrimIsPair(W4924)

if True == tmp12804 {
tmp12788 := MakeNative(func(__e *ControlFlow) {
W4925 := __e.Get(1)
_ = W4925
tmp12789 := MakeNative(func(__e *ControlFlow) {
W4926 := __e.Get(1)
_ = W4926
tmp12799 := PrimEqual(W4926, Nil)

if True == tmp12799 {
tmp12790 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12790

tmp12791 := Call(__e, PrimFunc(symshen_4deref), W4923, V4913)


tmp12792 := PrimIntern(MakeString(":"))

tmp12793 := PrimEqual(tmp12791, tmp12792)

tmp12794 := MakeNative(func(__e *ControlFlow) {
tmp12795 := Call(__e, PrimFunc(symshen_4type_1theory_1enabled_2))


tmp12796 := MakeNative(func(__e *ControlFlow) {
tmp12797 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4921, W4925, V4912, V4913, V4914, W4917, V4916)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4913, V4914, W4917, tmp12797)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp12795, V4913, V4914, W4917, tmp12796)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp12793, V4913, V4914, W4917, tmp12794)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12800 := PrimTail(W4924)

tmp12801 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12800, V4913)


__e.TailApply(tmp12789, tmp12801)
return


}, 1)

tmp12802 := PrimHead(W4924)

__e.TailApply(tmp12788, tmp12802)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12805 := PrimTail(W4922)

tmp12806 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12805, V4913)


__e.TailApply(tmp12787, tmp12806)
return


}, 1)

tmp12807 := PrimHead(W4922)

__e.TailApply(tmp12786, tmp12807)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12810 := PrimTail(W4920)

tmp12811 := Call(__e, PrimFunc(symshen_4lazyderef), tmp12810, V4913)


__e.TailApply(tmp12785, tmp12811)
return


}, 1)

tmp12812 := PrimHead(W4920)

__e.TailApply(tmp12784, tmp12812)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp12815 := Call(__e, PrimFunc(symshen_4lazyderef), V4911, V4913)


tmp12816 := Call(__e, tmp12783, tmp12815)


ifres12782 = tmp12816


} else {
ifres12782 = False


}

__e.TailApply(tmp12762, ifres12782)
return


} else {
__e.Return(W4918)
return
}


}, 1)

tmp12824 := Call(__e, PrimFunc(symshen_4unlocked_2), V4914)


var ifres12820 Obj

if True == tmp12824 {
tmp12821 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp12821

tmp12822 := Call(__e, PrimFunc(symshen_4maxinfexceeded_2))


tmp12823 := Call(__e, PrimFunc(symwhen), tmp12822, V4913, V4914, W4917, V4916)


ifres12820 = tmp12823


} else {
ifres12820 = False


}

__e.TailApply(tmp12761, ifres12820)
return


}, 1)

tmp12825 := PrimNumberAdd(V4915, MakeNumber(1))

__e.TailApply(tmp12760, tmp12825)
return


}, 6)

tmp12826 := Call(__e, ns2_1set, symshen_4system_1S, tmp12759)


_ = tmp12826

tmp12827 := MakeNative(func(__e *ControlFlow) {
V4935 := __e.Get(1)
_ = V4935
V4936 := __e.Get(2)
_ = V4936
V4937 := __e.Get(3)
_ = V4937
V4938 := __e.Get(4)
_ = V4938
V4939 := __e.Get(5)
_ = V4939
V4940 := __e.Get(6)
_ = V4940
tmp12828 := Call(__e, PrimFunc(symshen_4line))


_ = tmp12828

tmp12829 := Call(__e, PrimFunc(symshen_4deref), V4935, V4937)


tmp12830 := Call(__e, PrimFunc(symshen_4show_1p), tmp12829)


_ = tmp12830

tmp12831 := Call(__e, PrimFunc(symnl), MakeNumber(2))


_ = tmp12831

tmp12832 := Call(__e, PrimFunc(symshen_4deref), V4936, V4937)


tmp12833 := Call(__e, PrimFunc(symshen_4show_1assumptions), tmp12832, MakeNumber(1))


_ = tmp12833

tmp12834 := Call(__e, PrimFunc(symshen_4pause_1for_1user))


_ = tmp12834

__e.Return(False)
return


}, 6)

tmp12835 := Call(__e, ns2_1set, symshen_4show, tmp12827)


_ = tmp12835

tmp12836 := MakeNative(func(__e *ControlFlow) {
tmp12837 := MakeNative(func(__e *ControlFlow) {
W4941 := __e.Get(1)
_ = W4941
tmp12839 := PrimEqual(MakeNumber(1), W4941)

var ifres12838 Obj

if True == tmp12839 {
ifres12838 = MakeString("")


} else {
ifres12838 = MakeString("s")


}

tmp12840 := Call(__e, PrimFunc(symshen_4app), ifres12838, MakeString(" \n?- "), symshen_4a)


tmp12841 := PrimStringConcat(MakeString(" inference"), tmp12840)

tmp12842 := Call(__e, PrimFunc(symshen_4app), W4941, tmp12841, symshen_4a)


tmp12843 := PrimStringConcat(MakeString("____________________________________________________________ "), tmp12842)

tmp12844 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp12843, tmp12844)
return


}, 1)

tmp12845 := Call(__e, PrimFunc(syminferences))


__e.TailApply(tmp12837, tmp12845)
return


}, 0)

tmp12846 := Call(__e, ns2_1set, symshen_4line, tmp12836)


_ = tmp12846

tmp12847 := MakeNative(func(__e *ControlFlow) {
V4942 := __e.Get(1)
_ = V4942
tmp12879 := PrimIsPair(V4942)

var ifres12858 Obj

if True == tmp12879 {
tmp12877 := PrimTail(V4942)

tmp12878 := PrimIsPair(tmp12877)

var ifres12860 Obj

if True == tmp12878 {
tmp12874 := PrimTail(V4942)

tmp12875 := PrimTail(tmp12874)

tmp12876 := PrimIsPair(tmp12875)

var ifres12862 Obj

if True == tmp12876 {
tmp12870 := PrimTail(V4942)

tmp12871 := PrimTail(tmp12870)

tmp12872 := PrimTail(tmp12871)

tmp12873 := PrimEqual(Nil, tmp12872)

var ifres12864 Obj

if True == tmp12873 {
tmp12866 := PrimTail(V4942)

tmp12867 := PrimHead(tmp12866)

tmp12868 := PrimIntern(MakeString(":"))

tmp12869 := PrimEqual(tmp12867, tmp12868)

var ifres12865 Obj

if True == tmp12869 {
ifres12865 = True


} else {
ifres12865 = False


}

ifres12864 = ifres12865


} else {
ifres12864 = False


}

var ifres12863 Obj

if True == ifres12864 {
ifres12863 = True


} else {
ifres12863 = False


}

ifres12862 = ifres12863


} else {
ifres12862 = False


}

var ifres12861 Obj

if True == ifres12862 {
ifres12861 = True


} else {
ifres12861 = False


}

ifres12860 = ifres12861


} else {
ifres12860 = False


}

var ifres12859 Obj

if True == ifres12860 {
ifres12859 = True


} else {
ifres12859 = False


}

ifres12858 = ifres12859


} else {
ifres12858 = False


}

if True == ifres12858 {
tmp12848 := PrimHead(V4942)

tmp12849 := Call(__e, PrimFunc(symshen_4prterm), tmp12848)


_ = tmp12849

tmp12850 := Call(__e, PrimFunc(symstoutput))


tmp12851 := Call(__e, PrimFunc(sympr), MakeString(" : "), tmp12850)


_ = tmp12851

tmp12852 := PrimTail(V4942)

tmp12853 := PrimTail(tmp12852)

tmp12854 := PrimHead(tmp12853)

tmp12855 := Call(__e, PrimFunc(symshen_4app), tmp12854, MakeString(""), symshen_4r)


tmp12856 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp12855, tmp12856)
return


} else {
__e.TailApply(PrimFunc(symshen_4prterm), V4942)
return
}


}, 1)

tmp12880 := Call(__e, ns2_1set, symshen_4show_1p, tmp12847)


_ = tmp12880

tmp12881 := MakeNative(func(__e *ControlFlow) {
V4943 := __e.Get(1)
_ = V4943
tmp12924 := PrimIsPair(V4943)

var ifres12905 Obj

if True == tmp12924 {
tmp12922 := PrimHead(V4943)

tmp12923 := PrimEqual(symcons, tmp12922)

var ifres12907 Obj

if True == tmp12923 {
tmp12920 := PrimTail(V4943)

tmp12921 := PrimIsPair(tmp12920)

var ifres12909 Obj

if True == tmp12921 {
tmp12917 := PrimTail(V4943)

tmp12918 := PrimTail(tmp12917)

tmp12919 := PrimIsPair(tmp12918)

var ifres12911 Obj

if True == tmp12919 {
tmp12913 := PrimTail(V4943)

tmp12914 := PrimTail(tmp12913)

tmp12915 := PrimTail(tmp12914)

tmp12916 := PrimEqual(Nil, tmp12915)

var ifres12912 Obj

if True == tmp12916 {
ifres12912 = True


} else {
ifres12912 = False


}

ifres12911 = ifres12912


} else {
ifres12911 = False


}

var ifres12910 Obj

if True == ifres12911 {
ifres12910 = True


} else {
ifres12910 = False


}

ifres12909 = ifres12910


} else {
ifres12909 = False


}

var ifres12908 Obj

if True == ifres12909 {
ifres12908 = True


} else {
ifres12908 = False


}

ifres12907 = ifres12908


} else {
ifres12907 = False


}

var ifres12906 Obj

if True == ifres12907 {
ifres12906 = True


} else {
ifres12906 = False


}

ifres12905 = ifres12906


} else {
ifres12905 = False


}

if True == ifres12905 {
tmp12882 := Call(__e, PrimFunc(symstoutput))


tmp12883 := Call(__e, PrimFunc(sympr), MakeString("["), tmp12882)


_ = tmp12883

tmp12884 := PrimTail(V4943)

tmp12885 := PrimHead(tmp12884)

tmp12886 := Call(__e, PrimFunc(symshen_4prterm), tmp12885)


_ = tmp12886

tmp12887 := PrimTail(V4943)

tmp12888 := PrimTail(tmp12887)

tmp12889 := PrimHead(tmp12888)

tmp12890 := Call(__e, PrimFunc(symshen_4prtl), tmp12889)


_ = tmp12890

tmp12891 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("]"), tmp12891)
return


} else {
tmp12903 := PrimIsPair(V4943)

if True == tmp12903 {
tmp12892 := Call(__e, PrimFunc(symstoutput))


tmp12893 := Call(__e, PrimFunc(sympr), MakeString("("), tmp12892)


_ = tmp12893

tmp12894 := PrimHead(V4943)

tmp12895 := Call(__e, PrimFunc(symshen_4prterm), tmp12894)


_ = tmp12895

tmp12896 := MakeNative(func(__e *ControlFlow) {
Z4944 := __e.Get(1)
_ = Z4944
tmp12897 := Call(__e, PrimFunc(symstoutput))


tmp12898 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp12897)


_ = tmp12898

__e.TailApply(PrimFunc(symshen_4prterm), Z4944)
return


}, 1)

tmp12899 := PrimTail(V4943)

tmp12900 := Call(__e, PrimFunc(symmap), tmp12896, tmp12899)


_ = tmp12900

tmp12901 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(")"), tmp12901)
return


} else {
__e.TailApply(PrimFunc(symprint), V4943)
return
}


}


}, 1)

tmp12925 := Call(__e, ns2_1set, symshen_4prterm, tmp12881)


_ = tmp12925

tmp12926 := MakeNative(func(__e *ControlFlow) {
V4945 := __e.Get(1)
_ = V4945
tmp12959 := PrimEqual(Nil, V4945)

if True == tmp12959 {
__e.Return(MakeString(""))
return
} else {
tmp12957 := PrimIsPair(V4945)

var ifres12938 Obj

if True == tmp12957 {
tmp12955 := PrimHead(V4945)

tmp12956 := PrimEqual(symcons, tmp12955)

var ifres12940 Obj

if True == tmp12956 {
tmp12953 := PrimTail(V4945)

tmp12954 := PrimIsPair(tmp12953)

var ifres12942 Obj

if True == tmp12954 {
tmp12950 := PrimTail(V4945)

tmp12951 := PrimTail(tmp12950)

tmp12952 := PrimIsPair(tmp12951)

var ifres12944 Obj

if True == tmp12952 {
tmp12946 := PrimTail(V4945)

tmp12947 := PrimTail(tmp12946)

tmp12948 := PrimTail(tmp12947)

tmp12949 := PrimEqual(Nil, tmp12948)

var ifres12945 Obj

if True == tmp12949 {
ifres12945 = True


} else {
ifres12945 = False


}

ifres12944 = ifres12945


} else {
ifres12944 = False


}

var ifres12943 Obj

if True == ifres12944 {
ifres12943 = True


} else {
ifres12943 = False


}

ifres12942 = ifres12943


} else {
ifres12942 = False


}

var ifres12941 Obj

if True == ifres12942 {
ifres12941 = True


} else {
ifres12941 = False


}

ifres12940 = ifres12941


} else {
ifres12940 = False


}

var ifres12939 Obj

if True == ifres12940 {
ifres12939 = True


} else {
ifres12939 = False


}

ifres12938 = ifres12939


} else {
ifres12938 = False


}

if True == ifres12938 {
tmp12927 := Call(__e, PrimFunc(symstoutput))


tmp12928 := Call(__e, PrimFunc(sympr), MakeString(" "), tmp12927)


_ = tmp12928

tmp12929 := PrimTail(V4945)

tmp12930 := PrimHead(tmp12929)

tmp12931 := Call(__e, PrimFunc(symshen_4prterm), tmp12930)


_ = tmp12931

tmp12932 := PrimTail(V4945)

tmp12933 := PrimTail(tmp12932)

tmp12934 := PrimHead(tmp12933)

__e.TailApply(PrimFunc(symshen_4prtl), tmp12934)
return


} else {
tmp12935 := Call(__e, PrimFunc(symstoutput))


tmp12936 := Call(__e, PrimFunc(sympr), MakeString(" | "), tmp12935)


_ = tmp12936

__e.TailApply(PrimFunc(symshen_4prterm), V4945)
return


}


}


}, 1)

tmp12960 := Call(__e, ns2_1set, symshen_4prtl, tmp12926)


_ = tmp12960

tmp12961 := MakeNative(func(__e *ControlFlow) {
V4952 := __e.Get(1)
_ = V4952
V4953 := __e.Get(2)
_ = V4953
tmp12974 := PrimEqual(Nil, V4952)

if True == tmp12974 {
tmp12962 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString("\n> "), tmp12962)
return


} else {
tmp12972 := PrimIsPair(V4952)

if True == tmp12972 {
tmp12963 := Call(__e, PrimFunc(symshen_4app), V4953, MakeString(". "), symshen_4a)


tmp12964 := Call(__e, PrimFunc(symstoutput))


tmp12965 := Call(__e, PrimFunc(sympr), tmp12963, tmp12964)


_ = tmp12965

tmp12966 := PrimHead(V4952)

tmp12967 := Call(__e, PrimFunc(symshen_4show_1p), tmp12966)


_ = tmp12967

tmp12968 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp12968

tmp12969 := PrimTail(V4952)

tmp12970 := PrimNumberAdd(V4953, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4show_1assumptions), tmp12969, tmp12970)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.show-assumptions")))
return
}


}


}, 2)

tmp12975 := Call(__e, ns2_1set, symshen_4show_1assumptions, tmp12961)


_ = tmp12975

tmp12976 := MakeNative(func(__e *ControlFlow) {
tmp12977 := MakeNative(func(__e *ControlFlow) {
W4954 := __e.Get(1)
_ = W4954
tmp12979 := PrimEqual(W4954, MakeNumber(94))

if True == tmp12979 {
__e.Return(PrimSimpleError(MakeString("input aborted\n")))
return
} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 1)

tmp12980 := Call(__e, PrimFunc(symstinput))


tmp12981 := PrimReadByte(tmp12980)

__e.TailApply(tmp12977, tmp12981)
return


}, 0)

tmp12982 := Call(__e, ns2_1set, symshen_4pause_1for_1user, tmp12976)


_ = tmp12982

tmp12983 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dshen_1type_1theory_1enabled_2_d))
return
}, 0)

tmp12984 := Call(__e, ns2_1set, symshen_4type_1theory_1enabled_2, tmp12983)


_ = tmp12984

tmp12985 := MakeNative(func(__e *ControlFlow) {
tmp12987 := Call(__e, PrimFunc(syminferences))


tmp12988 := PrimValue(symshen_4_dmaxinferences_d)

tmp12989 := PrimGreatThan(tmp12987, tmp12988)

if True == tmp12989 {
__e.Return(PrimSimpleError(MakeString("maximum inferences exceeded")))
return
} else {
__e.Return(False)
return
}


}, 0)

tmp12990 := Call(__e, ns2_1set, symshen_4maxinfexceeded_2, tmp12985)


_ = tmp12990

tmp12991 := MakeNative(func(__e *ControlFlow) {
V4955 := __e.Get(1)
_ = V4955
V4956 := __e.Get(2)
_ = V4956
V4957 := __e.Get(3)
_ = V4957
V4958 := __e.Get(4)
_ = V4958
V4959 := __e.Get(5)
_ = V4959
V4960 := __e.Get(6)
_ = V4960
V4961 := __e.Get(7)
_ = V4961
tmp12992 := MakeNative(func(__e *ControlFlow) {
W4962 := __e.Get(1)
_ = W4962
tmp12993 := MakeNative(func(__e *ControlFlow) {
W4963 := __e.Get(1)
_ = W4963
tmp13907 := PrimEqual(W4963, False)

if True == tmp13907 {
tmp12994 := MakeNative(func(__e *ControlFlow) {
W4964 := __e.Get(1)
_ = W4964
tmp13897 := PrimEqual(W4964, False)

if True == tmp13897 {
tmp12995 := MakeNative(func(__e *ControlFlow) {
W4965 := __e.Get(1)
_ = W4965
tmp13891 := PrimEqual(W4965, False)

if True == tmp13891 {
tmp12996 := MakeNative(func(__e *ControlFlow) {
W4966 := __e.Get(1)
_ = W4966
tmp13872 := PrimEqual(W4966, False)

if True == tmp13872 {
tmp12997 := MakeNative(func(__e *ControlFlow) {
W4970 := __e.Get(1)
_ = W4970
tmp13839 := PrimEqual(W4970, False)

if True == tmp13839 {
tmp12998 := MakeNative(func(__e *ControlFlow) {
W4976 := __e.Get(1)
_ = W4976
tmp13812 := PrimEqual(W4976, False)

if True == tmp13812 {
tmp12999 := MakeNative(func(__e *ControlFlow) {
W4982 := __e.Get(1)
_ = W4982
tmp13777 := PrimEqual(W4982, False)

if True == tmp13777 {
tmp13000 := MakeNative(func(__e *ControlFlow) {
W4989 := __e.Get(1)
_ = W4989
tmp13746 := PrimEqual(W4989, False)

if True == tmp13746 {
tmp13001 := MakeNative(func(__e *ControlFlow) {
W4996 := __e.Get(1)
_ = W4996
tmp13661 := PrimEqual(W4996, False)

if True == tmp13661 {
tmp13002 := MakeNative(func(__e *ControlFlow) {
W5017 := __e.Get(1)
_ = W5017
tmp13555 := PrimEqual(W5017, False)

if True == tmp13555 {
tmp13003 := MakeNative(func(__e *ControlFlow) {
W5045 := __e.Get(1)
_ = W5045
tmp13470 := PrimEqual(W5045, False)

if True == tmp13470 {
tmp13004 := MakeNative(func(__e *ControlFlow) {
W5066 := __e.Get(1)
_ = W5066
tmp13427 := PrimEqual(W5066, False)

if True == tmp13427 {
tmp13005 := MakeNative(func(__e *ControlFlow) {
W5076 := __e.Get(1)
_ = W5076
tmp13303 := PrimEqual(W5076, False)

if True == tmp13303 {
tmp13006 := MakeNative(func(__e *ControlFlow) {
W5106 := __e.Get(1)
_ = W5106
tmp13239 := PrimEqual(W5106, False)

if True == tmp13239 {
tmp13007 := MakeNative(func(__e *ControlFlow) {
W5119 := __e.Get(1)
_ = W5119
tmp13151 := PrimEqual(W5119, False)

if True == tmp13151 {
tmp13008 := MakeNative(func(__e *ControlFlow) {
W5140 := __e.Get(1)
_ = W5140
tmp13113 := PrimEqual(W5140, False)

if True == tmp13113 {
tmp13009 := MakeNative(func(__e *ControlFlow) {
W5148 := __e.Get(1)
_ = W5148
tmp13074 := PrimEqual(W5148, False)

if True == tmp13074 {
tmp13010 := MakeNative(func(__e *ControlFlow) {
W5156 := __e.Get(1)
_ = W5156
tmp13036 := PrimEqual(W5156, False)

if True == tmp13036 {
tmp13011 := MakeNative(func(__e *ControlFlow) {
W5164 := __e.Get(1)
_ = W5164
tmp13025 := PrimEqual(W5164, False)

if True == tmp13025 {
tmp13012 := MakeNative(func(__e *ControlFlow) {
W5166 := __e.Get(1)
_ = W5166
tmp13014 := PrimEqual(W5166, False)

if True == tmp13014 {
__e.TailApply(PrimFunc(symshen_4unlock), V4959, W4962)
return
} else {
__e.Return(W5166)
return
}


}, 1)

tmp13023 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13015 Obj

if True == tmp13023 {
tmp13016 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13016

tmp13017 := PrimIntern(MakeString(":"))

tmp13018 := PrimCons(V4956, Nil)

tmp13019 := PrimCons(tmp13017, tmp13018)

tmp13020 := PrimCons(V4955, tmp13019)

tmp13021 := PrimValue(symshen_4_ddatatypes_d)

tmp13022 := Call(__e, PrimFunc(symshen_4search_1user_1datatypes), tmp13020, V4957, tmp13021, V4958, V4959, W4962, V4961)


ifres13015 = tmp13022


} else {
ifres13015 = False


}

__e.TailApply(tmp13012, ifres13015)
return


} else {
__e.Return(W5164)
return
}


}, 1)

tmp13034 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13026 Obj

if True == tmp13034 {
tmp13027 := MakeNative(func(__e *ControlFlow) {
W5165 := __e.Get(1)
_ = W5165
tmp13028 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13028

tmp13029 := MakeNative(func(__e *ControlFlow) {
tmp13030 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), V4955, V4956, W5165, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4958, V4959, W4962, tmp13030)
return


}, 0)

tmp13031 := Call(__e, PrimFunc(symshen_4l_1rules), V4957, W5165, False, V4958, V4959, W4962, tmp13029)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13031)
return


}, 1)

tmp13032 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13033 := Call(__e, tmp13027, tmp13032)


ifres13026 = tmp13033


} else {
ifres13026 = False


}

__e.TailApply(tmp13011, ifres13026)
return


} else {
__e.Return(W5156)
return
}


}, 1)

tmp13072 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13037 Obj

if True == tmp13072 {
tmp13038 := MakeNative(func(__e *ControlFlow) {
W5157 := __e.Get(1)
_ = W5157
tmp13069 := PrimIsPair(W5157)

if True == tmp13069 {
tmp13039 := MakeNative(func(__e *ControlFlow) {
W5158 := __e.Get(1)
_ = W5158
tmp13065 := PrimEqual(W5158, symset)

if True == tmp13065 {
tmp13040 := MakeNative(func(__e *ControlFlow) {
W5159 := __e.Get(1)
_ = W5159
tmp13061 := PrimIsPair(W5159)

if True == tmp13061 {
tmp13041 := MakeNative(func(__e *ControlFlow) {
W5160 := __e.Get(1)
_ = W5160
tmp13042 := MakeNative(func(__e *ControlFlow) {
W5161 := __e.Get(1)
_ = W5161
tmp13056 := PrimIsPair(W5161)

if True == tmp13056 {
tmp13043 := MakeNative(func(__e *ControlFlow) {
W5162 := __e.Get(1)
_ = W5162
tmp13044 := MakeNative(func(__e *ControlFlow) {
W5163 := __e.Get(1)
_ = W5163
tmp13051 := PrimEqual(W5163, Nil)

if True == tmp13051 {
tmp13045 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13045

tmp13046 := MakeNative(func(__e *ControlFlow) {
tmp13047 := PrimCons(W5160, Nil)

tmp13048 := PrimCons(symvalue, tmp13047)

tmp13049 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5162, V4956, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp13048, V4956, V4957, V4958, V4959, W4962, tmp13049)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5160, symsymbol, V4957, V4958, V4959, W4962, tmp13046)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13052 := PrimTail(W5161)

tmp13053 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13052, V4958)


__e.TailApply(tmp13044, tmp13053)
return


}, 1)

tmp13054 := PrimHead(W5161)

__e.TailApply(tmp13043, tmp13054)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13057 := PrimTail(W5159)

tmp13058 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13057, V4958)


__e.TailApply(tmp13042, tmp13058)
return


}, 1)

tmp13059 := PrimHead(W5159)

__e.TailApply(tmp13041, tmp13059)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13062 := PrimTail(W5157)

tmp13063 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13062, V4958)


__e.TailApply(tmp13040, tmp13063)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13066 := PrimHead(W5157)

tmp13067 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13066, V4958)


__e.TailApply(tmp13039, tmp13067)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13070 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13071 := Call(__e, tmp13038, tmp13070)


ifres13037 = tmp13071


} else {
ifres13037 = False


}

__e.TailApply(tmp13010, ifres13037)
return


} else {
__e.Return(W5148)
return
}


}, 1)

tmp13111 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13075 Obj

if True == tmp13111 {
tmp13076 := MakeNative(func(__e *ControlFlow) {
W5149 := __e.Get(1)
_ = W5149
tmp13108 := PrimIsPair(W5149)

if True == tmp13108 {
tmp13077 := MakeNative(func(__e *ControlFlow) {
W5150 := __e.Get(1)
_ = W5150
tmp13104 := PrimEqual(W5150, syminput_7)

if True == tmp13104 {
tmp13078 := MakeNative(func(__e *ControlFlow) {
W5151 := __e.Get(1)
_ = W5151
tmp13100 := PrimIsPair(W5151)

if True == tmp13100 {
tmp13079 := MakeNative(func(__e *ControlFlow) {
W5152 := __e.Get(1)
_ = W5152
tmp13080 := MakeNative(func(__e *ControlFlow) {
W5153 := __e.Get(1)
_ = W5153
tmp13095 := PrimIsPair(W5153)

if True == tmp13095 {
tmp13081 := MakeNative(func(__e *ControlFlow) {
W5154 := __e.Get(1)
_ = W5154
tmp13082 := MakeNative(func(__e *ControlFlow) {
W5155 := __e.Get(1)
_ = W5155
tmp13090 := PrimEqual(W5155, Nil)

if True == tmp13090 {
tmp13083 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13083

tmp13084 := Call(__e, PrimFunc(symshen_4deref), W5152, V4958)


tmp13085 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp13084)


tmp13086 := MakeNative(func(__e *ControlFlow) {
tmp13087 := PrimCons(symin, Nil)

tmp13088 := PrimCons(symstream, tmp13087)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5154, tmp13088, V4957, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), V4956, tmp13085, V4958, V4959, W4962, tmp13086)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13091 := PrimTail(W5153)

tmp13092 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13091, V4958)


__e.TailApply(tmp13082, tmp13092)
return


}, 1)

tmp13093 := PrimHead(W5153)

__e.TailApply(tmp13081, tmp13093)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13096 := PrimTail(W5151)

tmp13097 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13096, V4958)


__e.TailApply(tmp13080, tmp13097)
return


}, 1)

tmp13098 := PrimHead(W5151)

__e.TailApply(tmp13079, tmp13098)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13101 := PrimTail(W5149)

tmp13102 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13101, V4958)


__e.TailApply(tmp13078, tmp13102)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13105 := PrimHead(W5149)

tmp13106 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13105, V4958)


__e.TailApply(tmp13077, tmp13106)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13109 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13110 := Call(__e, tmp13076, tmp13109)


ifres13075 = tmp13110


} else {
ifres13075 = False


}

__e.TailApply(tmp13009, ifres13075)
return


} else {
__e.Return(W5140)
return
}


}, 1)

tmp13149 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13114 Obj

if True == tmp13149 {
tmp13115 := MakeNative(func(__e *ControlFlow) {
W5141 := __e.Get(1)
_ = W5141
tmp13146 := PrimIsPair(W5141)

if True == tmp13146 {
tmp13116 := MakeNative(func(__e *ControlFlow) {
W5142 := __e.Get(1)
_ = W5142
tmp13142 := PrimEqual(W5142, symtype)

if True == tmp13142 {
tmp13117 := MakeNative(func(__e *ControlFlow) {
W5143 := __e.Get(1)
_ = W5143
tmp13138 := PrimIsPair(W5143)

if True == tmp13138 {
tmp13118 := MakeNative(func(__e *ControlFlow) {
W5144 := __e.Get(1)
_ = W5144
tmp13119 := MakeNative(func(__e *ControlFlow) {
W5145 := __e.Get(1)
_ = W5145
tmp13133 := PrimIsPair(W5145)

if True == tmp13133 {
tmp13120 := MakeNative(func(__e *ControlFlow) {
W5146 := __e.Get(1)
_ = W5146
tmp13121 := MakeNative(func(__e *ControlFlow) {
W5147 := __e.Get(1)
_ = W5147
tmp13128 := PrimEqual(W5147, Nil)

if True == tmp13128 {
tmp13122 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13122

tmp13123 := MakeNative(func(__e *ControlFlow) {
tmp13124 := Call(__e, PrimFunc(symshen_4deref), W5146, V4958)


tmp13125 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp13124)


tmp13126 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5144, V4956, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symis_b), tmp13125, V4956, V4958, V4959, W4962, tmp13126)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4958, V4959, W4962, tmp13123)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13129 := PrimTail(W5145)

tmp13130 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13129, V4958)


__e.TailApply(tmp13121, tmp13130)
return


}, 1)

tmp13131 := PrimHead(W5145)

__e.TailApply(tmp13120, tmp13131)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13134 := PrimTail(W5143)

tmp13135 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13134, V4958)


__e.TailApply(tmp13119, tmp13135)
return


}, 1)

tmp13136 := PrimHead(W5143)

__e.TailApply(tmp13118, tmp13136)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13139 := PrimTail(W5141)

tmp13140 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13139, V4958)


__e.TailApply(tmp13117, tmp13140)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13143 := PrimHead(W5141)

tmp13144 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13143, V4958)


__e.TailApply(tmp13116, tmp13144)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13147 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13148 := Call(__e, tmp13115, tmp13147)


ifres13114 = tmp13148


} else {
ifres13114 = False


}

__e.TailApply(tmp13008, ifres13114)
return


} else {
__e.Return(W5119)
return
}


}, 1)

tmp13237 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13152 Obj

if True == tmp13237 {
tmp13153 := MakeNative(func(__e *ControlFlow) {
W5120 := __e.Get(1)
_ = W5120
tmp13234 := PrimIsPair(W5120)

if True == tmp13234 {
tmp13154 := MakeNative(func(__e *ControlFlow) {
W5121 := __e.Get(1)
_ = W5121
tmp13230 := PrimEqual(W5121, symopen)

if True == tmp13230 {
tmp13155 := MakeNative(func(__e *ControlFlow) {
W5122 := __e.Get(1)
_ = W5122
tmp13226 := PrimIsPair(W5122)

if True == tmp13226 {
tmp13156 := MakeNative(func(__e *ControlFlow) {
W5123 := __e.Get(1)
_ = W5123
tmp13157 := MakeNative(func(__e *ControlFlow) {
W5124 := __e.Get(1)
_ = W5124
tmp13221 := PrimIsPair(W5124)

if True == tmp13221 {
tmp13158 := MakeNative(func(__e *ControlFlow) {
W5125 := __e.Get(1)
_ = W5125
tmp13159 := MakeNative(func(__e *ControlFlow) {
W5126 := __e.Get(1)
_ = W5126
tmp13216 := PrimEqual(W5126, Nil)

if True == tmp13216 {
tmp13160 := MakeNative(func(__e *ControlFlow) {
W5127 := __e.Get(1)
_ = W5127
tmp13161 := MakeNative(func(__e *ControlFlow) {
W5128 := __e.Get(1)
_ = W5128
tmp13205 := PrimIsPair(W5127)

if True == tmp13205 {
tmp13162 := MakeNative(func(__e *ControlFlow) {
W5130 := __e.Get(1)
_ = W5130
tmp13163 := MakeNative(func(__e *ControlFlow) {
W5131 := __e.Get(1)
_ = W5131
tmp13167 := PrimEqual(W5130, symstream)

if True == tmp13167 {
__e.TailApply(PrimFunc(symthaw), W5131)
return
} else {
tmp13165 := Call(__e, PrimFunc(symshen_4pvar_2), W5130)


if True == tmp13165 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5130, symstream, V4958, W5131)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13168 := MakeNative(func(__e *ControlFlow) {
tmp13169 := MakeNative(func(__e *ControlFlow) {
W5132 := __e.Get(1)
_ = W5132
tmp13170 := MakeNative(func(__e *ControlFlow) {
W5133 := __e.Get(1)
_ = W5133
tmp13190 := PrimIsPair(W5132)

if True == tmp13190 {
tmp13171 := MakeNative(func(__e *ControlFlow) {
W5135 := __e.Get(1)
_ = W5135
tmp13172 := MakeNative(func(__e *ControlFlow) {
W5136 := __e.Get(1)
_ = W5136
tmp13173 := MakeNative(func(__e *ControlFlow) {
W5137 := __e.Get(1)
_ = W5137
tmp13177 := PrimEqual(W5136, Nil)

if True == tmp13177 {
__e.TailApply(PrimFunc(symthaw), W5137)
return
} else {
tmp13175 := Call(__e, PrimFunc(symshen_4pvar_2), W5136)


if True == tmp13175 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5136, Nil, V4958, W5137)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13178 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5133, W5135)
return
}, 0)

__e.TailApply(tmp13173, tmp13178)
return


}, 1)

tmp13179 := PrimTail(W5132)

tmp13180 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13179, V4958)


__e.TailApply(tmp13172, tmp13180)
return


}, 1)

tmp13181 := PrimHead(W5132)

__e.TailApply(tmp13171, tmp13181)
return


} else {
tmp13188 := Call(__e, PrimFunc(symshen_4pvar_2), W5132)


if True == tmp13188 {
tmp13182 := MakeNative(func(__e *ControlFlow) {
W5138 := __e.Get(1)
_ = W5138
tmp13183 := PrimCons(W5138, Nil)

tmp13184 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5133, W5138)
return
}, 0)

tmp13185 := Call(__e, PrimFunc(symshen_4bind_b), W5132, tmp13183, V4958, tmp13184)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13185)
return


}, 1)

tmp13186 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13182, tmp13186)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13191 := MakeNative(func(__e *ControlFlow) {
Z5134 := __e.Get(1)
_ = Z5134
__e.TailApply(W5128, Z5134)
return
}, 1)

__e.TailApply(tmp13170, tmp13191)
return


}, 1)

tmp13192 := PrimTail(W5127)

tmp13193 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13192, V4958)


__e.TailApply(tmp13169, tmp13193)
return


}, 0)

__e.TailApply(tmp13163, tmp13168)
return


}, 1)

tmp13194 := PrimHead(W5127)

tmp13195 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13194, V4958)


__e.TailApply(tmp13162, tmp13195)
return


} else {
tmp13203 := Call(__e, PrimFunc(symshen_4pvar_2), W5127)


if True == tmp13203 {
tmp13196 := MakeNative(func(__e *ControlFlow) {
W5139 := __e.Get(1)
_ = W5139
tmp13197 := PrimCons(W5139, Nil)

tmp13198 := PrimCons(symstream, tmp13197)

tmp13199 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5128, W5139)
return
}, 0)

tmp13200 := Call(__e, PrimFunc(symshen_4bind_b), W5127, tmp13198, V4958, tmp13199)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13200)
return


}, 1)

tmp13201 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13196, tmp13201)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13206 := MakeNative(func(__e *ControlFlow) {
Z5129 := __e.Get(1)
_ = Z5129
tmp13207 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13207

tmp13208 := MakeNative(func(__e *ControlFlow) {
tmp13209 := Call(__e, PrimFunc(symshen_4lazyderef), Z5129, V4958)


tmp13210 := PrimCons(symout, Nil)

tmp13211 := PrimCons(symin, tmp13210)

tmp13212 := Call(__e, PrimFunc(symelement_2), tmp13209, tmp13211)


tmp13213 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5123, symstring, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp13212, V4958, V4959, W4962, tmp13213)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5125, Z5129, V4958, V4959, W4962, tmp13208)
return


}, 1)

__e.TailApply(tmp13161, tmp13206)
return


}, 1)

tmp13214 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13160, tmp13214)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13217 := PrimTail(W5124)

tmp13218 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13217, V4958)


__e.TailApply(tmp13159, tmp13218)
return


}, 1)

tmp13219 := PrimHead(W5124)

__e.TailApply(tmp13158, tmp13219)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13222 := PrimTail(W5122)

tmp13223 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13222, V4958)


__e.TailApply(tmp13157, tmp13223)
return


}, 1)

tmp13224 := PrimHead(W5122)

__e.TailApply(tmp13156, tmp13224)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13227 := PrimTail(W5120)

tmp13228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13227, V4958)


__e.TailApply(tmp13155, tmp13228)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13231 := PrimHead(W5120)

tmp13232 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13231, V4958)


__e.TailApply(tmp13154, tmp13232)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13235 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13236 := Call(__e, tmp13153, tmp13235)


ifres13152 = tmp13236


} else {
ifres13152 = False


}

__e.TailApply(tmp13007, ifres13152)
return


} else {
__e.Return(W5106)
return
}


}, 1)

tmp13301 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13240 Obj

if True == tmp13301 {
tmp13241 := MakeNative(func(__e *ControlFlow) {
W5107 := __e.Get(1)
_ = W5107
tmp13298 := PrimIsPair(W5107)

if True == tmp13298 {
tmp13242 := MakeNative(func(__e *ControlFlow) {
W5108 := __e.Get(1)
_ = W5108
tmp13294 := PrimEqual(W5108, symlet)

if True == tmp13294 {
tmp13243 := MakeNative(func(__e *ControlFlow) {
W5109 := __e.Get(1)
_ = W5109
tmp13290 := PrimIsPair(W5109)

if True == tmp13290 {
tmp13244 := MakeNative(func(__e *ControlFlow) {
W5110 := __e.Get(1)
_ = W5110
tmp13245 := MakeNative(func(__e *ControlFlow) {
W5111 := __e.Get(1)
_ = W5111
tmp13285 := PrimIsPair(W5111)

if True == tmp13285 {
tmp13246 := MakeNative(func(__e *ControlFlow) {
W5112 := __e.Get(1)
_ = W5112
tmp13247 := MakeNative(func(__e *ControlFlow) {
W5113 := __e.Get(1)
_ = W5113
tmp13280 := PrimIsPair(W5113)

if True == tmp13280 {
tmp13248 := MakeNative(func(__e *ControlFlow) {
W5114 := __e.Get(1)
_ = W5114
tmp13249 := MakeNative(func(__e *ControlFlow) {
W5115 := __e.Get(1)
_ = W5115
tmp13275 := PrimEqual(W5115, Nil)

if True == tmp13275 {
tmp13250 := MakeNative(func(__e *ControlFlow) {
W5116 := __e.Get(1)
_ = W5116
tmp13251 := MakeNative(func(__e *ControlFlow) {
W5117 := __e.Get(1)
_ = W5117
tmp13252 := MakeNative(func(__e *ControlFlow) {
W5118 := __e.Get(1)
_ = W5118
tmp13253 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13253

tmp13254 := MakeNative(func(__e *ControlFlow) {
tmp13255 := Call(__e, PrimFunc(symshen_4lazyderef), W5110, V4958)


tmp13256 := Call(__e, PrimFunc(symshen_4freshterm), tmp13255)


tmp13257 := MakeNative(func(__e *ControlFlow) {
tmp13258 := Call(__e, PrimFunc(symshen_4lazyderef), W5110, V4958)


tmp13259 := Call(__e, PrimFunc(symshen_4lazyderef), W5117, V4958)


tmp13260 := Call(__e, PrimFunc(symshen_4lazyderef), W5114, V4958)


tmp13261 := Call(__e, PrimFunc(symshen_4beta), tmp13258, tmp13259, tmp13260)


tmp13262 := MakeNative(func(__e *ControlFlow) {
tmp13263 := PrimIntern(MakeString(":"))

tmp13264 := PrimCons(W5118, Nil)

tmp13265 := PrimCons(tmp13263, tmp13264)

tmp13266 := PrimCons(W5117, tmp13265)

tmp13267 := PrimCons(tmp13266, V4957)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5116, V4956, tmp13267, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5116, tmp13261, V4958, V4959, W4962, tmp13262)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5117, tmp13256, V4958, V4959, W4962, tmp13257)
return


}, 0)

tmp13268 := Call(__e, PrimFunc(symshen_4system_1S_1h), W5112, W5118, V4957, V4958, V4959, W4962, tmp13254)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13268)
return


}, 1)

tmp13269 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13270 := Call(__e, tmp13252, tmp13269)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13270)
return


}, 1)

tmp13271 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13272 := Call(__e, tmp13251, tmp13271)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13272)
return


}, 1)

tmp13273 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13250, tmp13273)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13276 := PrimTail(W5113)

tmp13277 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13276, V4958)


__e.TailApply(tmp13249, tmp13277)
return


}, 1)

tmp13278 := PrimHead(W5113)

__e.TailApply(tmp13248, tmp13278)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13281 := PrimTail(W5111)

tmp13282 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13281, V4958)


__e.TailApply(tmp13247, tmp13282)
return


}, 1)

tmp13283 := PrimHead(W5111)

__e.TailApply(tmp13246, tmp13283)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13286 := PrimTail(W5109)

tmp13287 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13286, V4958)


__e.TailApply(tmp13245, tmp13287)
return


}, 1)

tmp13288 := PrimHead(W5109)

__e.TailApply(tmp13244, tmp13288)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13291 := PrimTail(W5107)

tmp13292 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13291, V4958)


__e.TailApply(tmp13243, tmp13292)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13295 := PrimHead(W5107)

tmp13296 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13295, V4958)


__e.TailApply(tmp13242, tmp13296)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13299 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13300 := Call(__e, tmp13241, tmp13299)


ifres13240 = tmp13300


} else {
ifres13240 = False


}

__e.TailApply(tmp13006, ifres13240)
return


} else {
__e.Return(W5076)
return
}


}, 1)

tmp13425 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13304 Obj

if True == tmp13425 {
tmp13305 := MakeNative(func(__e *ControlFlow) {
W5077 := __e.Get(1)
_ = W5077
tmp13422 := PrimIsPair(W5077)

if True == tmp13422 {
tmp13306 := MakeNative(func(__e *ControlFlow) {
W5078 := __e.Get(1)
_ = W5078
tmp13418 := PrimEqual(W5078, symlambda)

if True == tmp13418 {
tmp13307 := MakeNative(func(__e *ControlFlow) {
W5079 := __e.Get(1)
_ = W5079
tmp13414 := PrimIsPair(W5079)

if True == tmp13414 {
tmp13308 := MakeNative(func(__e *ControlFlow) {
W5080 := __e.Get(1)
_ = W5080
tmp13309 := MakeNative(func(__e *ControlFlow) {
W5081 := __e.Get(1)
_ = W5081
tmp13409 := PrimIsPair(W5081)

if True == tmp13409 {
tmp13310 := MakeNative(func(__e *ControlFlow) {
W5082 := __e.Get(1)
_ = W5082
tmp13311 := MakeNative(func(__e *ControlFlow) {
W5083 := __e.Get(1)
_ = W5083
tmp13404 := PrimEqual(W5083, Nil)

if True == tmp13404 {
tmp13312 := MakeNative(func(__e *ControlFlow) {
W5084 := __e.Get(1)
_ = W5084
tmp13313 := MakeNative(func(__e *ControlFlow) {
W5085 := __e.Get(1)
_ = W5085
tmp13380 := PrimIsPair(W5084)

if True == tmp13380 {
tmp13314 := MakeNative(func(__e *ControlFlow) {
W5090 := __e.Get(1)
_ = W5090
tmp13315 := MakeNative(func(__e *ControlFlow) {
W5091 := __e.Get(1)
_ = W5091
tmp13316 := MakeNative(func(__e *ControlFlow) {
W5092 := __e.Get(1)
_ = W5092
tmp13360 := PrimIsPair(W5091)

if True == tmp13360 {
tmp13317 := MakeNative(func(__e *ControlFlow) {
W5094 := __e.Get(1)
_ = W5094
tmp13318 := MakeNative(func(__e *ControlFlow) {
W5095 := __e.Get(1)
_ = W5095
tmp13322 := PrimEqual(W5094, sym_1_1_6)

if True == tmp13322 {
__e.TailApply(PrimFunc(symthaw), W5095)
return
} else {
tmp13320 := Call(__e, PrimFunc(symshen_4pvar_2), W5094)


if True == tmp13320 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5094, sym_1_1_6, V4958, W5095)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13323 := MakeNative(func(__e *ControlFlow) {
tmp13324 := MakeNative(func(__e *ControlFlow) {
W5096 := __e.Get(1)
_ = W5096
tmp13325 := MakeNative(func(__e *ControlFlow) {
W5097 := __e.Get(1)
_ = W5097
tmp13345 := PrimIsPair(W5096)

if True == tmp13345 {
tmp13326 := MakeNative(func(__e *ControlFlow) {
W5099 := __e.Get(1)
_ = W5099
tmp13327 := MakeNative(func(__e *ControlFlow) {
W5100 := __e.Get(1)
_ = W5100
tmp13328 := MakeNative(func(__e *ControlFlow) {
W5101 := __e.Get(1)
_ = W5101
tmp13332 := PrimEqual(W5100, Nil)

if True == tmp13332 {
__e.TailApply(PrimFunc(symthaw), W5101)
return
} else {
tmp13330 := Call(__e, PrimFunc(symshen_4pvar_2), W5100)


if True == tmp13330 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5100, Nil, V4958, W5101)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13333 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5097, W5099)
return
}, 0)

__e.TailApply(tmp13328, tmp13333)
return


}, 1)

tmp13334 := PrimTail(W5096)

tmp13335 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13334, V4958)


__e.TailApply(tmp13327, tmp13335)
return


}, 1)

tmp13336 := PrimHead(W5096)

__e.TailApply(tmp13326, tmp13336)
return


} else {
tmp13343 := Call(__e, PrimFunc(symshen_4pvar_2), W5096)


if True == tmp13343 {
tmp13337 := MakeNative(func(__e *ControlFlow) {
W5102 := __e.Get(1)
_ = W5102
tmp13338 := PrimCons(W5102, Nil)

tmp13339 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5097, W5102)
return
}, 0)

tmp13340 := Call(__e, PrimFunc(symshen_4bind_b), W5096, tmp13338, V4958, tmp13339)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13340)
return


}, 1)

tmp13341 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13337, tmp13341)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13346 := MakeNative(func(__e *ControlFlow) {
Z5098 := __e.Get(1)
_ = Z5098
__e.TailApply(W5092, Z5098)
return
}, 1)

__e.TailApply(tmp13325, tmp13346)
return


}, 1)

tmp13347 := PrimTail(W5091)

tmp13348 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13347, V4958)


__e.TailApply(tmp13324, tmp13348)
return


}, 0)

__e.TailApply(tmp13318, tmp13323)
return


}, 1)

tmp13349 := PrimHead(W5091)

tmp13350 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13349, V4958)


__e.TailApply(tmp13317, tmp13350)
return


} else {
tmp13358 := Call(__e, PrimFunc(symshen_4pvar_2), W5091)


if True == tmp13358 {
tmp13351 := MakeNative(func(__e *ControlFlow) {
W5103 := __e.Get(1)
_ = W5103
tmp13352 := PrimCons(W5103, Nil)

tmp13353 := PrimCons(sym_1_1_6, tmp13352)

tmp13354 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5092, W5103)
return
}, 0)

tmp13355 := Call(__e, PrimFunc(symshen_4bind_b), W5091, tmp13353, V4958, tmp13354)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13355)
return


}, 1)

tmp13356 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13351, tmp13356)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13361 := MakeNative(func(__e *ControlFlow) {
Z5093 := __e.Get(1)
_ = Z5093
tmp13362 := Call(__e, W5085, W5090)


__e.TailApply(tmp13362, Z5093)
return


}, 1)

__e.TailApply(tmp13316, tmp13361)
return


}, 1)

tmp13363 := PrimTail(W5084)

tmp13364 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13363, V4958)


__e.TailApply(tmp13315, tmp13364)
return


}, 1)

tmp13365 := PrimHead(W5084)

__e.TailApply(tmp13314, tmp13365)
return


} else {
tmp13378 := Call(__e, PrimFunc(symshen_4pvar_2), W5084)


if True == tmp13378 {
tmp13366 := MakeNative(func(__e *ControlFlow) {
W5104 := __e.Get(1)
_ = W5104
tmp13367 := MakeNative(func(__e *ControlFlow) {
W5105 := __e.Get(1)
_ = W5105
tmp13368 := PrimCons(W5105, Nil)

tmp13369 := PrimCons(sym_1_1_6, tmp13368)

tmp13370 := PrimCons(W5104, tmp13369)

tmp13371 := MakeNative(func(__e *ControlFlow) {
tmp13372 := Call(__e, W5085, W5104)


__e.TailApply(tmp13372, W5105)
return


}, 0)

tmp13373 := Call(__e, PrimFunc(symshen_4bind_b), W5084, tmp13370, V4958, tmp13371)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13373)
return


}, 1)

tmp13374 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13375 := Call(__e, tmp13367, tmp13374)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13375)
return


}, 1)

tmp13376 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13366, tmp13376)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13381 := MakeNative(func(__e *ControlFlow) {
Z5086 := __e.Get(1)
_ = Z5086
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5087 := __e.Get(1)
_ = Z5087
tmp13382 := MakeNative(func(__e *ControlFlow) {
W5088 := __e.Get(1)
_ = W5088
tmp13383 := MakeNative(func(__e *ControlFlow) {
W5089 := __e.Get(1)
_ = W5089
tmp13384 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13384

tmp13385 := Call(__e, PrimFunc(symshen_4lazyderef), W5080, V4958)


tmp13386 := Call(__e, PrimFunc(symshen_4freshterm), tmp13385)


tmp13387 := MakeNative(func(__e *ControlFlow) {
tmp13388 := Call(__e, PrimFunc(symshen_4lazyderef), W5080, V4958)


tmp13389 := Call(__e, PrimFunc(symshen_4deref), W5089, V4958)


tmp13390 := Call(__e, PrimFunc(symshen_4deref), W5082, V4958)


tmp13391 := Call(__e, PrimFunc(symshen_4beta), tmp13388, tmp13389, tmp13390)


tmp13392 := MakeNative(func(__e *ControlFlow) {
tmp13393 := PrimIntern(MakeString(":"))

tmp13394 := PrimCons(Z5086, Nil)

tmp13395 := PrimCons(tmp13393, tmp13394)

tmp13396 := PrimCons(W5089, tmp13395)

tmp13397 := PrimCons(tmp13396, V4957)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5088, Z5087, tmp13397, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5088, tmp13391, V4958, V4959, W4962, tmp13392)
return


}, 0)

tmp13398 := Call(__e, PrimFunc(symbind), W5089, tmp13386, V4958, V4959, W4962, tmp13387)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13398)
return


}, 1)

tmp13399 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13400 := Call(__e, tmp13383, tmp13399)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13400)
return


}, 1)

tmp13401 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13382, tmp13401)
return


}, 1))
return
}, 1)

__e.TailApply(tmp13313, tmp13381)
return


}, 1)

tmp13402 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13312, tmp13402)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13405 := PrimTail(W5081)

tmp13406 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13405, V4958)


__e.TailApply(tmp13311, tmp13406)
return


}, 1)

tmp13407 := PrimHead(W5081)

__e.TailApply(tmp13310, tmp13407)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13410 := PrimTail(W5079)

tmp13411 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13410, V4958)


__e.TailApply(tmp13309, tmp13411)
return


}, 1)

tmp13412 := PrimHead(W5079)

__e.TailApply(tmp13308, tmp13412)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13415 := PrimTail(W5077)

tmp13416 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13415, V4958)


__e.TailApply(tmp13307, tmp13416)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13419 := PrimHead(W5077)

tmp13420 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13419, V4958)


__e.TailApply(tmp13306, tmp13420)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13423 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13424 := Call(__e, tmp13305, tmp13423)


ifres13304 = tmp13424


} else {
ifres13304 = False


}

__e.TailApply(tmp13005, ifres13304)
return


} else {
__e.Return(W5066)
return
}


}, 1)

tmp13468 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13428 Obj

if True == tmp13468 {
tmp13429 := MakeNative(func(__e *ControlFlow) {
W5067 := __e.Get(1)
_ = W5067
tmp13465 := PrimIsPair(W5067)

if True == tmp13465 {
tmp13430 := MakeNative(func(__e *ControlFlow) {
W5068 := __e.Get(1)
_ = W5068
tmp13461 := PrimEqual(W5068, sym_8s)

if True == tmp13461 {
tmp13431 := MakeNative(func(__e *ControlFlow) {
W5069 := __e.Get(1)
_ = W5069
tmp13457 := PrimIsPair(W5069)

if True == tmp13457 {
tmp13432 := MakeNative(func(__e *ControlFlow) {
W5070 := __e.Get(1)
_ = W5070
tmp13433 := MakeNative(func(__e *ControlFlow) {
W5071 := __e.Get(1)
_ = W5071
tmp13452 := PrimIsPair(W5071)

if True == tmp13452 {
tmp13434 := MakeNative(func(__e *ControlFlow) {
W5072 := __e.Get(1)
_ = W5072
tmp13435 := MakeNative(func(__e *ControlFlow) {
W5073 := __e.Get(1)
_ = W5073
tmp13447 := PrimEqual(W5073, Nil)

if True == tmp13447 {
tmp13436 := MakeNative(func(__e *ControlFlow) {
W5074 := __e.Get(1)
_ = W5074
tmp13437 := MakeNative(func(__e *ControlFlow) {
W5075 := __e.Get(1)
_ = W5075
tmp13441 := PrimEqual(W5074, symstring)

if True == tmp13441 {
__e.TailApply(PrimFunc(symthaw), W5075)
return
} else {
tmp13439 := Call(__e, PrimFunc(symshen_4pvar_2), W5074)


if True == tmp13439 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5074, symstring, V4958, W5075)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13442 := MakeNative(func(__e *ControlFlow) {
tmp13443 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13443

tmp13444 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5072, symstring, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5070, symstring, V4957, V4958, V4959, W4962, tmp13444)
return


}, 0)

__e.TailApply(tmp13437, tmp13442)
return


}, 1)

tmp13445 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13436, tmp13445)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13448 := PrimTail(W5071)

tmp13449 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13448, V4958)


__e.TailApply(tmp13435, tmp13449)
return


}, 1)

tmp13450 := PrimHead(W5071)

__e.TailApply(tmp13434, tmp13450)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13453 := PrimTail(W5069)

tmp13454 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13453, V4958)


__e.TailApply(tmp13433, tmp13454)
return


}, 1)

tmp13455 := PrimHead(W5069)

__e.TailApply(tmp13432, tmp13455)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13458 := PrimTail(W5067)

tmp13459 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13458, V4958)


__e.TailApply(tmp13431, tmp13459)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13462 := PrimHead(W5067)

tmp13463 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13462, V4958)


__e.TailApply(tmp13430, tmp13463)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13466 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13467 := Call(__e, tmp13429, tmp13466)


ifres13428 = tmp13467


} else {
ifres13428 = False


}

__e.TailApply(tmp13004, ifres13428)
return


} else {
__e.Return(W5045)
return
}


}, 1)

tmp13553 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13471 Obj

if True == tmp13553 {
tmp13472 := MakeNative(func(__e *ControlFlow) {
W5046 := __e.Get(1)
_ = W5046
tmp13550 := PrimIsPair(W5046)

if True == tmp13550 {
tmp13473 := MakeNative(func(__e *ControlFlow) {
W5047 := __e.Get(1)
_ = W5047
tmp13546 := PrimEqual(W5047, sym_8v)

if True == tmp13546 {
tmp13474 := MakeNative(func(__e *ControlFlow) {
W5048 := __e.Get(1)
_ = W5048
tmp13542 := PrimIsPair(W5048)

if True == tmp13542 {
tmp13475 := MakeNative(func(__e *ControlFlow) {
W5049 := __e.Get(1)
_ = W5049
tmp13476 := MakeNative(func(__e *ControlFlow) {
W5050 := __e.Get(1)
_ = W5050
tmp13537 := PrimIsPair(W5050)

if True == tmp13537 {
tmp13477 := MakeNative(func(__e *ControlFlow) {
W5051 := __e.Get(1)
_ = W5051
tmp13478 := MakeNative(func(__e *ControlFlow) {
W5052 := __e.Get(1)
_ = W5052
tmp13532 := PrimEqual(W5052, Nil)

if True == tmp13532 {
tmp13479 := MakeNative(func(__e *ControlFlow) {
W5053 := __e.Get(1)
_ = W5053
tmp13480 := MakeNative(func(__e *ControlFlow) {
W5054 := __e.Get(1)
_ = W5054
tmp13524 := PrimIsPair(W5053)

if True == tmp13524 {
tmp13481 := MakeNative(func(__e *ControlFlow) {
W5056 := __e.Get(1)
_ = W5056
tmp13482 := MakeNative(func(__e *ControlFlow) {
W5057 := __e.Get(1)
_ = W5057
tmp13486 := PrimEqual(W5056, symvector)

if True == tmp13486 {
__e.TailApply(PrimFunc(symthaw), W5057)
return
} else {
tmp13484 := Call(__e, PrimFunc(symshen_4pvar_2), W5056)


if True == tmp13484 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5056, symvector, V4958, W5057)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13487 := MakeNative(func(__e *ControlFlow) {
tmp13488 := MakeNative(func(__e *ControlFlow) {
W5058 := __e.Get(1)
_ = W5058
tmp13489 := MakeNative(func(__e *ControlFlow) {
W5059 := __e.Get(1)
_ = W5059
tmp13509 := PrimIsPair(W5058)

if True == tmp13509 {
tmp13490 := MakeNative(func(__e *ControlFlow) {
W5061 := __e.Get(1)
_ = W5061
tmp13491 := MakeNative(func(__e *ControlFlow) {
W5062 := __e.Get(1)
_ = W5062
tmp13492 := MakeNative(func(__e *ControlFlow) {
W5063 := __e.Get(1)
_ = W5063
tmp13496 := PrimEqual(W5062, Nil)

if True == tmp13496 {
__e.TailApply(PrimFunc(symthaw), W5063)
return
} else {
tmp13494 := Call(__e, PrimFunc(symshen_4pvar_2), W5062)


if True == tmp13494 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5062, Nil, V4958, W5063)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13497 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5059, W5061)
return
}, 0)

__e.TailApply(tmp13492, tmp13497)
return


}, 1)

tmp13498 := PrimTail(W5058)

tmp13499 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13498, V4958)


__e.TailApply(tmp13491, tmp13499)
return


}, 1)

tmp13500 := PrimHead(W5058)

__e.TailApply(tmp13490, tmp13500)
return


} else {
tmp13507 := Call(__e, PrimFunc(symshen_4pvar_2), W5058)


if True == tmp13507 {
tmp13501 := MakeNative(func(__e *ControlFlow) {
W5064 := __e.Get(1)
_ = W5064
tmp13502 := PrimCons(W5064, Nil)

tmp13503 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5059, W5064)
return
}, 0)

tmp13504 := Call(__e, PrimFunc(symshen_4bind_b), W5058, tmp13502, V4958, tmp13503)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13504)
return


}, 1)

tmp13505 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13501, tmp13505)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13510 := MakeNative(func(__e *ControlFlow) {
Z5060 := __e.Get(1)
_ = Z5060
__e.TailApply(W5054, Z5060)
return
}, 1)

__e.TailApply(tmp13489, tmp13510)
return


}, 1)

tmp13511 := PrimTail(W5053)

tmp13512 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13511, V4958)


__e.TailApply(tmp13488, tmp13512)
return


}, 0)

__e.TailApply(tmp13482, tmp13487)
return


}, 1)

tmp13513 := PrimHead(W5053)

tmp13514 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13513, V4958)


__e.TailApply(tmp13481, tmp13514)
return


} else {
tmp13522 := Call(__e, PrimFunc(symshen_4pvar_2), W5053)


if True == tmp13522 {
tmp13515 := MakeNative(func(__e *ControlFlow) {
W5065 := __e.Get(1)
_ = W5065
tmp13516 := PrimCons(W5065, Nil)

tmp13517 := PrimCons(symvector, tmp13516)

tmp13518 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5054, W5065)
return
}, 0)

tmp13519 := Call(__e, PrimFunc(symshen_4bind_b), W5053, tmp13517, V4958, tmp13518)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13519)
return


}, 1)

tmp13520 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13515, tmp13520)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13525 := MakeNative(func(__e *ControlFlow) {
Z5055 := __e.Get(1)
_ = Z5055
tmp13526 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13526

tmp13527 := MakeNative(func(__e *ControlFlow) {
tmp13528 := PrimCons(Z5055, Nil)

tmp13529 := PrimCons(symvector, tmp13528)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5051, tmp13529, V4957, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5049, Z5055, V4957, V4958, V4959, W4962, tmp13527)
return


}, 1)

__e.TailApply(tmp13480, tmp13525)
return


}, 1)

tmp13530 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13479, tmp13530)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13533 := PrimTail(W5050)

tmp13534 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13533, V4958)


__e.TailApply(tmp13478, tmp13534)
return


}, 1)

tmp13535 := PrimHead(W5050)

__e.TailApply(tmp13477, tmp13535)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13538 := PrimTail(W5048)

tmp13539 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13538, V4958)


__e.TailApply(tmp13476, tmp13539)
return


}, 1)

tmp13540 := PrimHead(W5048)

__e.TailApply(tmp13475, tmp13540)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13543 := PrimTail(W5046)

tmp13544 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13543, V4958)


__e.TailApply(tmp13474, tmp13544)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13547 := PrimHead(W5046)

tmp13548 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13547, V4958)


__e.TailApply(tmp13473, tmp13548)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13551 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13552 := Call(__e, tmp13472, tmp13551)


ifres13471 = tmp13552


} else {
ifres13471 = False


}

__e.TailApply(tmp13003, ifres13471)
return


} else {
__e.Return(W5017)
return
}


}, 1)

tmp13659 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13556 Obj

if True == tmp13659 {
tmp13557 := MakeNative(func(__e *ControlFlow) {
W5018 := __e.Get(1)
_ = W5018
tmp13656 := PrimIsPair(W5018)

if True == tmp13656 {
tmp13558 := MakeNative(func(__e *ControlFlow) {
W5019 := __e.Get(1)
_ = W5019
tmp13652 := PrimEqual(W5019, sym_8p)

if True == tmp13652 {
tmp13559 := MakeNative(func(__e *ControlFlow) {
W5020 := __e.Get(1)
_ = W5020
tmp13648 := PrimIsPair(W5020)

if True == tmp13648 {
tmp13560 := MakeNative(func(__e *ControlFlow) {
W5021 := __e.Get(1)
_ = W5021
tmp13561 := MakeNative(func(__e *ControlFlow) {
W5022 := __e.Get(1)
_ = W5022
tmp13643 := PrimIsPair(W5022)

if True == tmp13643 {
tmp13562 := MakeNative(func(__e *ControlFlow) {
W5023 := __e.Get(1)
_ = W5023
tmp13563 := MakeNative(func(__e *ControlFlow) {
W5024 := __e.Get(1)
_ = W5024
tmp13638 := PrimEqual(W5024, Nil)

if True == tmp13638 {
tmp13564 := MakeNative(func(__e *ControlFlow) {
W5025 := __e.Get(1)
_ = W5025
tmp13565 := MakeNative(func(__e *ControlFlow) {
W5026 := __e.Get(1)
_ = W5026
tmp13632 := PrimIsPair(W5025)

if True == tmp13632 {
tmp13566 := MakeNative(func(__e *ControlFlow) {
W5029 := __e.Get(1)
_ = W5029
tmp13567 := MakeNative(func(__e *ControlFlow) {
W5030 := __e.Get(1)
_ = W5030
tmp13568 := MakeNative(func(__e *ControlFlow) {
W5031 := __e.Get(1)
_ = W5031
tmp13612 := PrimIsPair(W5030)

if True == tmp13612 {
tmp13569 := MakeNative(func(__e *ControlFlow) {
W5033 := __e.Get(1)
_ = W5033
tmp13570 := MakeNative(func(__e *ControlFlow) {
W5034 := __e.Get(1)
_ = W5034
tmp13574 := PrimEqual(W5033, sym_d)

if True == tmp13574 {
__e.TailApply(PrimFunc(symthaw), W5034)
return
} else {
tmp13572 := Call(__e, PrimFunc(symshen_4pvar_2), W5033)


if True == tmp13572 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5033, sym_d, V4958, W5034)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13575 := MakeNative(func(__e *ControlFlow) {
tmp13576 := MakeNative(func(__e *ControlFlow) {
W5035 := __e.Get(1)
_ = W5035
tmp13577 := MakeNative(func(__e *ControlFlow) {
W5036 := __e.Get(1)
_ = W5036
tmp13597 := PrimIsPair(W5035)

if True == tmp13597 {
tmp13578 := MakeNative(func(__e *ControlFlow) {
W5038 := __e.Get(1)
_ = W5038
tmp13579 := MakeNative(func(__e *ControlFlow) {
W5039 := __e.Get(1)
_ = W5039
tmp13580 := MakeNative(func(__e *ControlFlow) {
W5040 := __e.Get(1)
_ = W5040
tmp13584 := PrimEqual(W5039, Nil)

if True == tmp13584 {
__e.TailApply(PrimFunc(symthaw), W5040)
return
} else {
tmp13582 := Call(__e, PrimFunc(symshen_4pvar_2), W5039)


if True == tmp13582 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5039, Nil, V4958, W5040)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13585 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5036, W5038)
return
}, 0)

__e.TailApply(tmp13580, tmp13585)
return


}, 1)

tmp13586 := PrimTail(W5035)

tmp13587 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13586, V4958)


__e.TailApply(tmp13579, tmp13587)
return


}, 1)

tmp13588 := PrimHead(W5035)

__e.TailApply(tmp13578, tmp13588)
return


} else {
tmp13595 := Call(__e, PrimFunc(symshen_4pvar_2), W5035)


if True == tmp13595 {
tmp13589 := MakeNative(func(__e *ControlFlow) {
W5041 := __e.Get(1)
_ = W5041
tmp13590 := PrimCons(W5041, Nil)

tmp13591 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5036, W5041)
return
}, 0)

tmp13592 := Call(__e, PrimFunc(symshen_4bind_b), W5035, tmp13590, V4958, tmp13591)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13592)
return


}, 1)

tmp13593 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13589, tmp13593)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13598 := MakeNative(func(__e *ControlFlow) {
Z5037 := __e.Get(1)
_ = Z5037
__e.TailApply(W5031, Z5037)
return
}, 1)

__e.TailApply(tmp13577, tmp13598)
return


}, 1)

tmp13599 := PrimTail(W5030)

tmp13600 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13599, V4958)


__e.TailApply(tmp13576, tmp13600)
return


}, 0)

__e.TailApply(tmp13570, tmp13575)
return


}, 1)

tmp13601 := PrimHead(W5030)

tmp13602 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13601, V4958)


__e.TailApply(tmp13569, tmp13602)
return


} else {
tmp13610 := Call(__e, PrimFunc(symshen_4pvar_2), W5030)


if True == tmp13610 {
tmp13603 := MakeNative(func(__e *ControlFlow) {
W5042 := __e.Get(1)
_ = W5042
tmp13604 := PrimCons(W5042, Nil)

tmp13605 := PrimCons(sym_d, tmp13604)

tmp13606 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5031, W5042)
return
}, 0)

tmp13607 := Call(__e, PrimFunc(symshen_4bind_b), W5030, tmp13605, V4958, tmp13606)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13607)
return


}, 1)

tmp13608 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13603, tmp13608)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13613 := MakeNative(func(__e *ControlFlow) {
Z5032 := __e.Get(1)
_ = Z5032
tmp13614 := Call(__e, W5026, W5029)


__e.TailApply(tmp13614, Z5032)
return


}, 1)

__e.TailApply(tmp13568, tmp13613)
return


}, 1)

tmp13615 := PrimTail(W5025)

tmp13616 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13615, V4958)


__e.TailApply(tmp13567, tmp13616)
return


}, 1)

tmp13617 := PrimHead(W5025)

__e.TailApply(tmp13566, tmp13617)
return


} else {
tmp13630 := Call(__e, PrimFunc(symshen_4pvar_2), W5025)


if True == tmp13630 {
tmp13618 := MakeNative(func(__e *ControlFlow) {
W5043 := __e.Get(1)
_ = W5043
tmp13619 := MakeNative(func(__e *ControlFlow) {
W5044 := __e.Get(1)
_ = W5044
tmp13620 := PrimCons(W5044, Nil)

tmp13621 := PrimCons(sym_d, tmp13620)

tmp13622 := PrimCons(W5043, tmp13621)

tmp13623 := MakeNative(func(__e *ControlFlow) {
tmp13624 := Call(__e, W5026, W5043)


__e.TailApply(tmp13624, W5044)
return


}, 0)

tmp13625 := Call(__e, PrimFunc(symshen_4bind_b), W5025, tmp13622, V4958, tmp13623)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13625)
return


}, 1)

tmp13626 := Call(__e, PrimFunc(symshen_4newpv), V4958)


tmp13627 := Call(__e, tmp13619, tmp13626)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13627)
return


}, 1)

tmp13628 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13618, tmp13628)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13633 := MakeNative(func(__e *ControlFlow) {
Z5027 := __e.Get(1)
_ = Z5027
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5028 := __e.Get(1)
_ = Z5028
tmp13634 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13634

tmp13635 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5023, Z5028, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5021, Z5027, V4957, V4958, V4959, W4962, tmp13635)
return


}, 1))
return
}, 1)

__e.TailApply(tmp13565, tmp13633)
return


}, 1)

tmp13636 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13564, tmp13636)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13639 := PrimTail(W5022)

tmp13640 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13639, V4958)


__e.TailApply(tmp13563, tmp13640)
return


}, 1)

tmp13641 := PrimHead(W5022)

__e.TailApply(tmp13562, tmp13641)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13644 := PrimTail(W5020)

tmp13645 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13644, V4958)


__e.TailApply(tmp13561, tmp13645)
return


}, 1)

tmp13646 := PrimHead(W5020)

__e.TailApply(tmp13560, tmp13646)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13649 := PrimTail(W5018)

tmp13650 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13649, V4958)


__e.TailApply(tmp13559, tmp13650)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13653 := PrimHead(W5018)

tmp13654 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13653, V4958)


__e.TailApply(tmp13558, tmp13654)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13657 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13658 := Call(__e, tmp13557, tmp13657)


ifres13556 = tmp13658


} else {
ifres13556 = False


}

__e.TailApply(tmp13002, ifres13556)
return


} else {
__e.Return(W4996)
return
}


}, 1)

tmp13744 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13662 Obj

if True == tmp13744 {
tmp13663 := MakeNative(func(__e *ControlFlow) {
W4997 := __e.Get(1)
_ = W4997
tmp13741 := PrimIsPair(W4997)

if True == tmp13741 {
tmp13664 := MakeNative(func(__e *ControlFlow) {
W4998 := __e.Get(1)
_ = W4998
tmp13737 := PrimEqual(W4998, symcons)

if True == tmp13737 {
tmp13665 := MakeNative(func(__e *ControlFlow) {
W4999 := __e.Get(1)
_ = W4999
tmp13733 := PrimIsPair(W4999)

if True == tmp13733 {
tmp13666 := MakeNative(func(__e *ControlFlow) {
W5000 := __e.Get(1)
_ = W5000
tmp13667 := MakeNative(func(__e *ControlFlow) {
W5001 := __e.Get(1)
_ = W5001
tmp13728 := PrimIsPair(W5001)

if True == tmp13728 {
tmp13668 := MakeNative(func(__e *ControlFlow) {
W5002 := __e.Get(1)
_ = W5002
tmp13669 := MakeNative(func(__e *ControlFlow) {
W5003 := __e.Get(1)
_ = W5003
tmp13723 := PrimEqual(W5003, Nil)

if True == tmp13723 {
tmp13670 := MakeNative(func(__e *ControlFlow) {
W5004 := __e.Get(1)
_ = W5004
tmp13671 := MakeNative(func(__e *ControlFlow) {
W5005 := __e.Get(1)
_ = W5005
tmp13715 := PrimIsPair(W5004)

if True == tmp13715 {
tmp13672 := MakeNative(func(__e *ControlFlow) {
W5007 := __e.Get(1)
_ = W5007
tmp13673 := MakeNative(func(__e *ControlFlow) {
W5008 := __e.Get(1)
_ = W5008
tmp13677 := PrimEqual(W5007, symlist)

if True == tmp13677 {
__e.TailApply(PrimFunc(symthaw), W5008)
return
} else {
tmp13675 := Call(__e, PrimFunc(symshen_4pvar_2), W5007)


if True == tmp13675 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5007, symlist, V4958, W5008)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13678 := MakeNative(func(__e *ControlFlow) {
tmp13679 := MakeNative(func(__e *ControlFlow) {
W5009 := __e.Get(1)
_ = W5009
tmp13680 := MakeNative(func(__e *ControlFlow) {
W5010 := __e.Get(1)
_ = W5010
tmp13700 := PrimIsPair(W5009)

if True == tmp13700 {
tmp13681 := MakeNative(func(__e *ControlFlow) {
W5012 := __e.Get(1)
_ = W5012
tmp13682 := MakeNative(func(__e *ControlFlow) {
W5013 := __e.Get(1)
_ = W5013
tmp13683 := MakeNative(func(__e *ControlFlow) {
W5014 := __e.Get(1)
_ = W5014
tmp13687 := PrimEqual(W5013, Nil)

if True == tmp13687 {
__e.TailApply(PrimFunc(symthaw), W5014)
return
} else {
tmp13685 := Call(__e, PrimFunc(symshen_4pvar_2), W5013)


if True == tmp13685 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5013, Nil, V4958, W5014)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13688 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5010, W5012)
return
}, 0)

__e.TailApply(tmp13683, tmp13688)
return


}, 1)

tmp13689 := PrimTail(W5009)

tmp13690 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13689, V4958)


__e.TailApply(tmp13682, tmp13690)
return


}, 1)

tmp13691 := PrimHead(W5009)

__e.TailApply(tmp13681, tmp13691)
return


} else {
tmp13698 := Call(__e, PrimFunc(symshen_4pvar_2), W5009)


if True == tmp13698 {
tmp13692 := MakeNative(func(__e *ControlFlow) {
W5015 := __e.Get(1)
_ = W5015
tmp13693 := PrimCons(W5015, Nil)

tmp13694 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5010, W5015)
return
}, 0)

tmp13695 := Call(__e, PrimFunc(symshen_4bind_b), W5009, tmp13693, V4958, tmp13694)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13695)
return


}, 1)

tmp13696 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13692, tmp13696)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13701 := MakeNative(func(__e *ControlFlow) {
Z5011 := __e.Get(1)
_ = Z5011
__e.TailApply(W5005, Z5011)
return
}, 1)

__e.TailApply(tmp13680, tmp13701)
return


}, 1)

tmp13702 := PrimTail(W5004)

tmp13703 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13702, V4958)


__e.TailApply(tmp13679, tmp13703)
return


}, 0)

__e.TailApply(tmp13673, tmp13678)
return


}, 1)

tmp13704 := PrimHead(W5004)

tmp13705 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13704, V4958)


__e.TailApply(tmp13672, tmp13705)
return


} else {
tmp13713 := Call(__e, PrimFunc(symshen_4pvar_2), W5004)


if True == tmp13713 {
tmp13706 := MakeNative(func(__e *ControlFlow) {
W5016 := __e.Get(1)
_ = W5016
tmp13707 := PrimCons(W5016, Nil)

tmp13708 := PrimCons(symlist, tmp13707)

tmp13709 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5005, W5016)
return
}, 0)

tmp13710 := Call(__e, PrimFunc(symshen_4bind_b), W5004, tmp13708, V4958, tmp13709)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13710)
return


}, 1)

tmp13711 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13706, tmp13711)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13716 := MakeNative(func(__e *ControlFlow) {
Z5006 := __e.Get(1)
_ = Z5006
tmp13717 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13717

tmp13718 := MakeNative(func(__e *ControlFlow) {
tmp13719 := PrimCons(Z5006, Nil)

tmp13720 := PrimCons(symlist, tmp13719)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5002, tmp13720, V4957, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5000, Z5006, V4957, V4958, V4959, W4962, tmp13718)
return


}, 1)

__e.TailApply(tmp13671, tmp13716)
return


}, 1)

tmp13721 := Call(__e, PrimFunc(symshen_4lazyderef), V4956, V4958)


__e.TailApply(tmp13670, tmp13721)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13724 := PrimTail(W5001)

tmp13725 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13724, V4958)


__e.TailApply(tmp13669, tmp13725)
return


}, 1)

tmp13726 := PrimHead(W5001)

__e.TailApply(tmp13668, tmp13726)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13729 := PrimTail(W4999)

tmp13730 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13729, V4958)


__e.TailApply(tmp13667, tmp13730)
return


}, 1)

tmp13731 := PrimHead(W4999)

__e.TailApply(tmp13666, tmp13731)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13734 := PrimTail(W4997)

tmp13735 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13734, V4958)


__e.TailApply(tmp13665, tmp13735)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13738 := PrimHead(W4997)

tmp13739 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13738, V4958)


__e.TailApply(tmp13664, tmp13739)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13742 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13743 := Call(__e, tmp13663, tmp13742)


ifres13662 = tmp13743


} else {
ifres13662 = False


}

__e.TailApply(tmp13001, ifres13662)
return


} else {
__e.Return(W4989)
return
}


}, 1)

tmp13775 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13747 Obj

if True == tmp13775 {
tmp13748 := MakeNative(func(__e *ControlFlow) {
W4990 := __e.Get(1)
_ = W4990
tmp13772 := PrimIsPair(W4990)

if True == tmp13772 {
tmp13749 := MakeNative(func(__e *ControlFlow) {
W4991 := __e.Get(1)
_ = W4991
tmp13750 := MakeNative(func(__e *ControlFlow) {
W4992 := __e.Get(1)
_ = W4992
tmp13767 := PrimIsPair(W4992)

if True == tmp13767 {
tmp13751 := MakeNative(func(__e *ControlFlow) {
W4993 := __e.Get(1)
_ = W4993
tmp13752 := MakeNative(func(__e *ControlFlow) {
W4994 := __e.Get(1)
_ = W4994
tmp13762 := PrimEqual(W4994, Nil)

if True == tmp13762 {
tmp13753 := MakeNative(func(__e *ControlFlow) {
W4995 := __e.Get(1)
_ = W4995
tmp13754 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13754

tmp13755 := PrimCons(V4956, Nil)

tmp13756 := PrimCons(sym_1_1_6, tmp13755)

tmp13757 := PrimCons(W4995, tmp13756)

tmp13758 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4993, W4995, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

tmp13759 := Call(__e, PrimFunc(symshen_4system_1S_1h), W4991, tmp13757, V4957, V4958, V4959, W4962, tmp13758)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13759)
return


}, 1)

tmp13760 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13753, tmp13760)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13763 := PrimTail(W4992)

tmp13764 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13763, V4958)


__e.TailApply(tmp13752, tmp13764)
return


}, 1)

tmp13765 := PrimHead(W4992)

__e.TailApply(tmp13751, tmp13765)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13768 := PrimTail(W4990)

tmp13769 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13768, V4958)


__e.TailApply(tmp13750, tmp13769)
return


}, 1)

tmp13770 := PrimHead(W4990)

__e.TailApply(tmp13749, tmp13770)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13773 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13774 := Call(__e, tmp13748, tmp13773)


ifres13747 = tmp13774


} else {
ifres13747 = False


}

__e.TailApply(tmp13000, ifres13747)
return


} else {
__e.Return(W4982)
return
}


}, 1)

tmp13810 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13778 Obj

if True == tmp13810 {
tmp13779 := MakeNative(func(__e *ControlFlow) {
W4983 := __e.Get(1)
_ = W4983
tmp13807 := PrimIsPair(W4983)

if True == tmp13807 {
tmp13780 := MakeNative(func(__e *ControlFlow) {
W4984 := __e.Get(1)
_ = W4984
tmp13781 := MakeNative(func(__e *ControlFlow) {
W4985 := __e.Get(1)
_ = W4985
tmp13802 := PrimIsPair(W4985)

if True == tmp13802 {
tmp13782 := MakeNative(func(__e *ControlFlow) {
W4986 := __e.Get(1)
_ = W4986
tmp13783 := MakeNative(func(__e *ControlFlow) {
W4987 := __e.Get(1)
_ = W4987
tmp13797 := PrimEqual(W4987, Nil)

if True == tmp13797 {
tmp13784 := MakeNative(func(__e *ControlFlow) {
W4988 := __e.Get(1)
_ = W4988
tmp13785 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13785

tmp13786 := Call(__e, PrimFunc(symshen_4lazyderef), W4984, V4958)


tmp13787 := PrimIsPair(tmp13786)

tmp13788 := PrimNot(tmp13787)

tmp13789 := MakeNative(func(__e *ControlFlow) {
tmp13790 := PrimCons(V4956, Nil)

tmp13791 := PrimCons(sym_1_1_6, tmp13790)

tmp13792 := PrimCons(W4988, tmp13791)

tmp13793 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4system_1S_1h), W4986, W4988, V4957, V4958, V4959, W4962, V4961)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4984, tmp13792, V4958, V4959, W4962, tmp13793)
return


}, 0)

tmp13794 := Call(__e, PrimFunc(symwhen), tmp13788, V4958, V4959, W4962, tmp13789)


__e.TailApply(PrimFunc(symshen_4gc), V4958, tmp13794)
return


}, 1)

tmp13795 := Call(__e, PrimFunc(symshen_4newpv), V4958)


__e.TailApply(tmp13784, tmp13795)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13798 := PrimTail(W4985)

tmp13799 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13798, V4958)


__e.TailApply(tmp13783, tmp13799)
return


}, 1)

tmp13800 := PrimHead(W4985)

__e.TailApply(tmp13782, tmp13800)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13803 := PrimTail(W4983)

tmp13804 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13803, V4958)


__e.TailApply(tmp13781, tmp13804)
return


}, 1)

tmp13805 := PrimHead(W4983)

__e.TailApply(tmp13780, tmp13805)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13808 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13809 := Call(__e, tmp13779, tmp13808)


ifres13778 = tmp13809


} else {
ifres13778 = False


}

__e.TailApply(tmp12999, ifres13778)
return


} else {
__e.Return(W4976)
return
}


}, 1)

tmp13837 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13813 Obj

if True == tmp13837 {
tmp13814 := MakeNative(func(__e *ControlFlow) {
W4977 := __e.Get(1)
_ = W4977
tmp13834 := PrimIsPair(W4977)

if True == tmp13834 {
tmp13815 := MakeNative(func(__e *ControlFlow) {
W4978 := __e.Get(1)
_ = W4978
tmp13830 := PrimEqual(W4978, symfn)

if True == tmp13830 {
tmp13816 := MakeNative(func(__e *ControlFlow) {
W4979 := __e.Get(1)
_ = W4979
tmp13826 := PrimIsPair(W4979)

if True == tmp13826 {
tmp13817 := MakeNative(func(__e *ControlFlow) {
W4980 := __e.Get(1)
_ = W4980
tmp13818 := MakeNative(func(__e *ControlFlow) {
W4981 := __e.Get(1)
_ = W4981
tmp13821 := PrimEqual(W4981, Nil)

if True == tmp13821 {
tmp13819 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13819

__e.TailApply(PrimFunc(symshen_4lookupsig), W4980, V4956, V4958, V4959, W4962, V4961)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13822 := PrimTail(W4979)

tmp13823 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13822, V4958)


__e.TailApply(tmp13818, tmp13823)
return


}, 1)

tmp13824 := PrimHead(W4979)

__e.TailApply(tmp13817, tmp13824)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13827 := PrimTail(W4977)

tmp13828 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13827, V4958)


__e.TailApply(tmp13816, tmp13828)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13831 := PrimHead(W4977)

tmp13832 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13831, V4958)


__e.TailApply(tmp13815, tmp13832)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13835 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13836 := Call(__e, tmp13814, tmp13835)


ifres13813 = tmp13836


} else {
ifres13813 = False


}

__e.TailApply(tmp12998, ifres13813)
return


} else {
__e.Return(W4970)
return
}


}, 1)

tmp13870 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13840 Obj

if True == tmp13870 {
tmp13841 := MakeNative(func(__e *ControlFlow) {
W4971 := __e.Get(1)
_ = W4971
tmp13867 := PrimIsPair(W4971)

if True == tmp13867 {
tmp13842 := MakeNative(func(__e *ControlFlow) {
W4972 := __e.Get(1)
_ = W4972
tmp13863 := PrimEqual(W4972, symfn)

if True == tmp13863 {
tmp13843 := MakeNative(func(__e *ControlFlow) {
W4973 := __e.Get(1)
_ = W4973
tmp13859 := PrimIsPair(W4973)

if True == tmp13859 {
tmp13844 := MakeNative(func(__e *ControlFlow) {
W4974 := __e.Get(1)
_ = W4974
tmp13845 := MakeNative(func(__e *ControlFlow) {
W4975 := __e.Get(1)
_ = W4975
tmp13854 := PrimEqual(W4975, Nil)

if True == tmp13854 {
tmp13846 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13846

tmp13847 := Call(__e, PrimFunc(symshen_4deref), W4974, V4958)


tmp13848 := Call(__e, PrimFunc(symarity), tmp13847)


tmp13849 := PrimEqual(tmp13848, MakeNumber(0))

tmp13850 := MakeNative(func(__e *ControlFlow) {
tmp13851 := MakeNative(func(__e *ControlFlow) {
tmp13852 := PrimCons(W4974, Nil)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), tmp13852, V4956, V4957, V4958, V4959, W4962, V4961)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V4958, V4959, W4962, tmp13851)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp13849, V4958, V4959, W4962, tmp13850)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13855 := PrimTail(W4973)

tmp13856 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13855, V4958)


__e.TailApply(tmp13845, tmp13856)
return


}, 1)

tmp13857 := PrimHead(W4973)

__e.TailApply(tmp13844, tmp13857)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13860 := PrimTail(W4971)

tmp13861 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13860, V4958)


__e.TailApply(tmp13843, tmp13861)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13864 := PrimHead(W4971)

tmp13865 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13864, V4958)


__e.TailApply(tmp13842, tmp13865)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13868 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13869 := Call(__e, tmp13841, tmp13868)


ifres13840 = tmp13869


} else {
ifres13840 = False


}

__e.TailApply(tmp12997, ifres13840)
return


} else {
__e.Return(W4966)
return
}


}, 1)

tmp13889 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13873 Obj

if True == tmp13889 {
tmp13874 := MakeNative(func(__e *ControlFlow) {
W4967 := __e.Get(1)
_ = W4967
tmp13886 := PrimIsPair(W4967)

if True == tmp13886 {
tmp13875 := MakeNative(func(__e *ControlFlow) {
W4968 := __e.Get(1)
_ = W4968
tmp13876 := MakeNative(func(__e *ControlFlow) {
W4969 := __e.Get(1)
_ = W4969
tmp13881 := PrimEqual(W4969, Nil)

if True == tmp13881 {
tmp13877 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13877

tmp13878 := PrimCons(V4956, Nil)

tmp13879 := PrimCons(sym_1_1_6, tmp13878)

__e.TailApply(PrimFunc(symshen_4lookupsig), W4968, tmp13879, V4958, V4959, W4962, V4961)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13882 := PrimTail(W4967)

tmp13883 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13882, V4958)


__e.TailApply(tmp13876, tmp13883)
return


}, 1)

tmp13884 := PrimHead(W4967)

__e.TailApply(tmp13875, tmp13884)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13887 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13888 := Call(__e, tmp13874, tmp13887)


ifres13873 = tmp13888


} else {
ifres13873 = False


}

__e.TailApply(tmp12996, ifres13873)
return


} else {
__e.Return(W4965)
return
}


}, 1)

tmp13895 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13892 Obj

if True == tmp13895 {
tmp13893 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13893

tmp13894 := Call(__e, PrimFunc(symshen_4by_1hypothesis), V4955, V4956, V4957, V4958, V4959, W4962, V4961)


ifres13892 = tmp13894


} else {
ifres13892 = False


}

__e.TailApply(tmp12995, ifres13892)
return


} else {
__e.Return(W4964)
return
}


}, 1)

tmp13905 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13898 Obj

if True == tmp13905 {
tmp13899 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13899

tmp13900 := Call(__e, PrimFunc(symshen_4lazyderef), V4955, V4958)


tmp13901 := PrimIsPair(tmp13900)

tmp13902 := PrimNot(tmp13901)

tmp13903 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4primitive), V4955, V4956, V4958, V4959, W4962, V4961)
return
}, 0)

tmp13904 := Call(__e, PrimFunc(symwhen), tmp13902, V4958, V4959, W4962, tmp13903)


ifres13898 = tmp13904


} else {
ifres13898 = False


}

__e.TailApply(tmp12994, ifres13898)
return


} else {
__e.Return(W4963)
return
}


}, 1)

tmp13917 := Call(__e, PrimFunc(symshen_4unlocked_2), V4959)


var ifres13908 Obj

if True == tmp13917 {
tmp13909 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13909

tmp13910 := PrimValue(symshen_4_dspy_d)

tmp13911 := MakeNative(func(__e *ControlFlow) {
tmp13912 := PrimIntern(MakeString(":"))

tmp13913 := PrimCons(V4956, Nil)

tmp13914 := PrimCons(tmp13912, tmp13913)

tmp13915 := PrimCons(V4955, tmp13914)

__e.TailApply(PrimFunc(symshen_4show), tmp13915, V4957, V4958, V4959, W4962, V4961)
return


}, 0)

tmp13916 := Call(__e, PrimFunc(symwhen), tmp13910, V4958, V4959, W4962, tmp13911)


ifres13908 = tmp13916


} else {
ifres13908 = False


}

__e.TailApply(tmp12993, ifres13908)
return


}, 1)

tmp13918 := PrimNumberAdd(V4960, MakeNumber(1))

__e.TailApply(tmp12992, tmp13918)
return


}, 7)

tmp13919 := Call(__e, ns2_1set, symshen_4system_1S_1h, tmp12991)


_ = tmp13919

tmp13920 := MakeNative(func(__e *ControlFlow) {
V5167 := __e.Get(1)
_ = V5167
V5168 := __e.Get(2)
_ = V5168
V5169 := __e.Get(3)
_ = V5169
V5170 := __e.Get(4)
_ = V5170
V5171 := __e.Get(5)
_ = V5171
V5172 := __e.Get(6)
_ = V5172
tmp13921 := MakeNative(func(__e *ControlFlow) {
W5173 := __e.Get(1)
_ = W5173
tmp14029 := PrimEqual(W5173, False)

if True == tmp14029 {
tmp13922 := MakeNative(func(__e *ControlFlow) {
W5176 := __e.Get(1)
_ = W5176
tmp14013 := PrimEqual(W5176, False)

if True == tmp14013 {
tmp13923 := MakeNative(func(__e *ControlFlow) {
W5179 := __e.Get(1)
_ = W5179
tmp13997 := PrimEqual(W5179, False)

if True == tmp13997 {
tmp13924 := MakeNative(func(__e *ControlFlow) {
W5182 := __e.Get(1)
_ = W5182
tmp13981 := PrimEqual(W5182, False)

if True == tmp13981 {
tmp13979 := Call(__e, PrimFunc(symshen_4unlocked_2), V5170)


if True == tmp13979 {
tmp13925 := MakeNative(func(__e *ControlFlow) {
W5185 := __e.Get(1)
_ = W5185
tmp13976 := PrimEqual(W5185, Nil)

if True == tmp13976 {
tmp13926 := MakeNative(func(__e *ControlFlow) {
W5186 := __e.Get(1)
_ = W5186
tmp13927 := MakeNative(func(__e *ControlFlow) {
W5187 := __e.Get(1)
_ = W5187
tmp13971 := PrimIsPair(W5186)

if True == tmp13971 {
tmp13928 := MakeNative(func(__e *ControlFlow) {
W5189 := __e.Get(1)
_ = W5189
tmp13929 := MakeNative(func(__e *ControlFlow) {
W5190 := __e.Get(1)
_ = W5190
tmp13933 := PrimEqual(W5189, symlist)

if True == tmp13933 {
__e.TailApply(PrimFunc(symthaw), W5190)
return
} else {
tmp13931 := Call(__e, PrimFunc(symshen_4pvar_2), W5189)


if True == tmp13931 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5189, symlist, V5169, W5190)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13934 := MakeNative(func(__e *ControlFlow) {
tmp13935 := MakeNative(func(__e *ControlFlow) {
W5191 := __e.Get(1)
_ = W5191
tmp13936 := MakeNative(func(__e *ControlFlow) {
W5192 := __e.Get(1)
_ = W5192
tmp13956 := PrimIsPair(W5191)

if True == tmp13956 {
tmp13937 := MakeNative(func(__e *ControlFlow) {
W5194 := __e.Get(1)
_ = W5194
tmp13938 := MakeNative(func(__e *ControlFlow) {
W5195 := __e.Get(1)
_ = W5195
tmp13939 := MakeNative(func(__e *ControlFlow) {
W5196 := __e.Get(1)
_ = W5196
tmp13943 := PrimEqual(W5195, Nil)

if True == tmp13943 {
__e.TailApply(PrimFunc(symthaw), W5196)
return
} else {
tmp13941 := Call(__e, PrimFunc(symshen_4pvar_2), W5195)


if True == tmp13941 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5195, Nil, V5169, W5196)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13944 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5192, W5194)
return
}, 0)

__e.TailApply(tmp13939, tmp13944)
return


}, 1)

tmp13945 := PrimTail(W5191)

tmp13946 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13945, V5169)


__e.TailApply(tmp13938, tmp13946)
return


}, 1)

tmp13947 := PrimHead(W5191)

__e.TailApply(tmp13937, tmp13947)
return


} else {
tmp13954 := Call(__e, PrimFunc(symshen_4pvar_2), W5191)


if True == tmp13954 {
tmp13948 := MakeNative(func(__e *ControlFlow) {
W5197 := __e.Get(1)
_ = W5197
tmp13949 := PrimCons(W5197, Nil)

tmp13950 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5192, W5197)
return
}, 0)

tmp13951 := Call(__e, PrimFunc(symshen_4bind_b), W5191, tmp13949, V5169, tmp13950)


__e.TailApply(PrimFunc(symshen_4gc), V5169, tmp13951)
return


}, 1)

tmp13952 := Call(__e, PrimFunc(symshen_4newpv), V5169)


__e.TailApply(tmp13948, tmp13952)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13957 := MakeNative(func(__e *ControlFlow) {
Z5193 := __e.Get(1)
_ = Z5193
__e.TailApply(W5187, Z5193)
return
}, 1)

__e.TailApply(tmp13936, tmp13957)
return


}, 1)

tmp13958 := PrimTail(W5186)

tmp13959 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13958, V5169)


__e.TailApply(tmp13935, tmp13959)
return


}, 0)

__e.TailApply(tmp13929, tmp13934)
return


}, 1)

tmp13960 := PrimHead(W5186)

tmp13961 := Call(__e, PrimFunc(symshen_4lazyderef), tmp13960, V5169)


__e.TailApply(tmp13928, tmp13961)
return


} else {
tmp13969 := Call(__e, PrimFunc(symshen_4pvar_2), W5186)


if True == tmp13969 {
tmp13962 := MakeNative(func(__e *ControlFlow) {
W5198 := __e.Get(1)
_ = W5198
tmp13963 := PrimCons(W5198, Nil)

tmp13964 := PrimCons(symlist, tmp13963)

tmp13965 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5187, W5198)
return
}, 0)

tmp13966 := Call(__e, PrimFunc(symshen_4bind_b), W5186, tmp13964, V5169, tmp13965)


__e.TailApply(PrimFunc(symshen_4gc), V5169, tmp13966)
return


}, 1)

tmp13967 := Call(__e, PrimFunc(symshen_4newpv), V5169)


__e.TailApply(tmp13962, tmp13967)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp13972 := MakeNative(func(__e *ControlFlow) {
Z5188 := __e.Get(1)
_ = Z5188
tmp13973 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13973

__e.TailApply(PrimFunc(symthaw), V5172)
return


}, 1)

__e.TailApply(tmp13927, tmp13972)
return


}, 1)

tmp13974 := Call(__e, PrimFunc(symshen_4lazyderef), V5168, V5169)


__e.TailApply(tmp13926, tmp13974)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp13977 := Call(__e, PrimFunc(symshen_4lazyderef), V5167, V5169)


__e.TailApply(tmp13925, tmp13977)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5182)
return
}


}, 1)

tmp13995 := Call(__e, PrimFunc(symshen_4unlocked_2), V5170)


var ifres13982 Obj

if True == tmp13995 {
tmp13983 := MakeNative(func(__e *ControlFlow) {
W5183 := __e.Get(1)
_ = W5183
tmp13984 := MakeNative(func(__e *ControlFlow) {
W5184 := __e.Get(1)
_ = W5184
tmp13988 := PrimEqual(W5183, symsymbol)

if True == tmp13988 {
__e.TailApply(PrimFunc(symthaw), W5184)
return
} else {
tmp13986 := Call(__e, PrimFunc(symshen_4pvar_2), W5183)


if True == tmp13986 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5183, symsymbol, V5169, W5184)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp13989 := MakeNative(func(__e *ControlFlow) {
tmp13990 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp13990

tmp13991 := Call(__e, PrimFunc(symshen_4lazyderef), V5167, V5169)


tmp13992 := PrimIsSymbol(tmp13991)

__e.TailApply(PrimFunc(symwhen), tmp13992, V5169, V5170, V5171, V5172)
return


}, 0)

__e.TailApply(tmp13984, tmp13989)
return


}, 1)

tmp13993 := Call(__e, PrimFunc(symshen_4lazyderef), V5168, V5169)


tmp13994 := Call(__e, tmp13983, tmp13993)


ifres13982 = tmp13994


} else {
ifres13982 = False


}

__e.TailApply(tmp13924, ifres13982)
return


} else {
__e.Return(W5179)
return
}


}, 1)

tmp14011 := Call(__e, PrimFunc(symshen_4unlocked_2), V5170)


var ifres13998 Obj

if True == tmp14011 {
tmp13999 := MakeNative(func(__e *ControlFlow) {
W5180 := __e.Get(1)
_ = W5180
tmp14000 := MakeNative(func(__e *ControlFlow) {
W5181 := __e.Get(1)
_ = W5181
tmp14004 := PrimEqual(W5180, symstring)

if True == tmp14004 {
__e.TailApply(PrimFunc(symthaw), W5181)
return
} else {
tmp14002 := Call(__e, PrimFunc(symshen_4pvar_2), W5180)


if True == tmp14002 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5180, symstring, V5169, W5181)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14005 := MakeNative(func(__e *ControlFlow) {
tmp14006 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14006

tmp14007 := Call(__e, PrimFunc(symshen_4lazyderef), V5167, V5169)


tmp14008 := PrimIsString(tmp14007)

__e.TailApply(PrimFunc(symwhen), tmp14008, V5169, V5170, V5171, V5172)
return


}, 0)

__e.TailApply(tmp14000, tmp14005)
return


}, 1)

tmp14009 := Call(__e, PrimFunc(symshen_4lazyderef), V5168, V5169)


tmp14010 := Call(__e, tmp13999, tmp14009)


ifres13998 = tmp14010


} else {
ifres13998 = False


}

__e.TailApply(tmp13923, ifres13998)
return


} else {
__e.Return(W5176)
return
}


}, 1)

tmp14027 := Call(__e, PrimFunc(symshen_4unlocked_2), V5170)


var ifres14014 Obj

if True == tmp14027 {
tmp14015 := MakeNative(func(__e *ControlFlow) {
W5177 := __e.Get(1)
_ = W5177
tmp14016 := MakeNative(func(__e *ControlFlow) {
W5178 := __e.Get(1)
_ = W5178
tmp14020 := PrimEqual(W5177, symboolean)

if True == tmp14020 {
__e.TailApply(PrimFunc(symthaw), W5178)
return
} else {
tmp14018 := Call(__e, PrimFunc(symshen_4pvar_2), W5177)


if True == tmp14018 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5177, symboolean, V5169, W5178)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14021 := MakeNative(func(__e *ControlFlow) {
tmp14022 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14022

tmp14023 := Call(__e, PrimFunc(symshen_4lazyderef), V5167, V5169)


tmp14024 := Call(__e, PrimFunc(symboolean_2), tmp14023)


__e.TailApply(PrimFunc(symwhen), tmp14024, V5169, V5170, V5171, V5172)
return


}, 0)

__e.TailApply(tmp14016, tmp14021)
return


}, 1)

tmp14025 := Call(__e, PrimFunc(symshen_4lazyderef), V5168, V5169)


tmp14026 := Call(__e, tmp14015, tmp14025)


ifres14014 = tmp14026


} else {
ifres14014 = False


}

__e.TailApply(tmp13922, ifres14014)
return


} else {
__e.Return(W5173)
return
}


}, 1)

tmp14043 := Call(__e, PrimFunc(symshen_4unlocked_2), V5170)


var ifres14030 Obj

if True == tmp14043 {
tmp14031 := MakeNative(func(__e *ControlFlow) {
W5174 := __e.Get(1)
_ = W5174
tmp14032 := MakeNative(func(__e *ControlFlow) {
W5175 := __e.Get(1)
_ = W5175
tmp14036 := PrimEqual(W5174, symnumber)

if True == tmp14036 {
__e.TailApply(PrimFunc(symthaw), W5175)
return
} else {
tmp14034 := Call(__e, PrimFunc(symshen_4pvar_2), W5174)


if True == tmp14034 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5174, symnumber, V5169, W5175)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp14037 := MakeNative(func(__e *ControlFlow) {
tmp14038 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14038

tmp14039 := Call(__e, PrimFunc(symshen_4lazyderef), V5167, V5169)


tmp14040 := PrimIsNumber(tmp14039)

__e.TailApply(PrimFunc(symwhen), tmp14040, V5169, V5170, V5171, V5172)
return


}, 0)

__e.TailApply(tmp14032, tmp14037)
return


}, 1)

tmp14041 := Call(__e, PrimFunc(symshen_4lazyderef), V5168, V5169)


tmp14042 := Call(__e, tmp14031, tmp14041)


ifres14030 = tmp14042


} else {
ifres14030 = False


}

__e.TailApply(tmp13921, ifres14030)
return


}, 6)

tmp14044 := Call(__e, ns2_1set, symshen_4primitive, tmp13920)


_ = tmp14044

tmp14045 := MakeNative(func(__e *ControlFlow) {
V5199 := __e.Get(1)
_ = V5199
V5200 := __e.Get(2)
_ = V5200
V5201 := __e.Get(3)
_ = V5201
V5202 := __e.Get(4)
_ = V5202
V5203 := __e.Get(5)
_ = V5203
V5204 := __e.Get(6)
_ = V5204
V5205 := __e.Get(7)
_ = V5205
tmp14046 := MakeNative(func(__e *ControlFlow) {
W5206 := __e.Get(1)
_ = W5206
tmp14057 := PrimEqual(W5206, False)

if True == tmp14057 {
tmp14055 := Call(__e, PrimFunc(symshen_4unlocked_2), V5203)


if True == tmp14055 {
tmp14047 := MakeNative(func(__e *ControlFlow) {
W5215 := __e.Get(1)
_ = W5215
tmp14052 := PrimIsPair(W5215)

if True == tmp14052 {
tmp14048 := MakeNative(func(__e *ControlFlow) {
W5216 := __e.Get(1)
_ = W5216
tmp14049 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14049

__e.TailApply(PrimFunc(symshen_4by_1hypothesis), V5199, V5200, W5216, V5202, V5203, V5204, V5205)
return


}, 1)

tmp14050 := PrimTail(W5215)

__e.TailApply(tmp14048, tmp14050)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14053 := Call(__e, PrimFunc(symshen_4lazyderef), V5201, V5202)


__e.TailApply(tmp14047, tmp14053)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5206)
return
}


}, 1)

tmp14099 := Call(__e, PrimFunc(symshen_4unlocked_2), V5203)


var ifres14058 Obj

if True == tmp14099 {
tmp14059 := MakeNative(func(__e *ControlFlow) {
W5207 := __e.Get(1)
_ = W5207
tmp14096 := PrimIsPair(W5207)

if True == tmp14096 {
tmp14060 := MakeNative(func(__e *ControlFlow) {
W5208 := __e.Get(1)
_ = W5208
tmp14092 := PrimIsPair(W5208)

if True == tmp14092 {
tmp14061 := MakeNative(func(__e *ControlFlow) {
W5209 := __e.Get(1)
_ = W5209
tmp14062 := MakeNative(func(__e *ControlFlow) {
W5210 := __e.Get(1)
_ = W5210
tmp14087 := PrimIsPair(W5210)

if True == tmp14087 {
tmp14063 := MakeNative(func(__e *ControlFlow) {
W5211 := __e.Get(1)
_ = W5211
tmp14064 := MakeNative(func(__e *ControlFlow) {
W5212 := __e.Get(1)
_ = W5212
tmp14082 := PrimIsPair(W5212)

if True == tmp14082 {
tmp14065 := MakeNative(func(__e *ControlFlow) {
W5213 := __e.Get(1)
_ = W5213
tmp14066 := MakeNative(func(__e *ControlFlow) {
W5214 := __e.Get(1)
_ = W5214
tmp14077 := PrimEqual(W5214, Nil)

if True == tmp14077 {
tmp14067 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14067

tmp14068 := Call(__e, PrimFunc(symshen_4deref), W5211, V5202)


tmp14069 := PrimIntern(MakeString(":"))

tmp14070 := PrimEqual(tmp14068, tmp14069)

tmp14071 := MakeNative(func(__e *ControlFlow) {
tmp14072 := Call(__e, PrimFunc(symshen_4deref), V5199, V5202)


tmp14073 := Call(__e, PrimFunc(symshen_4deref), W5209, V5202)


tmp14074 := PrimEqual(tmp14072, tmp14073)

tmp14075 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis_b), V5200, W5213, V5202, V5203, V5204, V5205)
return
}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14074, V5202, V5203, V5204, tmp14075)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14070, V5202, V5203, V5204, tmp14071)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14078 := PrimTail(W5212)

tmp14079 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14078, V5202)


__e.TailApply(tmp14066, tmp14079)
return


}, 1)

tmp14080 := PrimHead(W5212)

__e.TailApply(tmp14065, tmp14080)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14083 := PrimTail(W5210)

tmp14084 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14083, V5202)


__e.TailApply(tmp14064, tmp14084)
return


}, 1)

tmp14085 := PrimHead(W5210)

__e.TailApply(tmp14063, tmp14085)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14088 := PrimTail(W5208)

tmp14089 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14088, V5202)


__e.TailApply(tmp14062, tmp14089)
return


}, 1)

tmp14090 := PrimHead(W5208)

__e.TailApply(tmp14061, tmp14090)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14093 := PrimHead(W5207)

tmp14094 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14093, V5202)


__e.TailApply(tmp14060, tmp14094)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14097 := Call(__e, PrimFunc(symshen_4lazyderef), V5201, V5202)


tmp14098 := Call(__e, tmp14059, tmp14097)


ifres14058 = tmp14098


} else {
ifres14058 = False


}

__e.TailApply(tmp14046, ifres14058)
return


}, 7)

tmp14100 := Call(__e, ns2_1set, symshen_4by_1hypothesis, tmp14045)


_ = tmp14100

tmp14101 := MakeNative(func(__e *ControlFlow) {
V5217 := __e.Get(1)
_ = V5217
V5218 := __e.Get(2)
_ = V5218
V5219 := __e.Get(3)
_ = V5219
V5220 := __e.Get(4)
_ = V5220
V5221 := __e.Get(5)
_ = V5221
V5222 := __e.Get(6)
_ = V5222
tmp14106 := Call(__e, PrimFunc(symshen_4unlocked_2), V5220)


if True == tmp14106 {
tmp14102 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14102

tmp14103 := PrimValue(symshen_4_dsigf_d)

tmp14104 := Call(__e, PrimFunc(symassoc), V5217, tmp14103)


__e.TailApply(PrimFunc(symshen_4sigf), tmp14104, V5218, V5219, V5220, V5221, V5222)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp14107 := Call(__e, ns2_1set, symshen_4lookupsig, tmp14101)


_ = tmp14107

tmp14108 := MakeNative(func(__e *ControlFlow) {
V5237 := __e.Get(1)
_ = V5237
V5238 := __e.Get(2)
_ = V5238
V5239 := __e.Get(3)
_ = V5239
V5240 := __e.Get(4)
_ = V5240
V5241 := __e.Get(5)
_ = V5241
V5242 := __e.Get(6)
_ = V5242
tmp14115 := PrimIsPair(V5237)

if True == tmp14115 {
tmp14109 := PrimTail(V5237)

tmp14110 := Call(__e, tmp14109, V5238)


tmp14111 := Call(__e, tmp14110, V5239)


tmp14112 := Call(__e, tmp14111, V5240)


tmp14113 := Call(__e, tmp14112, V5241)


__e.TailApply(tmp14113, V5242)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp14116 := Call(__e, ns2_1set, symshen_4sigf, tmp14108)


_ = tmp14116

tmp14117 := MakeNative(func(__e *ControlFlow) {
V5243 := __e.Get(1)
_ = V5243
tmp14118 := MakeNative(func(__e *ControlFlow) {
W5244 := __e.Get(1)
_ = W5244
tmp14119 := MakeNative(func(__e *ControlFlow) {
W5245 := __e.Get(1)
_ = W5245
tmp14120 := MakeNative(func(__e *ControlFlow) {
W5246 := __e.Get(1)
_ = W5246
tmp14121 := MakeNative(func(__e *ControlFlow) {
W5247 := __e.Get(1)
_ = W5247
__e.Return(W5247)
return
}, 1)

tmp14122 := PrimValue(symshen_4_dgensym_d)

tmp14123 := PrimNumberAdd(MakeNumber(1), tmp14122)

tmp14124 := PrimSet(symshen_4_dgensym_d, tmp14123)

tmp14125 := PrimVectorSet(W5246, MakeNumber(2), tmp14124)

__e.TailApply(tmp14121, tmp14125)
return


}, 1)

tmp14126 := PrimVectorSet(W5245, MakeNumber(1), V5243)

__e.TailApply(tmp14120, tmp14126)
return


}, 1)

tmp14127 := PrimVectorSet(W5244, MakeNumber(0), symshen_4print_1freshterm)

__e.TailApply(tmp14119, tmp14127)
return


}, 1)

tmp14128 := PrimAbsvector(MakeNumber(3))

__e.TailApply(tmp14118, tmp14128)
return


}, 1)

tmp14129 := Call(__e, ns2_1set, symshen_4freshterm, tmp14117)


_ = tmp14129

tmp14130 := MakeNative(func(__e *ControlFlow) {
V5248 := __e.Get(1)
_ = V5248
tmp14131 := PrimVectorGet(V5248, MakeNumber(1))

tmp14132 := PrimStr(tmp14131)

__e.Return(PrimStringConcat(MakeString("&&"), tmp14132))
return


}, 1)

tmp14133 := Call(__e, ns2_1set, symshen_4print_1freshterm, tmp14130)


_ = tmp14133

tmp14134 := MakeNative(func(__e *ControlFlow) {
V5249 := __e.Get(1)
_ = V5249
V5250 := __e.Get(2)
_ = V5250
V5251 := __e.Get(3)
_ = V5251
V5252 := __e.Get(4)
_ = V5252
V5253 := __e.Get(5)
_ = V5253
V5254 := __e.Get(6)
_ = V5254
V5255 := __e.Get(7)
_ = V5255
tmp14135 := MakeNative(func(__e *ControlFlow) {
W5256 := __e.Get(1)
_ = W5256
tmp14146 := PrimEqual(W5256, False)

if True == tmp14146 {
tmp14144 := Call(__e, PrimFunc(symshen_4unlocked_2), V5253)


if True == tmp14144 {
tmp14136 := MakeNative(func(__e *ControlFlow) {
W5260 := __e.Get(1)
_ = W5260
tmp14141 := PrimIsPair(W5260)

if True == tmp14141 {
tmp14137 := MakeNative(func(__e *ControlFlow) {
W5261 := __e.Get(1)
_ = W5261
tmp14138 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14138

__e.TailApply(PrimFunc(symshen_4search_1user_1datatypes), V5249, V5250, W5261, V5252, V5253, V5254, V5255)
return


}, 1)

tmp14139 := PrimTail(W5260)

__e.TailApply(tmp14137, tmp14139)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14142 := Call(__e, PrimFunc(symshen_4lazyderef), V5251, V5252)


__e.TailApply(tmp14136, tmp14142)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5256)
return
}


}, 1)

tmp14166 := Call(__e, PrimFunc(symshen_4unlocked_2), V5253)


var ifres14147 Obj

if True == tmp14166 {
tmp14148 := MakeNative(func(__e *ControlFlow) {
W5257 := __e.Get(1)
_ = W5257
tmp14163 := PrimIsPair(W5257)

if True == tmp14163 {
tmp14149 := MakeNative(func(__e *ControlFlow) {
W5258 := __e.Get(1)
_ = W5258
tmp14159 := PrimIsPair(W5258)

if True == tmp14159 {
tmp14150 := MakeNative(func(__e *ControlFlow) {
W5259 := __e.Get(1)
_ = W5259
tmp14151 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14151

tmp14152 := Call(__e, PrimFunc(symshen_4deref), W5259, V5252)


tmp14153 := Call(__e, PrimFunc(symshen_4deref), V5249, V5252)


tmp14154 := Call(__e, tmp14152, tmp14153)


tmp14155 := Call(__e, PrimFunc(symshen_4deref), V5250, V5252)


tmp14156 := Call(__e, tmp14154, tmp14155)


__e.TailApply(PrimFunc(symcall), tmp14156, V5252, V5253, V5254, V5255)
return


}, 1)

tmp14157 := PrimTail(W5258)

__e.TailApply(tmp14150, tmp14157)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14160 := PrimHead(W5257)

tmp14161 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14160, V5252)


__e.TailApply(tmp14149, tmp14161)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14164 := Call(__e, PrimFunc(symshen_4lazyderef), V5251, V5252)


tmp14165 := Call(__e, tmp14148, tmp14164)


ifres14147 = tmp14165


} else {
ifres14147 = False


}

__e.TailApply(tmp14135, ifres14147)
return


}, 7)

tmp14167 := Call(__e, ns2_1set, symshen_4search_1user_1datatypes, tmp14134)


_ = tmp14167

tmp14168 := MakeNative(func(__e *ControlFlow) {
V5262 := __e.Get(1)
_ = V5262
V5263 := __e.Get(2)
_ = V5263
V5264 := __e.Get(3)
_ = V5264
V5265 := __e.Get(4)
_ = V5265
V5266 := __e.Get(5)
_ = V5266
V5267 := __e.Get(6)
_ = V5267
V5268 := __e.Get(7)
_ = V5268
tmp14169 := MakeNative(func(__e *ControlFlow) {
W5269 := __e.Get(1)
_ = W5269
tmp14170 := MakeNative(func(__e *ControlFlow) {
W5270 := __e.Get(1)
_ = W5270
tmp14600 := PrimEqual(W5270, False)

if True == tmp14600 {
tmp14171 := MakeNative(func(__e *ControlFlow) {
W5273 := __e.Get(1)
_ = W5273
tmp14500 := PrimEqual(W5273, False)

if True == tmp14500 {
tmp14172 := MakeNative(func(__e *ControlFlow) {
W5293 := __e.Get(1)
_ = W5293
tmp14395 := PrimEqual(W5293, False)

if True == tmp14395 {
tmp14173 := MakeNative(func(__e *ControlFlow) {
W5315 := __e.Get(1)
_ = W5315
tmp14314 := PrimEqual(W5315, False)

if True == tmp14314 {
tmp14174 := MakeNative(func(__e *ControlFlow) {
W5331 := __e.Get(1)
_ = W5331
tmp14214 := PrimEqual(W5331, False)

if True == tmp14214 {
tmp14175 := MakeNative(func(__e *ControlFlow) {
W5351 := __e.Get(1)
_ = W5351
tmp14177 := PrimEqual(W5351, False)

if True == tmp14177 {
__e.TailApply(PrimFunc(symshen_4unlock), V5266, W5269)
return
} else {
__e.Return(W5351)
return
}


}, 1)

tmp14212 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14178 Obj

if True == tmp14212 {
tmp14179 := MakeNative(func(__e *ControlFlow) {
W5352 := __e.Get(1)
_ = W5352
tmp14209 := PrimIsPair(W5352)

if True == tmp14209 {
tmp14180 := MakeNative(func(__e *ControlFlow) {
W5353 := __e.Get(1)
_ = W5353
tmp14181 := MakeNative(func(__e *ControlFlow) {
W5354 := __e.Get(1)
_ = W5354
tmp14182 := MakeNative(func(__e *ControlFlow) {
W5355 := __e.Get(1)
_ = W5355
tmp14183 := MakeNative(func(__e *ControlFlow) {
W5356 := __e.Get(1)
_ = W5356
tmp14201 := PrimIsPair(W5355)

if True == tmp14201 {
tmp14184 := MakeNative(func(__e *ControlFlow) {
W5359 := __e.Get(1)
_ = W5359
tmp14185 := MakeNative(func(__e *ControlFlow) {
W5360 := __e.Get(1)
_ = W5360
tmp14186 := Call(__e, W5356, W5359)


__e.TailApply(tmp14186, W5360)
return


}, 1)

tmp14187 := PrimTail(W5355)

__e.TailApply(tmp14185, tmp14187)
return


}, 1)

tmp14188 := PrimHead(W5355)

__e.TailApply(tmp14184, tmp14188)
return


} else {
tmp14199 := Call(__e, PrimFunc(symshen_4pvar_2), W5355)


if True == tmp14199 {
tmp14189 := MakeNative(func(__e *ControlFlow) {
W5361 := __e.Get(1)
_ = W5361
tmp14190 := MakeNative(func(__e *ControlFlow) {
W5362 := __e.Get(1)
_ = W5362
tmp14191 := PrimCons(W5361, W5362)

tmp14192 := MakeNative(func(__e *ControlFlow) {
tmp14193 := Call(__e, W5356, W5361)


__e.TailApply(tmp14193, W5362)
return


}, 0)

tmp14194 := Call(__e, PrimFunc(symshen_4bind_b), W5355, tmp14191, V5265, tmp14192)


__e.TailApply(PrimFunc(symshen_4gc), V5265, tmp14194)
return


}, 1)

tmp14195 := Call(__e, PrimFunc(symshen_4newpv), V5265)


tmp14196 := Call(__e, tmp14190, tmp14195)


__e.TailApply(PrimFunc(symshen_4gc), V5265, tmp14196)
return


}, 1)

tmp14197 := Call(__e, PrimFunc(symshen_4newpv), V5265)


__e.TailApply(tmp14189, tmp14197)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp14202 := MakeNative(func(__e *ControlFlow) {
Z5357 := __e.Get(1)
_ = Z5357
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5358 := __e.Get(1)
_ = Z5358
tmp14203 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14203

tmp14204 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4l_1rules), W5354, Z5358, V5264, V5265, V5266, W5269, V5268)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5357, W5353, V5265, V5266, W5269, tmp14204)
return


}, 1))
return
}, 1)

__e.TailApply(tmp14183, tmp14202)
return


}, 1)

tmp14205 := Call(__e, PrimFunc(symshen_4lazyderef), V5263, V5265)


__e.TailApply(tmp14182, tmp14205)
return


}, 1)

tmp14206 := PrimTail(W5352)

__e.TailApply(tmp14181, tmp14206)
return


}, 1)

tmp14207 := PrimHead(W5352)

__e.TailApply(tmp14180, tmp14207)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14210 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14211 := Call(__e, tmp14179, tmp14210)


ifres14178 = tmp14211


} else {
ifres14178 = False


}

__e.TailApply(tmp14175, ifres14178)
return


} else {
__e.Return(W5331)
return
}


}, 1)

tmp14312 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14215 Obj

if True == tmp14312 {
tmp14216 := MakeNative(func(__e *ControlFlow) {
W5332 := __e.Get(1)
_ = W5332
tmp14309 := PrimIsPair(W5332)

if True == tmp14309 {
tmp14217 := MakeNative(func(__e *ControlFlow) {
W5333 := __e.Get(1)
_ = W5333
tmp14305 := PrimIsPair(W5333)

if True == tmp14305 {
tmp14218 := MakeNative(func(__e *ControlFlow) {
W5334 := __e.Get(1)
_ = W5334
tmp14301 := PrimIsPair(W5334)

if True == tmp14301 {
tmp14219 := MakeNative(func(__e *ControlFlow) {
W5335 := __e.Get(1)
_ = W5335
tmp14297 := PrimEqual(W5335, sym_8v)

if True == tmp14297 {
tmp14220 := MakeNative(func(__e *ControlFlow) {
W5336 := __e.Get(1)
_ = W5336
tmp14293 := PrimIsPair(W5336)

if True == tmp14293 {
tmp14221 := MakeNative(func(__e *ControlFlow) {
W5337 := __e.Get(1)
_ = W5337
tmp14222 := MakeNative(func(__e *ControlFlow) {
W5338 := __e.Get(1)
_ = W5338
tmp14288 := PrimIsPair(W5338)

if True == tmp14288 {
tmp14223 := MakeNative(func(__e *ControlFlow) {
W5339 := __e.Get(1)
_ = W5339
tmp14224 := MakeNative(func(__e *ControlFlow) {
W5340 := __e.Get(1)
_ = W5340
tmp14283 := PrimEqual(W5340, Nil)

if True == tmp14283 {
tmp14225 := MakeNative(func(__e *ControlFlow) {
W5341 := __e.Get(1)
_ = W5341
tmp14279 := PrimIsPair(W5341)

if True == tmp14279 {
tmp14226 := MakeNative(func(__e *ControlFlow) {
W5342 := __e.Get(1)
_ = W5342
tmp14227 := MakeNative(func(__e *ControlFlow) {
W5343 := __e.Get(1)
_ = W5343
tmp14274 := PrimIsPair(W5343)

if True == tmp14274 {
tmp14228 := MakeNative(func(__e *ControlFlow) {
W5344 := __e.Get(1)
_ = W5344
tmp14270 := PrimIsPair(W5344)

if True == tmp14270 {
tmp14229 := MakeNative(func(__e *ControlFlow) {
W5345 := __e.Get(1)
_ = W5345
tmp14266 := PrimEqual(W5345, symvector)

if True == tmp14266 {
tmp14230 := MakeNative(func(__e *ControlFlow) {
W5346 := __e.Get(1)
_ = W5346
tmp14262 := PrimIsPair(W5346)

if True == tmp14262 {
tmp14231 := MakeNative(func(__e *ControlFlow) {
W5347 := __e.Get(1)
_ = W5347
tmp14232 := MakeNative(func(__e *ControlFlow) {
W5348 := __e.Get(1)
_ = W5348
tmp14257 := PrimEqual(W5348, Nil)

if True == tmp14257 {
tmp14233 := MakeNative(func(__e *ControlFlow) {
W5349 := __e.Get(1)
_ = W5349
tmp14253 := PrimEqual(W5349, Nil)

if True == tmp14253 {
tmp14234 := MakeNative(func(__e *ControlFlow) {
W5350 := __e.Get(1)
_ = W5350
tmp14235 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14235

tmp14236 := Call(__e, PrimFunc(symshen_4deref), W5342, V5265)


tmp14237 := PrimIntern(MakeString(":"))

tmp14238 := PrimEqual(tmp14236, tmp14237)

tmp14239 := MakeNative(func(__e *ControlFlow) {
tmp14240 := MakeNative(func(__e *ControlFlow) {
tmp14241 := PrimCons(W5347, Nil)

tmp14242 := PrimCons(W5342, tmp14241)

tmp14243 := PrimCons(W5337, tmp14242)

tmp14244 := PrimCons(W5347, Nil)

tmp14245 := PrimCons(symvector, tmp14244)

tmp14246 := PrimCons(tmp14245, Nil)

tmp14247 := PrimCons(W5342, tmp14246)

tmp14248 := PrimCons(W5339, tmp14247)

tmp14249 := PrimCons(tmp14248, W5350)

tmp14250 := PrimCons(tmp14243, tmp14249)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14250, V5263, True, V5265, V5266, W5269, V5268)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5265, V5266, W5269, tmp14240)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14238, V5265, V5266, W5269, tmp14239)
return


}, 1)

tmp14251 := PrimTail(W5332)

__e.TailApply(tmp14234, tmp14251)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14254 := PrimTail(W5343)

tmp14255 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14254, V5265)


__e.TailApply(tmp14233, tmp14255)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14258 := PrimTail(W5346)

tmp14259 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14258, V5265)


__e.TailApply(tmp14232, tmp14259)
return


}, 1)

tmp14260 := PrimHead(W5346)

__e.TailApply(tmp14231, tmp14260)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14263 := PrimTail(W5344)

tmp14264 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14263, V5265)


__e.TailApply(tmp14230, tmp14264)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14267 := PrimHead(W5344)

tmp14268 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14267, V5265)


__e.TailApply(tmp14229, tmp14268)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14271 := PrimHead(W5343)

tmp14272 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14271, V5265)


__e.TailApply(tmp14228, tmp14272)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14275 := PrimTail(W5341)

tmp14276 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14275, V5265)


__e.TailApply(tmp14227, tmp14276)
return


}, 1)

tmp14277 := PrimHead(W5341)

__e.TailApply(tmp14226, tmp14277)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14280 := PrimTail(W5333)

tmp14281 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14280, V5265)


__e.TailApply(tmp14225, tmp14281)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14284 := PrimTail(W5338)

tmp14285 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14284, V5265)


__e.TailApply(tmp14224, tmp14285)
return


}, 1)

tmp14286 := PrimHead(W5338)

__e.TailApply(tmp14223, tmp14286)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14289 := PrimTail(W5336)

tmp14290 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14289, V5265)


__e.TailApply(tmp14222, tmp14290)
return


}, 1)

tmp14291 := PrimHead(W5336)

__e.TailApply(tmp14221, tmp14291)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14294 := PrimTail(W5334)

tmp14295 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14294, V5265)


__e.TailApply(tmp14220, tmp14295)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14298 := PrimHead(W5334)

tmp14299 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14298, V5265)


__e.TailApply(tmp14219, tmp14299)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14302 := PrimHead(W5333)

tmp14303 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14302, V5265)


__e.TailApply(tmp14218, tmp14303)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14306 := PrimHead(W5332)

tmp14307 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14306, V5265)


__e.TailApply(tmp14217, tmp14307)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14310 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14311 := Call(__e, tmp14216, tmp14310)


ifres14215 = tmp14311


} else {
ifres14215 = False


}

__e.TailApply(tmp14174, ifres14215)
return


} else {
__e.Return(W5315)
return
}


}, 1)

tmp14393 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14315 Obj

if True == tmp14393 {
tmp14316 := MakeNative(func(__e *ControlFlow) {
W5316 := __e.Get(1)
_ = W5316
tmp14390 := PrimIsPair(W5316)

if True == tmp14390 {
tmp14317 := MakeNative(func(__e *ControlFlow) {
W5317 := __e.Get(1)
_ = W5317
tmp14386 := PrimIsPair(W5317)

if True == tmp14386 {
tmp14318 := MakeNative(func(__e *ControlFlow) {
W5318 := __e.Get(1)
_ = W5318
tmp14382 := PrimIsPair(W5318)

if True == tmp14382 {
tmp14319 := MakeNative(func(__e *ControlFlow) {
W5319 := __e.Get(1)
_ = W5319
tmp14378 := PrimEqual(W5319, sym_8s)

if True == tmp14378 {
tmp14320 := MakeNative(func(__e *ControlFlow) {
W5320 := __e.Get(1)
_ = W5320
tmp14374 := PrimIsPair(W5320)

if True == tmp14374 {
tmp14321 := MakeNative(func(__e *ControlFlow) {
W5321 := __e.Get(1)
_ = W5321
tmp14322 := MakeNative(func(__e *ControlFlow) {
W5322 := __e.Get(1)
_ = W5322
tmp14369 := PrimIsPair(W5322)

if True == tmp14369 {
tmp14323 := MakeNative(func(__e *ControlFlow) {
W5323 := __e.Get(1)
_ = W5323
tmp14324 := MakeNative(func(__e *ControlFlow) {
W5324 := __e.Get(1)
_ = W5324
tmp14364 := PrimEqual(W5324, Nil)

if True == tmp14364 {
tmp14325 := MakeNative(func(__e *ControlFlow) {
W5325 := __e.Get(1)
_ = W5325
tmp14360 := PrimIsPair(W5325)

if True == tmp14360 {
tmp14326 := MakeNative(func(__e *ControlFlow) {
W5326 := __e.Get(1)
_ = W5326
tmp14327 := MakeNative(func(__e *ControlFlow) {
W5327 := __e.Get(1)
_ = W5327
tmp14355 := PrimIsPair(W5327)

if True == tmp14355 {
tmp14328 := MakeNative(func(__e *ControlFlow) {
W5328 := __e.Get(1)
_ = W5328
tmp14351 := PrimEqual(W5328, symstring)

if True == tmp14351 {
tmp14329 := MakeNative(func(__e *ControlFlow) {
W5329 := __e.Get(1)
_ = W5329
tmp14347 := PrimEqual(W5329, Nil)

if True == tmp14347 {
tmp14330 := MakeNative(func(__e *ControlFlow) {
W5330 := __e.Get(1)
_ = W5330
tmp14331 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14331

tmp14332 := Call(__e, PrimFunc(symshen_4deref), W5326, V5265)


tmp14333 := PrimIntern(MakeString(":"))

tmp14334 := PrimEqual(tmp14332, tmp14333)

tmp14335 := MakeNative(func(__e *ControlFlow) {
tmp14336 := MakeNative(func(__e *ControlFlow) {
tmp14337 := PrimCons(symstring, Nil)

tmp14338 := PrimCons(W5326, tmp14337)

tmp14339 := PrimCons(W5321, tmp14338)

tmp14340 := PrimCons(symstring, Nil)

tmp14341 := PrimCons(W5326, tmp14340)

tmp14342 := PrimCons(W5323, tmp14341)

tmp14343 := PrimCons(tmp14342, W5330)

tmp14344 := PrimCons(tmp14339, tmp14343)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14344, V5263, True, V5265, V5266, W5269, V5268)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5265, V5266, W5269, tmp14336)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14334, V5265, V5266, W5269, tmp14335)
return


}, 1)

tmp14345 := PrimTail(W5316)

__e.TailApply(tmp14330, tmp14345)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14348 := PrimTail(W5327)

tmp14349 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14348, V5265)


__e.TailApply(tmp14329, tmp14349)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14352 := PrimHead(W5327)

tmp14353 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14352, V5265)


__e.TailApply(tmp14328, tmp14353)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14356 := PrimTail(W5325)

tmp14357 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14356, V5265)


__e.TailApply(tmp14327, tmp14357)
return


}, 1)

tmp14358 := PrimHead(W5325)

__e.TailApply(tmp14326, tmp14358)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14361 := PrimTail(W5317)

tmp14362 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14361, V5265)


__e.TailApply(tmp14325, tmp14362)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14365 := PrimTail(W5322)

tmp14366 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14365, V5265)


__e.TailApply(tmp14324, tmp14366)
return


}, 1)

tmp14367 := PrimHead(W5322)

__e.TailApply(tmp14323, tmp14367)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14370 := PrimTail(W5320)

tmp14371 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14370, V5265)


__e.TailApply(tmp14322, tmp14371)
return


}, 1)

tmp14372 := PrimHead(W5320)

__e.TailApply(tmp14321, tmp14372)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14375 := PrimTail(W5318)

tmp14376 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14375, V5265)


__e.TailApply(tmp14320, tmp14376)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14379 := PrimHead(W5318)

tmp14380 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14379, V5265)


__e.TailApply(tmp14319, tmp14380)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14383 := PrimHead(W5317)

tmp14384 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14383, V5265)


__e.TailApply(tmp14318, tmp14384)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14387 := PrimHead(W5316)

tmp14388 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14387, V5265)


__e.TailApply(tmp14317, tmp14388)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14391 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14392 := Call(__e, tmp14316, tmp14391)


ifres14315 = tmp14392


} else {
ifres14315 = False


}

__e.TailApply(tmp14173, ifres14315)
return


} else {
__e.Return(W5293)
return
}


}, 1)

tmp14498 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14396 Obj

if True == tmp14498 {
tmp14397 := MakeNative(func(__e *ControlFlow) {
W5294 := __e.Get(1)
_ = W5294
tmp14495 := PrimIsPair(W5294)

if True == tmp14495 {
tmp14398 := MakeNative(func(__e *ControlFlow) {
W5295 := __e.Get(1)
_ = W5295
tmp14491 := PrimIsPair(W5295)

if True == tmp14491 {
tmp14399 := MakeNative(func(__e *ControlFlow) {
W5296 := __e.Get(1)
_ = W5296
tmp14487 := PrimIsPair(W5296)

if True == tmp14487 {
tmp14400 := MakeNative(func(__e *ControlFlow) {
W5297 := __e.Get(1)
_ = W5297
tmp14483 := PrimEqual(W5297, sym_8p)

if True == tmp14483 {
tmp14401 := MakeNative(func(__e *ControlFlow) {
W5298 := __e.Get(1)
_ = W5298
tmp14479 := PrimIsPair(W5298)

if True == tmp14479 {
tmp14402 := MakeNative(func(__e *ControlFlow) {
W5299 := __e.Get(1)
_ = W5299
tmp14403 := MakeNative(func(__e *ControlFlow) {
W5300 := __e.Get(1)
_ = W5300
tmp14474 := PrimIsPair(W5300)

if True == tmp14474 {
tmp14404 := MakeNative(func(__e *ControlFlow) {
W5301 := __e.Get(1)
_ = W5301
tmp14405 := MakeNative(func(__e *ControlFlow) {
W5302 := __e.Get(1)
_ = W5302
tmp14469 := PrimEqual(W5302, Nil)

if True == tmp14469 {
tmp14406 := MakeNative(func(__e *ControlFlow) {
W5303 := __e.Get(1)
_ = W5303
tmp14465 := PrimIsPair(W5303)

if True == tmp14465 {
tmp14407 := MakeNative(func(__e *ControlFlow) {
W5304 := __e.Get(1)
_ = W5304
tmp14408 := MakeNative(func(__e *ControlFlow) {
W5305 := __e.Get(1)
_ = W5305
tmp14460 := PrimIsPair(W5305)

if True == tmp14460 {
tmp14409 := MakeNative(func(__e *ControlFlow) {
W5306 := __e.Get(1)
_ = W5306
tmp14456 := PrimIsPair(W5306)

if True == tmp14456 {
tmp14410 := MakeNative(func(__e *ControlFlow) {
W5307 := __e.Get(1)
_ = W5307
tmp14411 := MakeNative(func(__e *ControlFlow) {
W5308 := __e.Get(1)
_ = W5308
tmp14451 := PrimIsPair(W5308)

if True == tmp14451 {
tmp14412 := MakeNative(func(__e *ControlFlow) {
W5309 := __e.Get(1)
_ = W5309
tmp14447 := PrimEqual(W5309, sym_d)

if True == tmp14447 {
tmp14413 := MakeNative(func(__e *ControlFlow) {
W5310 := __e.Get(1)
_ = W5310
tmp14443 := PrimIsPair(W5310)

if True == tmp14443 {
tmp14414 := MakeNative(func(__e *ControlFlow) {
W5311 := __e.Get(1)
_ = W5311
tmp14415 := MakeNative(func(__e *ControlFlow) {
W5312 := __e.Get(1)
_ = W5312
tmp14438 := PrimEqual(W5312, Nil)

if True == tmp14438 {
tmp14416 := MakeNative(func(__e *ControlFlow) {
W5313 := __e.Get(1)
_ = W5313
tmp14434 := PrimEqual(W5313, Nil)

if True == tmp14434 {
tmp14417 := MakeNative(func(__e *ControlFlow) {
W5314 := __e.Get(1)
_ = W5314
tmp14418 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14418

tmp14419 := Call(__e, PrimFunc(symshen_4deref), W5304, V5265)


tmp14420 := PrimIntern(MakeString(":"))

tmp14421 := PrimEqual(tmp14419, tmp14420)

tmp14422 := MakeNative(func(__e *ControlFlow) {
tmp14423 := MakeNative(func(__e *ControlFlow) {
tmp14424 := PrimCons(W5307, Nil)

tmp14425 := PrimCons(W5304, tmp14424)

tmp14426 := PrimCons(W5299, tmp14425)

tmp14427 := PrimCons(W5311, Nil)

tmp14428 := PrimCons(W5304, tmp14427)

tmp14429 := PrimCons(W5301, tmp14428)

tmp14430 := PrimCons(tmp14429, W5314)

tmp14431 := PrimCons(tmp14426, tmp14430)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14431, V5263, True, V5265, V5266, W5269, V5268)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5265, V5266, W5269, tmp14423)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14421, V5265, V5266, W5269, tmp14422)
return


}, 1)

tmp14432 := PrimTail(W5294)

__e.TailApply(tmp14417, tmp14432)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14435 := PrimTail(W5305)

tmp14436 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14435, V5265)


__e.TailApply(tmp14416, tmp14436)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14439 := PrimTail(W5310)

tmp14440 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14439, V5265)


__e.TailApply(tmp14415, tmp14440)
return


}, 1)

tmp14441 := PrimHead(W5310)

__e.TailApply(tmp14414, tmp14441)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14444 := PrimTail(W5308)

tmp14445 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14444, V5265)


__e.TailApply(tmp14413, tmp14445)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14448 := PrimHead(W5308)

tmp14449 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14448, V5265)


__e.TailApply(tmp14412, tmp14449)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14452 := PrimTail(W5306)

tmp14453 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14452, V5265)


__e.TailApply(tmp14411, tmp14453)
return


}, 1)

tmp14454 := PrimHead(W5306)

__e.TailApply(tmp14410, tmp14454)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14457 := PrimHead(W5305)

tmp14458 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14457, V5265)


__e.TailApply(tmp14409, tmp14458)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14461 := PrimTail(W5303)

tmp14462 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14461, V5265)


__e.TailApply(tmp14408, tmp14462)
return


}, 1)

tmp14463 := PrimHead(W5303)

__e.TailApply(tmp14407, tmp14463)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14466 := PrimTail(W5295)

tmp14467 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14466, V5265)


__e.TailApply(tmp14406, tmp14467)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14470 := PrimTail(W5300)

tmp14471 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14470, V5265)


__e.TailApply(tmp14405, tmp14471)
return


}, 1)

tmp14472 := PrimHead(W5300)

__e.TailApply(tmp14404, tmp14472)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14475 := PrimTail(W5298)

tmp14476 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14475, V5265)


__e.TailApply(tmp14403, tmp14476)
return


}, 1)

tmp14477 := PrimHead(W5298)

__e.TailApply(tmp14402, tmp14477)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14480 := PrimTail(W5296)

tmp14481 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14480, V5265)


__e.TailApply(tmp14401, tmp14481)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14484 := PrimHead(W5296)

tmp14485 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14484, V5265)


__e.TailApply(tmp14400, tmp14485)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14488 := PrimHead(W5295)

tmp14489 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14488, V5265)


__e.TailApply(tmp14399, tmp14489)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14492 := PrimHead(W5294)

tmp14493 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14492, V5265)


__e.TailApply(tmp14398, tmp14493)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14496 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14497 := Call(__e, tmp14397, tmp14496)


ifres14396 = tmp14497


} else {
ifres14396 = False


}

__e.TailApply(tmp14172, ifres14396)
return


} else {
__e.Return(W5273)
return
}


}, 1)

tmp14598 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14501 Obj

if True == tmp14598 {
tmp14502 := MakeNative(func(__e *ControlFlow) {
W5274 := __e.Get(1)
_ = W5274
tmp14595 := PrimIsPair(W5274)

if True == tmp14595 {
tmp14503 := MakeNative(func(__e *ControlFlow) {
W5275 := __e.Get(1)
_ = W5275
tmp14591 := PrimIsPair(W5275)

if True == tmp14591 {
tmp14504 := MakeNative(func(__e *ControlFlow) {
W5276 := __e.Get(1)
_ = W5276
tmp14587 := PrimIsPair(W5276)

if True == tmp14587 {
tmp14505 := MakeNative(func(__e *ControlFlow) {
W5277 := __e.Get(1)
_ = W5277
tmp14583 := PrimEqual(W5277, symcons)

if True == tmp14583 {
tmp14506 := MakeNative(func(__e *ControlFlow) {
W5278 := __e.Get(1)
_ = W5278
tmp14579 := PrimIsPair(W5278)

if True == tmp14579 {
tmp14507 := MakeNative(func(__e *ControlFlow) {
W5279 := __e.Get(1)
_ = W5279
tmp14508 := MakeNative(func(__e *ControlFlow) {
W5280 := __e.Get(1)
_ = W5280
tmp14574 := PrimIsPair(W5280)

if True == tmp14574 {
tmp14509 := MakeNative(func(__e *ControlFlow) {
W5281 := __e.Get(1)
_ = W5281
tmp14510 := MakeNative(func(__e *ControlFlow) {
W5282 := __e.Get(1)
_ = W5282
tmp14569 := PrimEqual(W5282, Nil)

if True == tmp14569 {
tmp14511 := MakeNative(func(__e *ControlFlow) {
W5283 := __e.Get(1)
_ = W5283
tmp14565 := PrimIsPair(W5283)

if True == tmp14565 {
tmp14512 := MakeNative(func(__e *ControlFlow) {
W5284 := __e.Get(1)
_ = W5284
tmp14513 := MakeNative(func(__e *ControlFlow) {
W5285 := __e.Get(1)
_ = W5285
tmp14560 := PrimIsPair(W5285)

if True == tmp14560 {
tmp14514 := MakeNative(func(__e *ControlFlow) {
W5286 := __e.Get(1)
_ = W5286
tmp14556 := PrimIsPair(W5286)

if True == tmp14556 {
tmp14515 := MakeNative(func(__e *ControlFlow) {
W5287 := __e.Get(1)
_ = W5287
tmp14552 := PrimEqual(W5287, symlist)

if True == tmp14552 {
tmp14516 := MakeNative(func(__e *ControlFlow) {
W5288 := __e.Get(1)
_ = W5288
tmp14548 := PrimIsPair(W5288)

if True == tmp14548 {
tmp14517 := MakeNative(func(__e *ControlFlow) {
W5289 := __e.Get(1)
_ = W5289
tmp14518 := MakeNative(func(__e *ControlFlow) {
W5290 := __e.Get(1)
_ = W5290
tmp14543 := PrimEqual(W5290, Nil)

if True == tmp14543 {
tmp14519 := MakeNative(func(__e *ControlFlow) {
W5291 := __e.Get(1)
_ = W5291
tmp14539 := PrimEqual(W5291, Nil)

if True == tmp14539 {
tmp14520 := MakeNative(func(__e *ControlFlow) {
W5292 := __e.Get(1)
_ = W5292
tmp14521 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14521

tmp14522 := Call(__e, PrimFunc(symshen_4deref), W5284, V5265)


tmp14523 := PrimIntern(MakeString(":"))

tmp14524 := PrimEqual(tmp14522, tmp14523)

tmp14525 := MakeNative(func(__e *ControlFlow) {
tmp14526 := MakeNative(func(__e *ControlFlow) {
tmp14527 := PrimCons(W5289, Nil)

tmp14528 := PrimCons(W5284, tmp14527)

tmp14529 := PrimCons(W5279, tmp14528)

tmp14530 := PrimCons(W5289, Nil)

tmp14531 := PrimCons(symlist, tmp14530)

tmp14532 := PrimCons(tmp14531, Nil)

tmp14533 := PrimCons(W5284, tmp14532)

tmp14534 := PrimCons(W5281, tmp14533)

tmp14535 := PrimCons(tmp14534, W5292)

tmp14536 := PrimCons(tmp14529, tmp14535)

__e.TailApply(PrimFunc(symshen_4l_1rules), tmp14536, V5263, True, V5265, V5266, W5269, V5268)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5265, V5266, W5269, tmp14526)
return


}, 0)

__e.TailApply(PrimFunc(symwhen), tmp14524, V5265, V5266, W5269, tmp14525)
return


}, 1)

tmp14537 := PrimTail(W5274)

__e.TailApply(tmp14520, tmp14537)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14540 := PrimTail(W5285)

tmp14541 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14540, V5265)


__e.TailApply(tmp14519, tmp14541)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14544 := PrimTail(W5288)

tmp14545 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14544, V5265)


__e.TailApply(tmp14518, tmp14545)
return


}, 1)

tmp14546 := PrimHead(W5288)

__e.TailApply(tmp14517, tmp14546)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14549 := PrimTail(W5286)

tmp14550 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14549, V5265)


__e.TailApply(tmp14516, tmp14550)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14553 := PrimHead(W5286)

tmp14554 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14553, V5265)


__e.TailApply(tmp14515, tmp14554)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14557 := PrimHead(W5285)

tmp14558 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14557, V5265)


__e.TailApply(tmp14514, tmp14558)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14561 := PrimTail(W5283)

tmp14562 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14561, V5265)


__e.TailApply(tmp14513, tmp14562)
return


}, 1)

tmp14563 := PrimHead(W5283)

__e.TailApply(tmp14512, tmp14563)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14566 := PrimTail(W5275)

tmp14567 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14566, V5265)


__e.TailApply(tmp14511, tmp14567)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14570 := PrimTail(W5280)

tmp14571 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14570, V5265)


__e.TailApply(tmp14510, tmp14571)
return


}, 1)

tmp14572 := PrimHead(W5280)

__e.TailApply(tmp14509, tmp14572)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14575 := PrimTail(W5278)

tmp14576 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14575, V5265)


__e.TailApply(tmp14508, tmp14576)
return


}, 1)

tmp14577 := PrimHead(W5278)

__e.TailApply(tmp14507, tmp14577)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14580 := PrimTail(W5276)

tmp14581 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14580, V5265)


__e.TailApply(tmp14506, tmp14581)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14584 := PrimHead(W5276)

tmp14585 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14584, V5265)


__e.TailApply(tmp14505, tmp14585)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14588 := PrimHead(W5275)

tmp14589 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14588, V5265)


__e.TailApply(tmp14504, tmp14589)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14592 := PrimHead(W5274)

tmp14593 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14592, V5265)


__e.TailApply(tmp14503, tmp14593)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14596 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14597 := Call(__e, tmp14502, tmp14596)


ifres14501 = tmp14597


} else {
ifres14501 = False


}

__e.TailApply(tmp14171, ifres14501)
return


} else {
__e.Return(W5270)
return
}


}, 1)

tmp14613 := Call(__e, PrimFunc(symshen_4unlocked_2), V5266)


var ifres14601 Obj

if True == tmp14613 {
tmp14602 := MakeNative(func(__e *ControlFlow) {
W5271 := __e.Get(1)
_ = W5271
tmp14610 := PrimEqual(W5271, Nil)

if True == tmp14610 {
tmp14603 := MakeNative(func(__e *ControlFlow) {
W5272 := __e.Get(1)
_ = W5272
tmp14607 := PrimEqual(W5272, True)

if True == tmp14607 {
tmp14604 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14604

tmp14605 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symbind), V5263, Nil, V5265, V5266, W5269, V5268)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5265, V5266, W5269, tmp14605)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14608 := Call(__e, PrimFunc(symshen_4lazyderef), V5264, V5265)


__e.TailApply(tmp14603, tmp14608)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14611 := Call(__e, PrimFunc(symshen_4lazyderef), V5262, V5265)


tmp14612 := Call(__e, tmp14602, tmp14611)


ifres14601 = tmp14612


} else {
ifres14601 = False


}

__e.TailApply(tmp14170, ifres14601)
return


}, 1)

tmp14614 := PrimNumberAdd(V5267, MakeNumber(1))

__e.TailApply(tmp14169, tmp14614)
return


}, 7)

tmp14615 := Call(__e, ns2_1set, symshen_4l_1rules, tmp14168)


_ = tmp14615

tmp14616 := MakeNative(func(__e *ControlFlow) {
V5363 := __e.Get(1)
_ = V5363
V5364 := __e.Get(2)
_ = V5364
V5365 := __e.Get(3)
_ = V5365
V5366 := __e.Get(4)
_ = V5366
V5367 := __e.Get(5)
_ = V5367
V5368 := __e.Get(6)
_ = V5368
tmp14617 := MakeNative(func(__e *ControlFlow) {
W5369 := __e.Get(1)
_ = W5369
tmp14618 := MakeNative(func(__e *ControlFlow) {
W5370 := __e.Get(1)
_ = W5370
tmp14620 := PrimEqual(W5370, False)

if True == tmp14620 {
__e.TailApply(PrimFunc(symshen_4unlock), V5366, W5369)
return
} else {
__e.Return(W5370)
return
}


}, 1)

tmp14668 := Call(__e, PrimFunc(symshen_4unlocked_2), V5366)


var ifres14621 Obj

if True == tmp14668 {
tmp14622 := MakeNative(func(__e *ControlFlow) {
W5371 := __e.Get(1)
_ = W5371
tmp14665 := PrimIsPair(W5371)

if True == tmp14665 {
tmp14623 := MakeNative(func(__e *ControlFlow) {
W5372 := __e.Get(1)
_ = W5372
tmp14661 := PrimEqual(W5372, symdefine)

if True == tmp14661 {
tmp14624 := MakeNative(func(__e *ControlFlow) {
W5373 := __e.Get(1)
_ = W5373
tmp14657 := PrimIsPair(W5373)

if True == tmp14657 {
tmp14625 := MakeNative(func(__e *ControlFlow) {
W5374 := __e.Get(1)
_ = W5374
tmp14626 := MakeNative(func(__e *ControlFlow) {
W5375 := __e.Get(1)
_ = W5375
tmp14627 := MakeNative(func(__e *ControlFlow) {
W5376 := __e.Get(1)
_ = W5376
tmp14628 := MakeNative(func(__e *ControlFlow) {
W5377 := __e.Get(1)
_ = W5377
tmp14629 := MakeNative(func(__e *ControlFlow) {
W5378 := __e.Get(1)
_ = W5378
tmp14630 := MakeNative(func(__e *ControlFlow) {
W5379 := __e.Get(1)
_ = W5379
tmp14631 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp14631

tmp14632 := MakeNative(func(__e *ControlFlow) {
tmp14633 := PrimCons(W5374, W5375)

tmp14634 := Call(__e, PrimFunc(symshen_4sigxrules), tmp14633)


tmp14635 := MakeNative(func(__e *ControlFlow) {
tmp14636 := Call(__e, PrimFunc(symshen_4lazyderef), W5376, V5365)


tmp14637 := Call(__e, PrimFunc(symfst), tmp14636)


tmp14638 := MakeNative(func(__e *ControlFlow) {
tmp14639 := Call(__e, PrimFunc(symshen_4lazyderef), W5376, V5365)


tmp14640 := Call(__e, PrimFunc(symsnd), tmp14639)


tmp14641 := MakeNative(func(__e *ControlFlow) {
tmp14642 := Call(__e, PrimFunc(symshen_4deref), W5379, V5365)


tmp14643 := Call(__e, PrimFunc(symshen_4freshen_1sig), tmp14642)


tmp14644 := MakeNative(func(__e *ControlFlow) {
tmp14645 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symis), W5379, V5364, V5365, V5366, W5369, V5368)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rules), W5374, W5377, W5378, MakeNumber(1), V5365, V5366, W5369, tmp14645)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5378, tmp14643, V5365, V5366, W5369, tmp14644)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5377, tmp14640, V5365, V5366, W5369, tmp14641)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5379, tmp14637, V5365, V5366, W5369, tmp14638)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5376, tmp14634, V5365, V5366, W5369, tmp14635)
return


}, 0)

tmp14646 := Call(__e, PrimFunc(symshen_4cut), V5365, V5366, W5369, tmp14632)


__e.TailApply(PrimFunc(symshen_4gc), V5365, tmp14646)
return


}, 1)

tmp14647 := Call(__e, PrimFunc(symshen_4newpv), V5365)


tmp14648 := Call(__e, tmp14630, tmp14647)


__e.TailApply(PrimFunc(symshen_4gc), V5365, tmp14648)
return


}, 1)

tmp14649 := Call(__e, PrimFunc(symshen_4newpv), V5365)


tmp14650 := Call(__e, tmp14629, tmp14649)


__e.TailApply(PrimFunc(symshen_4gc), V5365, tmp14650)
return


}, 1)

tmp14651 := Call(__e, PrimFunc(symshen_4newpv), V5365)


tmp14652 := Call(__e, tmp14628, tmp14651)


__e.TailApply(PrimFunc(symshen_4gc), V5365, tmp14652)
return


}, 1)

tmp14653 := Call(__e, PrimFunc(symshen_4newpv), V5365)


__e.TailApply(tmp14627, tmp14653)
return


}, 1)

tmp14654 := PrimTail(W5373)

__e.TailApply(tmp14626, tmp14654)
return


}, 1)

tmp14655 := PrimHead(W5373)

__e.TailApply(tmp14625, tmp14655)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14658 := PrimTail(W5371)

tmp14659 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14658, V5365)


__e.TailApply(tmp14624, tmp14659)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14662 := PrimHead(W5371)

tmp14663 := Call(__e, PrimFunc(symshen_4lazyderef), tmp14662, V5365)


__e.TailApply(tmp14623, tmp14663)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp14666 := Call(__e, PrimFunc(symshen_4lazyderef), V5363, V5365)


tmp14667 := Call(__e, tmp14622, tmp14666)


ifres14621 = tmp14667


} else {
ifres14621 = False


}

__e.TailApply(tmp14618, ifres14621)
return


}, 1)

tmp14669 := PrimNumberAdd(V5367, MakeNumber(1))

__e.TailApply(tmp14617, tmp14669)
return


}, 6)

tmp14670 := Call(__e, ns2_1set, symshen_4t_d, tmp14616)


_ = tmp14670

tmp14671 := MakeNative(func(__e *ControlFlow) {
V5380 := __e.Get(1)
_ = V5380
tmp14672 := MakeNative(func(__e *ControlFlow) {
Z5381 := __e.Get(1)
_ = Z5381
__e.TailApply(PrimFunc(symshen_4_5sig_drules_6), Z5381)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp14672, V5380)
return


}, 1)

tmp14673 := Call(__e, ns2_1set, symshen_4sigxrules, tmp14671)


_ = tmp14673

tmp14674 := MakeNative(func(__e *ControlFlow) {
V5382 := __e.Get(1)
_ = V5382
tmp14675 := MakeNative(func(__e *ControlFlow) {
W5383 := __e.Get(1)
_ = W5383
tmp14677 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5383)


if True == tmp14677 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5383)
return
}


}, 1)

tmp14710 := PrimIsPair(V5382)

var ifres14678 Obj

if True == tmp14710 {
tmp14679 := MakeNative(func(__e *ControlFlow) {
W5384 := __e.Get(1)
_ = W5384
tmp14706 := Call(__e, PrimFunc(symshen_4hds_a_2), W5384, sym_i)


if True == tmp14706 {
tmp14680 := MakeNative(func(__e *ControlFlow) {
W5385 := __e.Get(1)
_ = W5385
tmp14681 := MakeNative(func(__e *ControlFlow) {
W5386 := __e.Get(1)
_ = W5386
tmp14702 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5386)


if True == tmp14702 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14682 := MakeNative(func(__e *ControlFlow) {
W5387 := __e.Get(1)
_ = W5387
tmp14683 := MakeNative(func(__e *ControlFlow) {
W5388 := __e.Get(1)
_ = W5388
tmp14698 := Call(__e, PrimFunc(symshen_4hds_a_2), W5388, sym_j)


if True == tmp14698 {
tmp14684 := MakeNative(func(__e *ControlFlow) {
W5389 := __e.Get(1)
_ = W5389
tmp14685 := MakeNative(func(__e *ControlFlow) {
W5390 := __e.Get(1)
_ = W5390
tmp14694 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5390)


if True == tmp14694 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14686 := MakeNative(func(__e *ControlFlow) {
W5391 := __e.Get(1)
_ = W5391
tmp14687 := MakeNative(func(__e *ControlFlow) {
W5392 := __e.Get(1)
_ = W5392
tmp14688 := MakeNative(func(__e *ControlFlow) {
W5393 := __e.Get(1)
_ = W5393
__e.TailApply(PrimFunc(sym_8p), W5393, W5391)
return
}, 1)

tmp14689 := Call(__e, PrimFunc(symshen_4rectify_1type), W5387)


tmp14690 := Call(__e, tmp14688, tmp14689)


__e.TailApply(PrimFunc(symshen_4comb), W5392, tmp14690)
return


}, 1)

tmp14691 := Call(__e, PrimFunc(symshen_4in_1_6), W5390)


__e.TailApply(tmp14687, tmp14691)
return


}, 1)

tmp14692 := Call(__e, PrimFunc(symshen_4_5_1out), W5390)


__e.TailApply(tmp14686, tmp14692)
return


}


}, 1)

tmp14695 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5389)


__e.TailApply(tmp14685, tmp14695)
return


}, 1)

tmp14696 := Call(__e, PrimFunc(symtail), W5388)


__e.TailApply(tmp14684, tmp14696)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14699 := Call(__e, PrimFunc(symshen_4in_1_6), W5386)


__e.TailApply(tmp14683, tmp14699)
return


}, 1)

tmp14700 := Call(__e, PrimFunc(symshen_4_5_1out), W5386)


__e.TailApply(tmp14682, tmp14700)
return


}


}, 1)

tmp14703 := Call(__e, PrimFunc(symshen_4_5signature_6), W5385)


__e.TailApply(tmp14681, tmp14703)
return


}, 1)

tmp14704 := Call(__e, PrimFunc(symtail), W5384)


__e.TailApply(tmp14680, tmp14704)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14707 := Call(__e, PrimFunc(symtail), V5382)


tmp14708 := Call(__e, tmp14679, tmp14707)


ifres14678 = tmp14708


} else {
tmp14709 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres14678 = tmp14709


}

__e.TailApply(tmp14675, ifres14678)
return


}, 1)

tmp14711 := Call(__e, ns2_1set, symshen_4_5sig_drules_6, tmp14674)


_ = tmp14711

tmp14712 := MakeNative(func(__e *ControlFlow) {
V5394 := __e.Get(1)
_ = V5394
tmp14713 := MakeNative(func(__e *ControlFlow) {
W5395 := __e.Get(1)
_ = W5395
tmp14714 := MakeNative(func(__e *ControlFlow) {
W5396 := __e.Get(1)
_ = W5396
__e.TailApply(PrimFunc(symshen_4freshen_1type), W5396, V5394)
return
}, 1)

tmp14715 := MakeNative(func(__e *ControlFlow) {
Z5397 := __e.Get(1)
_ = Z5397
tmp14716 := Call(__e, PrimFunc(symconcat), sym_e, Z5397)


tmp14717 := Call(__e, PrimFunc(symshen_4freshterm), tmp14716)


__e.Return(PrimCons(Z5397, tmp14717))
return


}, 1)

tmp14718 := Call(__e, PrimFunc(symmap), tmp14715, W5395)


__e.TailApply(tmp14714, tmp14718)
return


}, 1)

tmp14719 := Call(__e, PrimFunc(symshen_4extract_1vars), V5394)


__e.TailApply(tmp14713, tmp14719)
return


}, 1)

tmp14720 := Call(__e, ns2_1set, symshen_4freshen_1sig, tmp14712)


_ = tmp14720

tmp14721 := MakeNative(func(__e *ControlFlow) {
V5398 := __e.Get(1)
_ = V5398
V5399 := __e.Get(2)
_ = V5399
tmp14735 := PrimEqual(Nil, V5398)

if True == tmp14735 {
__e.Return(V5399)
return
} else {
tmp14733 := PrimIsPair(V5398)

var ifres14729 Obj

if True == tmp14733 {
tmp14731 := PrimHead(V5398)

tmp14732 := PrimIsPair(tmp14731)

var ifres14730 Obj

if True == tmp14732 {
ifres14730 = True


} else {
ifres14730 = False


}

ifres14729 = ifres14730


} else {
ifres14729 = False


}

if True == ifres14729 {
tmp14722 := PrimTail(V5398)

tmp14723 := PrimHead(V5398)

tmp14724 := PrimTail(tmp14723)

tmp14725 := PrimHead(V5398)

tmp14726 := PrimHead(tmp14725)

tmp14727 := Call(__e, PrimFunc(symsubst), tmp14724, tmp14726, V5399)


__e.TailApply(PrimFunc(symshen_4freshen_1type), tmp14722, tmp14727)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1type)
return
}


}


}, 2)

tmp14736 := Call(__e, ns2_1set, symshen_4freshen_1type, tmp14721)


_ = tmp14736

tmp14737 := MakeNative(func(__e *ControlFlow) {
V5400 := __e.Get(1)
_ = V5400
tmp14738 := MakeNative(func(__e *ControlFlow) {
W5401 := __e.Get(1)
_ = W5401
tmp14753 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5401)


if True == tmp14753 {
tmp14739 := MakeNative(func(__e *ControlFlow) {
W5408 := __e.Get(1)
_ = W5408
tmp14741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5408)


if True == tmp14741 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5408)
return
}


}, 1)

tmp14742 := MakeNative(func(__e *ControlFlow) {
W5409 := __e.Get(1)
_ = W5409
tmp14749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5409)


if True == tmp14749 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14743 := MakeNative(func(__e *ControlFlow) {
W5410 := __e.Get(1)
_ = W5410
tmp14744 := MakeNative(func(__e *ControlFlow) {
W5411 := __e.Get(1)
_ = W5411
tmp14745 := PrimCons(W5410, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W5411, tmp14745)
return


}, 1)

tmp14746 := Call(__e, PrimFunc(symshen_4in_1_6), W5409)


__e.TailApply(tmp14744, tmp14746)
return


}, 1)

tmp14747 := Call(__e, PrimFunc(symshen_4_5_1out), W5409)


__e.TailApply(tmp14743, tmp14747)
return


}


}, 1)

tmp14750 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5400)


tmp14751 := Call(__e, tmp14742, tmp14750)


__e.TailApply(tmp14739, tmp14751)
return


} else {
__e.Return(W5401)
return
}


}, 1)

tmp14754 := MakeNative(func(__e *ControlFlow) {
W5402 := __e.Get(1)
_ = W5402
tmp14769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5402)


if True == tmp14769 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14755 := MakeNative(func(__e *ControlFlow) {
W5403 := __e.Get(1)
_ = W5403
tmp14756 := MakeNative(func(__e *ControlFlow) {
W5404 := __e.Get(1)
_ = W5404
tmp14757 := MakeNative(func(__e *ControlFlow) {
W5405 := __e.Get(1)
_ = W5405
tmp14764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5405)


if True == tmp14764 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14758 := MakeNative(func(__e *ControlFlow) {
W5406 := __e.Get(1)
_ = W5406
tmp14759 := MakeNative(func(__e *ControlFlow) {
W5407 := __e.Get(1)
_ = W5407
tmp14760 := PrimCons(W5403, W5406)

__e.TailApply(PrimFunc(symshen_4comb), W5407, tmp14760)
return


}, 1)

tmp14761 := Call(__e, PrimFunc(symshen_4in_1_6), W5405)


__e.TailApply(tmp14759, tmp14761)
return


}, 1)

tmp14762 := Call(__e, PrimFunc(symshen_4_5_1out), W5405)


__e.TailApply(tmp14758, tmp14762)
return


}


}, 1)

tmp14765 := Call(__e, PrimFunc(symshen_4_5rules_d_6), W5404)


__e.TailApply(tmp14757, tmp14765)
return


}, 1)

tmp14766 := Call(__e, PrimFunc(symshen_4in_1_6), W5402)


__e.TailApply(tmp14756, tmp14766)
return


}, 1)

tmp14767 := Call(__e, PrimFunc(symshen_4_5_1out), W5402)


__e.TailApply(tmp14755, tmp14767)
return


}


}, 1)

tmp14770 := Call(__e, PrimFunc(symshen_4_5rule_d_6), V5400)


tmp14771 := Call(__e, tmp14754, tmp14770)


__e.TailApply(tmp14738, tmp14771)
return


}, 1)

tmp14772 := Call(__e, ns2_1set, symshen_4_5rules_d_6, tmp14737)


_ = tmp14772

tmp14773 := MakeNative(func(__e *ControlFlow) {
V5412 := __e.Get(1)
_ = V5412
tmp14774 := MakeNative(func(__e *ControlFlow) {
W5413 := __e.Get(1)
_ = W5413
tmp14860 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5413)


if True == tmp14860 {
tmp14775 := MakeNative(func(__e *ControlFlow) {
W5423 := __e.Get(1)
_ = W5423
tmp14824 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5423)


if True == tmp14824 {
tmp14776 := MakeNative(func(__e *ControlFlow) {
W5433 := __e.Get(1)
_ = W5433
tmp14801 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5433)


if True == tmp14801 {
tmp14777 := MakeNative(func(__e *ControlFlow) {
W5440 := __e.Get(1)
_ = W5440
tmp14779 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5440)


if True == tmp14779 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W5440)
return
}


}, 1)

tmp14780 := MakeNative(func(__e *ControlFlow) {
W5441 := __e.Get(1)
_ = W5441
tmp14797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5441)


if True == tmp14797 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14781 := MakeNative(func(__e *ControlFlow) {
W5442 := __e.Get(1)
_ = W5442
tmp14782 := MakeNative(func(__e *ControlFlow) {
W5443 := __e.Get(1)
_ = W5443
tmp14793 := Call(__e, PrimFunc(symshen_4hds_a_2), W5443, sym_1_6)


if True == tmp14793 {
tmp14783 := MakeNative(func(__e *ControlFlow) {
W5444 := __e.Get(1)
_ = W5444
tmp14790 := PrimIsPair(W5444)

if True == tmp14790 {
tmp14784 := MakeNative(func(__e *ControlFlow) {
W5445 := __e.Get(1)
_ = W5445
tmp14785 := MakeNative(func(__e *ControlFlow) {
W5446 := __e.Get(1)
_ = W5446
tmp14786 := Call(__e, PrimFunc(sym_8p), W5442, W5445)


__e.TailApply(PrimFunc(symshen_4comb), W5446, tmp14786)
return


}, 1)

tmp14787 := Call(__e, PrimFunc(symtail), W5444)


__e.TailApply(tmp14785, tmp14787)
return


}, 1)

tmp14788 := Call(__e, PrimFunc(symhead), W5444)


__e.TailApply(tmp14784, tmp14788)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14791 := Call(__e, PrimFunc(symtail), W5443)


__e.TailApply(tmp14783, tmp14791)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14794 := Call(__e, PrimFunc(symshen_4in_1_6), W5441)


__e.TailApply(tmp14782, tmp14794)
return


}, 1)

tmp14795 := Call(__e, PrimFunc(symshen_4_5_1out), W5441)


__e.TailApply(tmp14781, tmp14795)
return


}


}, 1)

tmp14798 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5412)


tmp14799 := Call(__e, tmp14780, tmp14798)


__e.TailApply(tmp14777, tmp14799)
return


} else {
__e.Return(W5433)
return
}


}, 1)

tmp14802 := MakeNative(func(__e *ControlFlow) {
W5434 := __e.Get(1)
_ = W5434
tmp14820 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5434)


if True == tmp14820 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14803 := MakeNative(func(__e *ControlFlow) {
W5435 := __e.Get(1)
_ = W5435
tmp14804 := MakeNative(func(__e *ControlFlow) {
W5436 := __e.Get(1)
_ = W5436
tmp14816 := Call(__e, PrimFunc(symshen_4hds_a_2), W5436, sym_5_1)


if True == tmp14816 {
tmp14805 := MakeNative(func(__e *ControlFlow) {
W5437 := __e.Get(1)
_ = W5437
tmp14813 := PrimIsPair(W5437)

if True == tmp14813 {
tmp14806 := MakeNative(func(__e *ControlFlow) {
W5438 := __e.Get(1)
_ = W5438
tmp14807 := MakeNative(func(__e *ControlFlow) {
W5439 := __e.Get(1)
_ = W5439
tmp14808 := Call(__e, PrimFunc(symshen_4correct), W5438)


tmp14809 := Call(__e, PrimFunc(sym_8p), W5435, tmp14808)


__e.TailApply(PrimFunc(symshen_4comb), W5439, tmp14809)
return


}, 1)

tmp14810 := Call(__e, PrimFunc(symtail), W5437)


__e.TailApply(tmp14807, tmp14810)
return


}, 1)

tmp14811 := Call(__e, PrimFunc(symhead), W5437)


__e.TailApply(tmp14806, tmp14811)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14814 := Call(__e, PrimFunc(symtail), W5436)


__e.TailApply(tmp14805, tmp14814)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14817 := Call(__e, PrimFunc(symshen_4in_1_6), W5434)


__e.TailApply(tmp14804, tmp14817)
return


}, 1)

tmp14818 := Call(__e, PrimFunc(symshen_4_5_1out), W5434)


__e.TailApply(tmp14803, tmp14818)
return


}


}, 1)

tmp14821 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5412)


tmp14822 := Call(__e, tmp14802, tmp14821)


__e.TailApply(tmp14776, tmp14822)
return


} else {
__e.Return(W5423)
return
}


}, 1)

tmp14825 := MakeNative(func(__e *ControlFlow) {
W5424 := __e.Get(1)
_ = W5424
tmp14856 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5424)


if True == tmp14856 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14826 := MakeNative(func(__e *ControlFlow) {
W5425 := __e.Get(1)
_ = W5425
tmp14827 := MakeNative(func(__e *ControlFlow) {
W5426 := __e.Get(1)
_ = W5426
tmp14852 := Call(__e, PrimFunc(symshen_4hds_a_2), W5426, sym_5_1)


if True == tmp14852 {
tmp14828 := MakeNative(func(__e *ControlFlow) {
W5427 := __e.Get(1)
_ = W5427
tmp14849 := PrimIsPair(W5427)

if True == tmp14849 {
tmp14829 := MakeNative(func(__e *ControlFlow) {
W5428 := __e.Get(1)
_ = W5428
tmp14830 := MakeNative(func(__e *ControlFlow) {
W5429 := __e.Get(1)
_ = W5429
tmp14845 := Call(__e, PrimFunc(symshen_4hds_a_2), W5429, symwhere)


if True == tmp14845 {
tmp14831 := MakeNative(func(__e *ControlFlow) {
W5430 := __e.Get(1)
_ = W5430
tmp14842 := PrimIsPair(W5430)

if True == tmp14842 {
tmp14832 := MakeNative(func(__e *ControlFlow) {
W5431 := __e.Get(1)
_ = W5431
tmp14833 := MakeNative(func(__e *ControlFlow) {
W5432 := __e.Get(1)
_ = W5432
tmp14834 := PrimCons(W5428, Nil)

tmp14835 := PrimCons(W5431, tmp14834)

tmp14836 := PrimCons(symwhere, tmp14835)

tmp14837 := Call(__e, PrimFunc(symshen_4correct), tmp14836)


tmp14838 := Call(__e, PrimFunc(sym_8p), W5425, tmp14837)


__e.TailApply(PrimFunc(symshen_4comb), W5432, tmp14838)
return


}, 1)

tmp14839 := Call(__e, PrimFunc(symtail), W5430)


__e.TailApply(tmp14833, tmp14839)
return


}, 1)

tmp14840 := Call(__e, PrimFunc(symhead), W5430)


__e.TailApply(tmp14832, tmp14840)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14843 := Call(__e, PrimFunc(symtail), W5429)


__e.TailApply(tmp14831, tmp14843)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14846 := Call(__e, PrimFunc(symtail), W5427)


__e.TailApply(tmp14830, tmp14846)
return


}, 1)

tmp14847 := Call(__e, PrimFunc(symhead), W5427)


__e.TailApply(tmp14829, tmp14847)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14850 := Call(__e, PrimFunc(symtail), W5426)


__e.TailApply(tmp14828, tmp14850)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14853 := Call(__e, PrimFunc(symshen_4in_1_6), W5424)


__e.TailApply(tmp14827, tmp14853)
return


}, 1)

tmp14854 := Call(__e, PrimFunc(symshen_4_5_1out), W5424)


__e.TailApply(tmp14826, tmp14854)
return


}


}, 1)

tmp14857 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5412)


tmp14858 := Call(__e, tmp14825, tmp14857)


__e.TailApply(tmp14775, tmp14858)
return


} else {
__e.Return(W5413)
return
}


}, 1)

tmp14861 := MakeNative(func(__e *ControlFlow) {
W5414 := __e.Get(1)
_ = W5414
tmp14891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W5414)


if True == tmp14891 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp14862 := MakeNative(func(__e *ControlFlow) {
W5415 := __e.Get(1)
_ = W5415
tmp14863 := MakeNative(func(__e *ControlFlow) {
W5416 := __e.Get(1)
_ = W5416
tmp14887 := Call(__e, PrimFunc(symshen_4hds_a_2), W5416, sym_1_6)


if True == tmp14887 {
tmp14864 := MakeNative(func(__e *ControlFlow) {
W5417 := __e.Get(1)
_ = W5417
tmp14884 := PrimIsPair(W5417)

if True == tmp14884 {
tmp14865 := MakeNative(func(__e *ControlFlow) {
W5418 := __e.Get(1)
_ = W5418
tmp14866 := MakeNative(func(__e *ControlFlow) {
W5419 := __e.Get(1)
_ = W5419
tmp14880 := Call(__e, PrimFunc(symshen_4hds_a_2), W5419, symwhere)


if True == tmp14880 {
tmp14867 := MakeNative(func(__e *ControlFlow) {
W5420 := __e.Get(1)
_ = W5420
tmp14877 := PrimIsPair(W5420)

if True == tmp14877 {
tmp14868 := MakeNative(func(__e *ControlFlow) {
W5421 := __e.Get(1)
_ = W5421
tmp14869 := MakeNative(func(__e *ControlFlow) {
W5422 := __e.Get(1)
_ = W5422
tmp14870 := PrimCons(W5418, Nil)

tmp14871 := PrimCons(W5421, tmp14870)

tmp14872 := PrimCons(symwhere, tmp14871)

tmp14873 := Call(__e, PrimFunc(sym_8p), W5415, tmp14872)


__e.TailApply(PrimFunc(symshen_4comb), W5422, tmp14873)
return


}, 1)

tmp14874 := Call(__e, PrimFunc(symtail), W5420)


__e.TailApply(tmp14869, tmp14874)
return


}, 1)

tmp14875 := Call(__e, PrimFunc(symhead), W5420)


__e.TailApply(tmp14868, tmp14875)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14878 := Call(__e, PrimFunc(symtail), W5419)


__e.TailApply(tmp14867, tmp14878)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14881 := Call(__e, PrimFunc(symtail), W5417)


__e.TailApply(tmp14866, tmp14881)
return


}, 1)

tmp14882 := Call(__e, PrimFunc(symhead), W5417)


__e.TailApply(tmp14865, tmp14882)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14885 := Call(__e, PrimFunc(symtail), W5416)


__e.TailApply(tmp14864, tmp14885)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp14888 := Call(__e, PrimFunc(symshen_4in_1_6), W5414)


__e.TailApply(tmp14863, tmp14888)
return


}, 1)

tmp14889 := Call(__e, PrimFunc(symshen_4_5_1out), W5414)


__e.TailApply(tmp14862, tmp14889)
return


}


}, 1)

tmp14892 := Call(__e, PrimFunc(symshen_4_5patterns_6), V5412)


tmp14893 := Call(__e, tmp14861, tmp14892)


__e.TailApply(tmp14774, tmp14893)
return


}, 1)

tmp14894 := Call(__e, ns2_1set, symshen_4_5rule_d_6, tmp14773)


_ = tmp14894

tmp14895 := MakeNative(func(__e *ControlFlow) {
V5447 := __e.Get(1)
_ = V5447
tmp15043 := PrimIsPair(V5447)

var ifres14987 Obj

if True == tmp15043 {
tmp15041 := PrimHead(V5447)

tmp15042 := PrimEqual(symwhere, tmp15041)

var ifres14989 Obj

if True == tmp15042 {
tmp15039 := PrimTail(V5447)

tmp15040 := PrimIsPair(tmp15039)

var ifres14991 Obj

if True == tmp15040 {
tmp15036 := PrimTail(V5447)

tmp15037 := PrimTail(tmp15036)

tmp15038 := PrimIsPair(tmp15037)

var ifres14993 Obj

if True == tmp15038 {
tmp15032 := PrimTail(V5447)

tmp15033 := PrimTail(tmp15032)

tmp15034 := PrimHead(tmp15033)

tmp15035 := PrimIsPair(tmp15034)

var ifres14995 Obj

if True == tmp15035 {
tmp15027 := PrimTail(V5447)

tmp15028 := PrimTail(tmp15027)

tmp15029 := PrimHead(tmp15028)

tmp15030 := PrimHead(tmp15029)

tmp15031 := PrimEqual(symfail_1if, tmp15030)

var ifres14997 Obj

if True == tmp15031 {
tmp15022 := PrimTail(V5447)

tmp15023 := PrimTail(tmp15022)

tmp15024 := PrimHead(tmp15023)

tmp15025 := PrimTail(tmp15024)

tmp15026 := PrimIsPair(tmp15025)

var ifres14999 Obj

if True == tmp15026 {
tmp15016 := PrimTail(V5447)

tmp15017 := PrimTail(tmp15016)

tmp15018 := PrimHead(tmp15017)

tmp15019 := PrimTail(tmp15018)

tmp15020 := PrimTail(tmp15019)

tmp15021 := PrimIsPair(tmp15020)

var ifres15001 Obj

if True == tmp15021 {
tmp15009 := PrimTail(V5447)

tmp15010 := PrimTail(tmp15009)

tmp15011 := PrimHead(tmp15010)

tmp15012 := PrimTail(tmp15011)

tmp15013 := PrimTail(tmp15012)

tmp15014 := PrimTail(tmp15013)

tmp15015 := PrimEqual(Nil, tmp15014)

var ifres15003 Obj

if True == tmp15015 {
tmp15005 := PrimTail(V5447)

tmp15006 := PrimTail(tmp15005)

tmp15007 := PrimTail(tmp15006)

tmp15008 := PrimEqual(Nil, tmp15007)

var ifres15004 Obj

if True == tmp15008 {
ifres15004 = True


} else {
ifres15004 = False


}

ifres15003 = ifres15004


} else {
ifres15003 = False


}

var ifres15002 Obj

if True == ifres15003 {
ifres15002 = True


} else {
ifres15002 = False


}

ifres15001 = ifres15002


} else {
ifres15001 = False


}

var ifres15000 Obj

if True == ifres15001 {
ifres15000 = True


} else {
ifres15000 = False


}

ifres14999 = ifres15000


} else {
ifres14999 = False


}

var ifres14998 Obj

if True == ifres14999 {
ifres14998 = True


} else {
ifres14998 = False


}

ifres14997 = ifres14998


} else {
ifres14997 = False


}

var ifres14996 Obj

if True == ifres14997 {
ifres14996 = True


} else {
ifres14996 = False


}

ifres14995 = ifres14996


} else {
ifres14995 = False


}

var ifres14994 Obj

if True == ifres14995 {
ifres14994 = True


} else {
ifres14994 = False


}

ifres14993 = ifres14994


} else {
ifres14993 = False


}

var ifres14992 Obj

if True == ifres14993 {
ifres14992 = True


} else {
ifres14992 = False


}

ifres14991 = ifres14992


} else {
ifres14991 = False


}

var ifres14990 Obj

if True == ifres14991 {
ifres14990 = True


} else {
ifres14990 = False


}

ifres14989 = ifres14990


} else {
ifres14989 = False


}

var ifres14988 Obj

if True == ifres14989 {
ifres14988 = True


} else {
ifres14988 = False


}

ifres14987 = ifres14988


} else {
ifres14987 = False


}

if True == ifres14987 {
tmp14896 := PrimTail(V5447)

tmp14897 := PrimHead(tmp14896)

tmp14898 := PrimTail(V5447)

tmp14899 := PrimTail(tmp14898)

tmp14900 := PrimHead(tmp14899)

tmp14901 := PrimTail(tmp14900)

tmp14902 := PrimCons(tmp14901, Nil)

tmp14903 := PrimCons(symnot, tmp14902)

tmp14904 := PrimCons(tmp14903, Nil)

tmp14905 := PrimCons(tmp14897, tmp14904)

tmp14906 := PrimCons(symand, tmp14905)

tmp14907 := PrimTail(V5447)

tmp14908 := PrimTail(tmp14907)

tmp14909 := PrimHead(tmp14908)

tmp14910 := PrimTail(tmp14909)

tmp14911 := PrimTail(tmp14910)

tmp14912 := PrimCons(tmp14906, tmp14911)

__e.Return(PrimCons(symwhere, tmp14912))
return


} else {
tmp14985 := PrimIsPair(V5447)

var ifres14966 Obj

if True == tmp14985 {
tmp14983 := PrimHead(V5447)

tmp14984 := PrimEqual(symwhere, tmp14983)

var ifres14968 Obj

if True == tmp14984 {
tmp14981 := PrimTail(V5447)

tmp14982 := PrimIsPair(tmp14981)

var ifres14970 Obj

if True == tmp14982 {
tmp14978 := PrimTail(V5447)

tmp14979 := PrimTail(tmp14978)

tmp14980 := PrimIsPair(tmp14979)

var ifres14972 Obj

if True == tmp14980 {
tmp14974 := PrimTail(V5447)

tmp14975 := PrimTail(tmp14974)

tmp14976 := PrimTail(tmp14975)

tmp14977 := PrimEqual(Nil, tmp14976)

var ifres14973 Obj

if True == tmp14977 {
ifres14973 = True


} else {
ifres14973 = False


}

ifres14972 = ifres14973


} else {
ifres14972 = False


}

var ifres14971 Obj

if True == ifres14972 {
ifres14971 = True


} else {
ifres14971 = False


}

ifres14970 = ifres14971


} else {
ifres14970 = False


}

var ifres14969 Obj

if True == ifres14970 {
ifres14969 = True


} else {
ifres14969 = False


}

ifres14968 = ifres14969


} else {
ifres14968 = False


}

var ifres14967 Obj

if True == ifres14968 {
ifres14967 = True


} else {
ifres14967 = False


}

ifres14966 = ifres14967


} else {
ifres14966 = False


}

if True == ifres14966 {
tmp14913 := PrimTail(V5447)

tmp14914 := PrimHead(tmp14913)

tmp14915 := PrimTail(V5447)

tmp14916 := PrimTail(tmp14915)

tmp14917 := PrimHead(tmp14916)

tmp14918 := PrimCons(symfail, Nil)

tmp14919 := PrimCons(tmp14918, Nil)

tmp14920 := PrimCons(tmp14917, tmp14919)

tmp14921 := PrimCons(sym_a, tmp14920)

tmp14922 := PrimCons(tmp14921, Nil)

tmp14923 := PrimCons(symnot, tmp14922)

tmp14924 := PrimCons(tmp14923, Nil)

tmp14925 := PrimCons(tmp14914, tmp14924)

tmp14926 := PrimCons(symand, tmp14925)

tmp14927 := PrimTail(V5447)

tmp14928 := PrimTail(tmp14927)

tmp14929 := PrimCons(tmp14926, tmp14928)

__e.Return(PrimCons(symwhere, tmp14929))
return


} else {
tmp14964 := PrimIsPair(V5447)

var ifres14945 Obj

if True == tmp14964 {
tmp14962 := PrimHead(V5447)

tmp14963 := PrimEqual(symfail_1if, tmp14962)

var ifres14947 Obj

if True == tmp14963 {
tmp14960 := PrimTail(V5447)

tmp14961 := PrimIsPair(tmp14960)

var ifres14949 Obj

if True == tmp14961 {
tmp14957 := PrimTail(V5447)

tmp14958 := PrimTail(tmp14957)

tmp14959 := PrimIsPair(tmp14958)

var ifres14951 Obj

if True == tmp14959 {
tmp14953 := PrimTail(V5447)

tmp14954 := PrimTail(tmp14953)

tmp14955 := PrimTail(tmp14954)

tmp14956 := PrimEqual(Nil, tmp14955)

var ifres14952 Obj

if True == tmp14956 {
ifres14952 = True


} else {
ifres14952 = False


}

ifres14951 = ifres14952


} else {
ifres14951 = False


}

var ifres14950 Obj

if True == ifres14951 {
ifres14950 = True


} else {
ifres14950 = False


}

ifres14949 = ifres14950


} else {
ifres14949 = False


}

var ifres14948 Obj

if True == ifres14949 {
ifres14948 = True


} else {
ifres14948 = False


}

ifres14947 = ifres14948


} else {
ifres14947 = False


}

var ifres14946 Obj

if True == ifres14947 {
ifres14946 = True


} else {
ifres14946 = False


}

ifres14945 = ifres14946


} else {
ifres14945 = False


}

if True == ifres14945 {
tmp14930 := PrimTail(V5447)

tmp14931 := PrimCons(tmp14930, Nil)

tmp14932 := PrimCons(symnot, tmp14931)

tmp14933 := PrimTail(V5447)

tmp14934 := PrimTail(tmp14933)

tmp14935 := PrimCons(tmp14932, tmp14934)

__e.Return(PrimCons(symwhere, tmp14935))
return


} else {
tmp14936 := PrimCons(symfail, Nil)

tmp14937 := PrimCons(tmp14936, Nil)

tmp14938 := PrimCons(V5447, tmp14937)

tmp14939 := PrimCons(sym_a, tmp14938)

tmp14940 := PrimCons(tmp14939, Nil)

tmp14941 := PrimCons(symnot, tmp14940)

tmp14942 := PrimCons(V5447, Nil)

tmp14943 := PrimCons(tmp14941, tmp14942)

__e.Return(PrimCons(symwhere, tmp14943))
return


}


}


}


}, 1)

tmp15044 := Call(__e, ns2_1set, symshen_4correct, tmp14895)


_ = tmp15044

tmp15045 := MakeNative(func(__e *ControlFlow) {
V5448 := __e.Get(1)
_ = V5448
V5449 := __e.Get(2)
_ = V5449
V5450 := __e.Get(3)
_ = V5450
V5451 := __e.Get(4)
_ = V5451
V5452 := __e.Get(5)
_ = V5452
V5453 := __e.Get(6)
_ = V5453
V5454 := __e.Get(7)
_ = V5454
V5455 := __e.Get(8)
_ = V5455
tmp15046 := MakeNative(func(__e *ControlFlow) {
W5456 := __e.Get(1)
_ = W5456
tmp15047 := MakeNative(func(__e *ControlFlow) {
W5457 := __e.Get(1)
_ = W5457
tmp15077 := PrimEqual(W5457, False)

if True == tmp15077 {
tmp15048 := MakeNative(func(__e *ControlFlow) {
W5459 := __e.Get(1)
_ = W5459
tmp15050 := PrimEqual(W5459, False)

if True == tmp15050 {
__e.TailApply(PrimFunc(symshen_4unlock), V5453, W5456)
return
} else {
__e.Return(W5459)
return
}


}, 1)

tmp15075 := Call(__e, PrimFunc(symshen_4unlocked_2), V5453)


var ifres15051 Obj

if True == tmp15075 {
tmp15052 := MakeNative(func(__e *ControlFlow) {
W5460 := __e.Get(1)
_ = W5460
tmp15072 := PrimIsPair(W5460)

if True == tmp15072 {
tmp15053 := MakeNative(func(__e *ControlFlow) {
W5461 := __e.Get(1)
_ = W5461
tmp15054 := MakeNative(func(__e *ControlFlow) {
W5462 := __e.Get(1)
_ = W5462
tmp15055 := MakeNative(func(__e *ControlFlow) {
W5463 := __e.Get(1)
_ = W5463
tmp15056 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15056

tmp15057 := Call(__e, PrimFunc(symshen_4deref), W5461, V5452)


tmp15058 := Call(__e, PrimFunc(symshen_4freshen_1rule), tmp15057)


tmp15059 := MakeNative(func(__e *ControlFlow) {
tmp15060 := Call(__e, PrimFunc(symshen_4lazyderef), W5463, V5452)


tmp15061 := Call(__e, PrimFunc(symfst), tmp15060)


tmp15062 := Call(__e, PrimFunc(symshen_4lazyderef), W5463, V5452)


tmp15063 := Call(__e, PrimFunc(symsnd), tmp15062)


tmp15064 := MakeNative(func(__e *ControlFlow) {
tmp15065 := MakeNative(func(__e *ControlFlow) {
tmp15066 := PrimNumberAdd(V5451, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4t_d_1rules), V5448, W5462, V5450, tmp15066, V5452, V5453, W5456, V5455)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5452, V5453, W5456, tmp15065)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1rule), V5448, V5451, tmp15061, tmp15063, V5450, V5452, V5453, W5456, tmp15064)
return


}, 0)

tmp15067 := Call(__e, PrimFunc(symbind), W5463, tmp15058, V5452, V5453, W5456, tmp15059)


__e.TailApply(PrimFunc(symshen_4gc), V5452, tmp15067)
return


}, 1)

tmp15068 := Call(__e, PrimFunc(symshen_4newpv), V5452)


__e.TailApply(tmp15055, tmp15068)
return


}, 1)

tmp15069 := PrimTail(W5460)

__e.TailApply(tmp15054, tmp15069)
return


}, 1)

tmp15070 := PrimHead(W5460)

__e.TailApply(tmp15053, tmp15070)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15073 := Call(__e, PrimFunc(symshen_4lazyderef), V5449, V5452)


tmp15074 := Call(__e, tmp15052, tmp15073)


ifres15051 = tmp15074


} else {
ifres15051 = False


}

__e.TailApply(tmp15048, ifres15051)
return


} else {
__e.Return(W5457)
return
}


}, 1)

tmp15085 := Call(__e, PrimFunc(symshen_4unlocked_2), V5453)


var ifres15078 Obj

if True == tmp15085 {
tmp15079 := MakeNative(func(__e *ControlFlow) {
W5458 := __e.Get(1)
_ = W5458
tmp15082 := PrimEqual(W5458, Nil)

if True == tmp15082 {
tmp15080 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15080

__e.TailApply(PrimFunc(symthaw), V5455)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15083 := Call(__e, PrimFunc(symshen_4lazyderef), V5449, V5452)


tmp15084 := Call(__e, tmp15079, tmp15083)


ifres15078 = tmp15084


} else {
ifres15078 = False


}

__e.TailApply(tmp15047, ifres15078)
return


}, 1)

tmp15086 := PrimNumberAdd(V5454, MakeNumber(1))

__e.TailApply(tmp15046, tmp15086)
return


}, 8)

tmp15087 := Call(__e, ns2_1set, symshen_4t_d_1rules, tmp15045)


_ = tmp15087

tmp15088 := MakeNative(func(__e *ControlFlow) {
V5464 := __e.Get(1)
_ = V5464
tmp15101 := Call(__e, PrimFunc(symtuple_2), V5464)


if True == tmp15101 {
tmp15089 := MakeNative(func(__e *ControlFlow) {
W5465 := __e.Get(1)
_ = W5465
tmp15090 := MakeNative(func(__e *ControlFlow) {
W5466 := __e.Get(1)
_ = W5466
tmp15091 := Call(__e, PrimFunc(symfst), V5464)


tmp15092 := Call(__e, PrimFunc(symshen_4freshen), W5466, tmp15091)


tmp15093 := Call(__e, PrimFunc(symsnd), V5464)


tmp15094 := Call(__e, PrimFunc(symshen_4freshen), W5466, tmp15093)


__e.TailApply(PrimFunc(sym_8p), tmp15092, tmp15094)
return


}, 1)

tmp15095 := MakeNative(func(__e *ControlFlow) {
Z5467 := __e.Get(1)
_ = Z5467
tmp15096 := Call(__e, PrimFunc(symshen_4freshterm), Z5467)


__e.Return(PrimCons(Z5467, tmp15096))
return


}, 1)

tmp15097 := Call(__e, PrimFunc(symmap), tmp15095, W5465)


__e.TailApply(tmp15090, tmp15097)
return


}, 1)

tmp15098 := Call(__e, PrimFunc(symfst), V5464)


tmp15099 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp15098)


__e.TailApply(tmp15089, tmp15099)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen_1rule)
return
}


}, 1)

tmp15102 := Call(__e, ns2_1set, symshen_4freshen_1rule, tmp15088)


_ = tmp15102

tmp15103 := MakeNative(func(__e *ControlFlow) {
V5468 := __e.Get(1)
_ = V5468
V5469 := __e.Get(2)
_ = V5469
tmp15117 := PrimEqual(Nil, V5468)

if True == tmp15117 {
__e.Return(V5469)
return
} else {
tmp15115 := PrimIsPair(V5468)

var ifres15111 Obj

if True == tmp15115 {
tmp15113 := PrimHead(V5468)

tmp15114 := PrimIsPair(tmp15113)

var ifres15112 Obj

if True == tmp15114 {
ifres15112 = True


} else {
ifres15112 = False


}

ifres15111 = ifres15112


} else {
ifres15111 = False


}

if True == ifres15111 {
tmp15104 := PrimTail(V5468)

tmp15105 := PrimHead(V5468)

tmp15106 := PrimHead(tmp15105)

tmp15107 := PrimHead(V5468)

tmp15108 := PrimTail(tmp15107)

tmp15109 := Call(__e, PrimFunc(symshen_4beta), tmp15106, tmp15108, V5469)


__e.TailApply(PrimFunc(symshen_4freshen), tmp15104, tmp15109)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshen)
return
}


}


}, 2)

tmp15118 := Call(__e, ns2_1set, symshen_4freshen, tmp15103)


_ = tmp15118

tmp15119 := MakeNative(func(__e *ControlFlow) {
V5470 := __e.Get(1)
_ = V5470
V5471 := __e.Get(2)
_ = V5471
V5472 := __e.Get(3)
_ = V5472
V5473 := __e.Get(4)
_ = V5473
V5474 := __e.Get(5)
_ = V5474
V5475 := __e.Get(6)
_ = V5475
V5476 := __e.Get(7)
_ = V5476
V5477 := __e.Get(8)
_ = V5477
V5478 := __e.Get(9)
_ = V5478
tmp15120 := MakeNative(func(__e *ControlFlow) {
W5479 := __e.Get(1)
_ = W5479
tmp15133 := PrimEqual(W5479, False)

if True == tmp15133 {
tmp15131 := Call(__e, PrimFunc(symshen_4unlocked_2), V5476)


if True == tmp15131 {
tmp15121 := MakeNative(func(__e *ControlFlow) {
W5480 := __e.Get(1)
_ = W5480
tmp15122 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15122

tmp15123 := Call(__e, PrimFunc(symshen_4app), V5470, MakeString("\n"), symshen_4a)


tmp15124 := PrimStringConcat(MakeString(" of "), tmp15123)

tmp15125 := Call(__e, PrimFunc(symshen_4app), V5471, tmp15124, symshen_4a)


tmp15126 := PrimStringConcat(MakeString("type error in rule "), tmp15125)

tmp15127 := PrimSimpleError(tmp15126)

tmp15128 := Call(__e, PrimFunc(symbind), W5480, tmp15127, V5475, V5476, V5477, V5478)


__e.TailApply(PrimFunc(symshen_4gc), V5475, tmp15128)
return


}, 1)

tmp15129 := Call(__e, PrimFunc(symshen_4newpv), V5475)


__e.TailApply(tmp15121, tmp15129)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5479)
return
}


}, 1)

tmp15137 := Call(__e, PrimFunc(symshen_4unlocked_2), V5476)


var ifres15134 Obj

if True == tmp15137 {
tmp15135 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15135

tmp15136 := Call(__e, PrimFunc(symshen_4t_d_1rule_1h), V5472, V5473, V5474, V5475, V5476, V5477, V5478)


ifres15134 = tmp15136


} else {
ifres15134 = False


}

__e.TailApply(tmp15120, ifres15134)
return


}, 9)

tmp15138 := Call(__e, ns2_1set, symshen_4t_d_1rule, tmp15119)


_ = tmp15138

tmp15139 := MakeNative(func(__e *ControlFlow) {
V5481 := __e.Get(1)
_ = V5481
V5482 := __e.Get(2)
_ = V5482
V5483 := __e.Get(3)
_ = V5483
V5484 := __e.Get(4)
_ = V5484
V5485 := __e.Get(5)
_ = V5485
V5486 := __e.Get(6)
_ = V5486
V5487 := __e.Get(7)
_ = V5487
tmp15140 := MakeNative(func(__e *ControlFlow) {
W5488 := __e.Get(1)
_ = W5488
tmp15141 := MakeNative(func(__e *ControlFlow) {
W5489 := __e.Get(1)
_ = W5489
tmp15164 := PrimEqual(W5489, False)

if True == tmp15164 {
tmp15142 := MakeNative(func(__e *ControlFlow) {
W5496 := __e.Get(1)
_ = W5496
tmp15144 := PrimEqual(W5496, False)

if True == tmp15144 {
__e.TailApply(PrimFunc(symshen_4unlock), V5485, W5488)
return
} else {
__e.Return(W5496)
return
}


}, 1)

tmp15162 := Call(__e, PrimFunc(symshen_4unlocked_2), V5485)


var ifres15145 Obj

if True == tmp15162 {
tmp15146 := MakeNative(func(__e *ControlFlow) {
W5497 := __e.Get(1)
_ = W5497
tmp15147 := MakeNative(func(__e *ControlFlow) {
W5498 := __e.Get(1)
_ = W5498
tmp15148 := MakeNative(func(__e *ControlFlow) {
W5499 := __e.Get(1)
_ = W5499
tmp15149 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15149

tmp15150 := Call(__e, PrimFunc(symshen_4freshterms), V5481)


tmp15151 := MakeNative(func(__e *ControlFlow) {
tmp15152 := MakeNative(func(__e *ControlFlow) {
tmp15153 := MakeNative(func(__e *ControlFlow) {
tmp15154 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5482, W5498, W5499, V5484, V5485, W5488, V5487)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4myassume), V5481, V5483, W5499, V5484, V5485, W5488, tmp15154)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5484, V5485, W5488, tmp15153)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4t_d_1integrity), V5481, V5483, W5497, W5498, V5484, V5485, W5488, tmp15152)
return


}, 0)

tmp15155 := Call(__e, PrimFunc(symshen_4p_1hyps), tmp15150, W5497, V5484, V5485, W5488, tmp15151)


__e.TailApply(PrimFunc(symshen_4gc), V5484, tmp15155)
return


}, 1)

tmp15156 := Call(__e, PrimFunc(symshen_4newpv), V5484)


tmp15157 := Call(__e, tmp15148, tmp15156)


__e.TailApply(PrimFunc(symshen_4gc), V5484, tmp15157)
return


}, 1)

tmp15158 := Call(__e, PrimFunc(symshen_4newpv), V5484)


tmp15159 := Call(__e, tmp15147, tmp15158)


__e.TailApply(PrimFunc(symshen_4gc), V5484, tmp15159)
return


}, 1)

tmp15160 := Call(__e, PrimFunc(symshen_4newpv), V5484)


tmp15161 := Call(__e, tmp15146, tmp15160)


ifres15145 = tmp15161


} else {
ifres15145 = False


}

__e.TailApply(tmp15142, ifres15145)
return


} else {
__e.Return(W5489)
return
}


}, 1)

tmp15194 := Call(__e, PrimFunc(symshen_4unlocked_2), V5485)


var ifres15165 Obj

if True == tmp15194 {
tmp15166 := MakeNative(func(__e *ControlFlow) {
W5490 := __e.Get(1)
_ = W5490
tmp15191 := PrimEqual(W5490, Nil)

if True == tmp15191 {
tmp15167 := MakeNative(func(__e *ControlFlow) {
W5491 := __e.Get(1)
_ = W5491
tmp15188 := PrimIsPair(W5491)

if True == tmp15188 {
tmp15168 := MakeNative(func(__e *ControlFlow) {
W5492 := __e.Get(1)
_ = W5492
tmp15184 := PrimEqual(W5492, sym_1_1_6)

if True == tmp15184 {
tmp15169 := MakeNative(func(__e *ControlFlow) {
W5493 := __e.Get(1)
_ = W5493
tmp15180 := PrimIsPair(W5493)

if True == tmp15180 {
tmp15170 := MakeNative(func(__e *ControlFlow) {
W5494 := __e.Get(1)
_ = W5494
tmp15171 := MakeNative(func(__e *ControlFlow) {
W5495 := __e.Get(1)
_ = W5495
tmp15175 := PrimEqual(W5495, Nil)

if True == tmp15175 {
tmp15172 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15172

tmp15173 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1correct), V5482, W5494, Nil, V5484, V5485, W5488, V5487)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5484, V5485, W5488, tmp15173)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15176 := PrimTail(W5493)

tmp15177 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15176, V5484)


__e.TailApply(tmp15171, tmp15177)
return


}, 1)

tmp15178 := PrimHead(W5493)

__e.TailApply(tmp15170, tmp15178)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15181 := PrimTail(W5491)

tmp15182 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15181, V5484)


__e.TailApply(tmp15169, tmp15182)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15185 := PrimHead(W5491)

tmp15186 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15185, V5484)


__e.TailApply(tmp15168, tmp15186)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15189 := Call(__e, PrimFunc(symshen_4lazyderef), V5483, V5484)


__e.TailApply(tmp15167, tmp15189)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15192 := Call(__e, PrimFunc(symshen_4lazyderef), V5481, V5484)


tmp15193 := Call(__e, tmp15166, tmp15192)


ifres15165 = tmp15193


} else {
ifres15165 = False


}

__e.TailApply(tmp15141, ifres15165)
return


}, 1)

tmp15195 := PrimNumberAdd(V5486, MakeNumber(1))

__e.TailApply(tmp15140, tmp15195)
return


}, 7)

tmp15196 := Call(__e, ns2_1set, symshen_4t_d_1rule_1h, tmp15139)


_ = tmp15196

tmp15197 := MakeNative(func(__e *ControlFlow) {
V5500 := __e.Get(1)
_ = V5500
V5501 := __e.Get(2)
_ = V5501
V5502 := __e.Get(3)
_ = V5502
V5503 := __e.Get(4)
_ = V5503
V5504 := __e.Get(5)
_ = V5504
V5505 := __e.Get(6)
_ = V5505
V5506 := __e.Get(7)
_ = V5506
tmp15198 := MakeNative(func(__e *ControlFlow) {
W5507 := __e.Get(1)
_ = W5507
tmp15351 := PrimEqual(W5507, False)

if True == tmp15351 {
tmp15349 := Call(__e, PrimFunc(symshen_4unlocked_2), V5504)


if True == tmp15349 {
tmp15199 := MakeNative(func(__e *ControlFlow) {
W5511 := __e.Get(1)
_ = W5511
tmp15346 := PrimIsPair(W5511)

if True == tmp15346 {
tmp15200 := MakeNative(func(__e *ControlFlow) {
W5512 := __e.Get(1)
_ = W5512
tmp15201 := MakeNative(func(__e *ControlFlow) {
W5513 := __e.Get(1)
_ = W5513
tmp15202 := MakeNative(func(__e *ControlFlow) {
W5514 := __e.Get(1)
_ = W5514
tmp15341 := PrimIsPair(W5514)

if True == tmp15341 {
tmp15203 := MakeNative(func(__e *ControlFlow) {
W5515 := __e.Get(1)
_ = W5515
tmp15204 := MakeNative(func(__e *ControlFlow) {
W5516 := __e.Get(1)
_ = W5516
tmp15336 := PrimIsPair(W5516)

if True == tmp15336 {
tmp15205 := MakeNative(func(__e *ControlFlow) {
W5517 := __e.Get(1)
_ = W5517
tmp15332 := PrimEqual(W5517, sym_1_1_6)

if True == tmp15332 {
tmp15206 := MakeNative(func(__e *ControlFlow) {
W5518 := __e.Get(1)
_ = W5518
tmp15328 := PrimIsPair(W5518)

if True == tmp15328 {
tmp15207 := MakeNative(func(__e *ControlFlow) {
W5519 := __e.Get(1)
_ = W5519
tmp15208 := MakeNative(func(__e *ControlFlow) {
W5520 := __e.Get(1)
_ = W5520
tmp15323 := PrimEqual(W5520, Nil)

if True == tmp15323 {
tmp15209 := MakeNative(func(__e *ControlFlow) {
W5521 := __e.Get(1)
_ = W5521
tmp15210 := MakeNative(func(__e *ControlFlow) {
W5522 := __e.Get(1)
_ = W5522
tmp15314 := PrimIsPair(W5521)

if True == tmp15314 {
tmp15211 := MakeNative(func(__e *ControlFlow) {
W5527 := __e.Get(1)
_ = W5527
tmp15212 := MakeNative(func(__e *ControlFlow) {
W5528 := __e.Get(1)
_ = W5528
tmp15282 := PrimIsPair(W5527)

if True == tmp15282 {
tmp15213 := MakeNative(func(__e *ControlFlow) {
W5533 := __e.Get(1)
_ = W5533
tmp15214 := MakeNative(func(__e *ControlFlow) {
W5534 := __e.Get(1)
_ = W5534
tmp15215 := MakeNative(func(__e *ControlFlow) {
W5535 := __e.Get(1)
_ = W5535
tmp15257 := PrimIsPair(W5534)

if True == tmp15257 {
tmp15216 := MakeNative(func(__e *ControlFlow) {
W5538 := __e.Get(1)
_ = W5538
tmp15217 := MakeNative(func(__e *ControlFlow) {
W5539 := __e.Get(1)
_ = W5539
tmp15218 := MakeNative(func(__e *ControlFlow) {
W5540 := __e.Get(1)
_ = W5540
tmp15238 := PrimIsPair(W5539)

if True == tmp15238 {
tmp15219 := MakeNative(func(__e *ControlFlow) {
W5542 := __e.Get(1)
_ = W5542
tmp15220 := MakeNative(func(__e *ControlFlow) {
W5543 := __e.Get(1)
_ = W5543
tmp15221 := MakeNative(func(__e *ControlFlow) {
W5544 := __e.Get(1)
_ = W5544
tmp15225 := PrimEqual(W5543, Nil)

if True == tmp15225 {
__e.TailApply(PrimFunc(symthaw), W5544)
return
} else {
tmp15223 := Call(__e, PrimFunc(symshen_4pvar_2), W5543)


if True == tmp15223 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5543, Nil, V5503, W5544)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15226 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5540, W5542)
return
}, 0)

__e.TailApply(tmp15221, tmp15226)
return


}, 1)

tmp15227 := PrimTail(W5539)

tmp15228 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15227, V5503)


__e.TailApply(tmp15220, tmp15228)
return


}, 1)

tmp15229 := PrimHead(W5539)

__e.TailApply(tmp15219, tmp15229)
return


} else {
tmp15236 := Call(__e, PrimFunc(symshen_4pvar_2), W5539)


if True == tmp15236 {
tmp15230 := MakeNative(func(__e *ControlFlow) {
W5545 := __e.Get(1)
_ = W5545
tmp15231 := PrimCons(W5545, Nil)

tmp15232 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5540, W5545)
return
}, 0)

tmp15233 := Call(__e, PrimFunc(symshen_4bind_b), W5539, tmp15231, V5503, tmp15232)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15233)
return


}, 1)

tmp15234 := Call(__e, PrimFunc(symshen_4newpv), V5503)


__e.TailApply(tmp15230, tmp15234)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15239 := MakeNative(func(__e *ControlFlow) {
Z5541 := __e.Get(1)
_ = Z5541
tmp15240 := Call(__e, W5535, W5538)


__e.TailApply(tmp15240, Z5541)
return


}, 1)

__e.TailApply(tmp15218, tmp15239)
return


}, 1)

tmp15241 := PrimTail(W5534)

tmp15242 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15241, V5503)


__e.TailApply(tmp15217, tmp15242)
return


}, 1)

tmp15243 := PrimHead(W5534)

__e.TailApply(tmp15216, tmp15243)
return


} else {
tmp15255 := Call(__e, PrimFunc(symshen_4pvar_2), W5534)


if True == tmp15255 {
tmp15244 := MakeNative(func(__e *ControlFlow) {
W5546 := __e.Get(1)
_ = W5546
tmp15245 := MakeNative(func(__e *ControlFlow) {
W5547 := __e.Get(1)
_ = W5547
tmp15246 := PrimCons(W5547, Nil)

tmp15247 := PrimCons(W5546, tmp15246)

tmp15248 := MakeNative(func(__e *ControlFlow) {
tmp15249 := Call(__e, W5535, W5546)


__e.TailApply(tmp15249, W5547)
return


}, 0)

tmp15250 := Call(__e, PrimFunc(symshen_4bind_b), W5534, tmp15247, V5503, tmp15248)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15250)
return


}, 1)

tmp15251 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15252 := Call(__e, tmp15245, tmp15251)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15252)
return


}, 1)

tmp15253 := Call(__e, PrimFunc(symshen_4newpv), V5503)


__e.TailApply(tmp15244, tmp15253)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15258 := MakeNative(func(__e *ControlFlow) {
Z5536 := __e.Get(1)
_ = Z5536
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5537 := __e.Get(1)
_ = Z5537
tmp15259 := Call(__e, W5528, W5533)


tmp15260 := Call(__e, tmp15259, Z5536)


__e.TailApply(tmp15260, Z5537)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15215, tmp15258)
return


}, 1)

tmp15261 := PrimTail(W5527)

tmp15262 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15261, V5503)


__e.TailApply(tmp15214, tmp15262)
return


}, 1)

tmp15263 := PrimHead(W5527)

__e.TailApply(tmp15213, tmp15263)
return


} else {
tmp15280 := Call(__e, PrimFunc(symshen_4pvar_2), W5527)


if True == tmp15280 {
tmp15264 := MakeNative(func(__e *ControlFlow) {
W5548 := __e.Get(1)
_ = W5548
tmp15265 := MakeNative(func(__e *ControlFlow) {
W5549 := __e.Get(1)
_ = W5549
tmp15266 := MakeNative(func(__e *ControlFlow) {
W5550 := __e.Get(1)
_ = W5550
tmp15267 := PrimCons(W5550, Nil)

tmp15268 := PrimCons(W5549, tmp15267)

tmp15269 := PrimCons(W5548, tmp15268)

tmp15270 := MakeNative(func(__e *ControlFlow) {
tmp15271 := Call(__e, W5528, W5548)


tmp15272 := Call(__e, tmp15271, W5549)


__e.TailApply(tmp15272, W5550)
return


}, 0)

tmp15273 := Call(__e, PrimFunc(symshen_4bind_b), W5527, tmp15269, V5503, tmp15270)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15273)
return


}, 1)

tmp15274 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15275 := Call(__e, tmp15266, tmp15274)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15275)
return


}, 1)

tmp15276 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15277 := Call(__e, tmp15265, tmp15276)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15277)
return


}, 1)

tmp15278 := Call(__e, PrimFunc(symshen_4newpv), V5503)


__e.TailApply(tmp15264, tmp15278)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15283 := MakeNative(func(__e *ControlFlow) {
Z5529 := __e.Get(1)
_ = Z5529
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5530 := __e.Get(1)
_ = Z5530
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5531 := __e.Get(1)
_ = Z5531
tmp15284 := MakeNative(func(__e *ControlFlow) {
W5532 := __e.Get(1)
_ = W5532
tmp15285 := Call(__e, W5522, Z5529)


tmp15286 := Call(__e, tmp15285, Z5530)


tmp15287 := Call(__e, tmp15286, Z5531)


__e.TailApply(tmp15287, W5532)
return


}, 1)

tmp15288 := PrimTail(W5521)

__e.TailApply(tmp15284, tmp15288)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp15212, tmp15283)
return


}, 1)

tmp15289 := PrimHead(W5521)

tmp15290 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15289, V5503)


__e.TailApply(tmp15211, tmp15290)
return


} else {
tmp15312 := Call(__e, PrimFunc(symshen_4pvar_2), W5521)


if True == tmp15312 {
tmp15291 := MakeNative(func(__e *ControlFlow) {
W5551 := __e.Get(1)
_ = W5551
tmp15292 := MakeNative(func(__e *ControlFlow) {
W5552 := __e.Get(1)
_ = W5552
tmp15293 := MakeNative(func(__e *ControlFlow) {
W5553 := __e.Get(1)
_ = W5553
tmp15294 := MakeNative(func(__e *ControlFlow) {
W5554 := __e.Get(1)
_ = W5554
tmp15295 := PrimCons(W5553, Nil)

tmp15296 := PrimCons(W5552, tmp15295)

tmp15297 := PrimCons(W5551, tmp15296)

tmp15298 := PrimCons(tmp15297, W5554)

tmp15299 := MakeNative(func(__e *ControlFlow) {
tmp15300 := Call(__e, W5522, W5551)


tmp15301 := Call(__e, tmp15300, W5552)


tmp15302 := Call(__e, tmp15301, W5553)


__e.TailApply(tmp15302, W5554)
return


}, 0)

tmp15303 := Call(__e, PrimFunc(symshen_4bind_b), W5521, tmp15298, V5503, tmp15299)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15303)
return


}, 1)

tmp15304 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15305 := Call(__e, tmp15294, tmp15304)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15305)
return


}, 1)

tmp15306 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15307 := Call(__e, tmp15293, tmp15306)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15307)
return


}, 1)

tmp15308 := Call(__e, PrimFunc(symshen_4newpv), V5503)


tmp15309 := Call(__e, tmp15292, tmp15308)


__e.TailApply(PrimFunc(symshen_4gc), V5503, tmp15309)
return


}, 1)

tmp15310 := Call(__e, PrimFunc(symshen_4newpv), V5503)


__e.TailApply(tmp15291, tmp15310)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15315 := MakeNative(func(__e *ControlFlow) {
Z5523 := __e.Get(1)
_ = Z5523
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5524 := __e.Get(1)
_ = Z5524
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5525 := __e.Get(1)
_ = Z5525
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5526 := __e.Get(1)
_ = Z5526
tmp15316 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15316

tmp15317 := MakeNative(func(__e *ControlFlow) {
tmp15318 := MakeNative(func(__e *ControlFlow) {
tmp15319 := PrimIntern(MakeString(":"))

tmp15320 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4myassume), W5513, W5519, Z5526, V5503, V5504, V5505, V5506)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5524, tmp15319, V5503, V5504, V5505, tmp15320)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5512, Z5523, V5503, V5504, V5505, tmp15318)
return


}, 0)

__e.TailApply(PrimFunc(symis_b), W5515, Z5525, V5503, V5504, V5505, tmp15317)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp15210, tmp15315)
return


}, 1)

tmp15321 := Call(__e, PrimFunc(symshen_4lazyderef), V5502, V5503)


__e.TailApply(tmp15209, tmp15321)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15324 := PrimTail(W5518)

tmp15325 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15324, V5503)


__e.TailApply(tmp15208, tmp15325)
return


}, 1)

tmp15326 := PrimHead(W5518)

__e.TailApply(tmp15207, tmp15326)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15329 := PrimTail(W5516)

tmp15330 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15329, V5503)


__e.TailApply(tmp15206, tmp15330)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15333 := PrimHead(W5516)

tmp15334 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15333, V5503)


__e.TailApply(tmp15205, tmp15334)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15337 := PrimTail(W5514)

tmp15338 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15337, V5503)


__e.TailApply(tmp15204, tmp15338)
return


}, 1)

tmp15339 := PrimHead(W5514)

__e.TailApply(tmp15203, tmp15339)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15342 := Call(__e, PrimFunc(symshen_4lazyderef), V5501, V5503)


__e.TailApply(tmp15202, tmp15342)
return


}, 1)

tmp15343 := PrimTail(W5511)

__e.TailApply(tmp15201, tmp15343)
return


}, 1)

tmp15344 := PrimHead(W5511)

__e.TailApply(tmp15200, tmp15344)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15347 := Call(__e, PrimFunc(symshen_4lazyderef), V5500, V5503)


__e.TailApply(tmp15199, tmp15347)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5507)
return
}


}, 1)

tmp15367 := Call(__e, PrimFunc(symshen_4unlocked_2), V5504)


var ifres15352 Obj

if True == tmp15367 {
tmp15353 := MakeNative(func(__e *ControlFlow) {
W5508 := __e.Get(1)
_ = W5508
tmp15364 := PrimEqual(W5508, Nil)

if True == tmp15364 {
tmp15354 := MakeNative(func(__e *ControlFlow) {
W5509 := __e.Get(1)
_ = W5509
tmp15355 := MakeNative(func(__e *ControlFlow) {
W5510 := __e.Get(1)
_ = W5510
tmp15359 := PrimEqual(W5509, Nil)

if True == tmp15359 {
__e.TailApply(PrimFunc(symthaw), W5510)
return
} else {
tmp15357 := Call(__e, PrimFunc(symshen_4pvar_2), W5509)


if True == tmp15357 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5509, Nil, V5503, W5510)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15360 := MakeNative(func(__e *ControlFlow) {
tmp15361 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15361

__e.TailApply(PrimFunc(symthaw), V5506)
return


}, 0)

__e.TailApply(tmp15355, tmp15360)
return


}, 1)

tmp15362 := Call(__e, PrimFunc(symshen_4lazyderef), V5502, V5503)


__e.TailApply(tmp15354, tmp15362)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15365 := Call(__e, PrimFunc(symshen_4lazyderef), V5500, V5503)


tmp15366 := Call(__e, tmp15353, tmp15365)


ifres15352 = tmp15366


} else {
ifres15352 = False


}

__e.TailApply(tmp15198, ifres15352)
return


}, 7)

tmp15368 := Call(__e, ns2_1set, symshen_4myassume, tmp15197)


_ = tmp15368

tmp15369 := MakeNative(func(__e *ControlFlow) {
V5557 := __e.Get(1)
_ = V5557
tmp15392 := PrimEqual(Nil, V5557)

if True == tmp15392 {
__e.Return(Nil)
return
} else {
tmp15390 := PrimIsPair(V5557)

var ifres15386 Obj

if True == tmp15390 {
tmp15388 := PrimHead(V5557)

tmp15389 := PrimIsPair(tmp15388)

var ifres15387 Obj

if True == tmp15389 {
ifres15387 = True


} else {
ifres15387 = False


}

ifres15386 = ifres15387


} else {
ifres15386 = False


}

if True == ifres15386 {
tmp15370 := PrimHead(V5557)

tmp15371 := PrimTail(V5557)

tmp15372 := Call(__e, PrimFunc(symappend), tmp15370, tmp15371)


__e.TailApply(PrimFunc(symshen_4freshterms), tmp15372)
return


} else {
tmp15384 := PrimIsPair(V5557)

var ifres15380 Obj

if True == tmp15384 {
tmp15382 := PrimHead(V5557)

tmp15383 := Call(__e, PrimFunc(symshen_4freshterm_2), tmp15382)


var ifres15381 Obj

if True == tmp15383 {
ifres15381 = True


} else {
ifres15381 = False


}

ifres15380 = ifres15381


} else {
ifres15380 = False


}

if True == ifres15380 {
tmp15373 := PrimHead(V5557)

tmp15374 := PrimTail(V5557)

tmp15375 := Call(__e, PrimFunc(symshen_4freshterms), tmp15374)


__e.TailApply(PrimFunc(symadjoin), tmp15373, tmp15375)
return


} else {
tmp15378 := PrimIsPair(V5557)

if True == tmp15378 {
tmp15376 := PrimTail(V5557)

__e.TailApply(PrimFunc(symshen_4freshterms), tmp15376)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4freshterms)
return
}


}


}


}


}, 1)

tmp15393 := Call(__e, ns2_1set, symshen_4freshterms, tmp15369)


_ = tmp15393

tmp15394 := MakeNative(func(__e *ControlFlow) {
V5558 := __e.Get(1)
_ = V5558
V5559 := __e.Get(2)
_ = V5559
V5560 := __e.Get(3)
_ = V5560
V5561 := __e.Get(4)
_ = V5561
V5562 := __e.Get(5)
_ = V5562
V5563 := __e.Get(6)
_ = V5563
tmp15395 := MakeNative(func(__e *ControlFlow) {
W5564 := __e.Get(1)
_ = W5564
tmp15519 := PrimEqual(W5564, False)

if True == tmp15519 {
tmp15517 := Call(__e, PrimFunc(symshen_4unlocked_2), V5561)


if True == tmp15517 {
tmp15396 := MakeNative(func(__e *ControlFlow) {
W5568 := __e.Get(1)
_ = W5568
tmp15514 := PrimIsPair(W5568)

if True == tmp15514 {
tmp15397 := MakeNative(func(__e *ControlFlow) {
W5569 := __e.Get(1)
_ = W5569
tmp15398 := MakeNative(func(__e *ControlFlow) {
W5570 := __e.Get(1)
_ = W5570
tmp15399 := MakeNative(func(__e *ControlFlow) {
W5571 := __e.Get(1)
_ = W5571
tmp15400 := MakeNative(func(__e *ControlFlow) {
W5572 := __e.Get(1)
_ = W5572
tmp15504 := PrimIsPair(W5571)

if True == tmp15504 {
tmp15401 := MakeNative(func(__e *ControlFlow) {
W5577 := __e.Get(1)
_ = W5577
tmp15402 := MakeNative(func(__e *ControlFlow) {
W5578 := __e.Get(1)
_ = W5578
tmp15472 := PrimIsPair(W5577)

if True == tmp15472 {
tmp15403 := MakeNative(func(__e *ControlFlow) {
W5583 := __e.Get(1)
_ = W5583
tmp15404 := MakeNative(func(__e *ControlFlow) {
W5584 := __e.Get(1)
_ = W5584
tmp15405 := MakeNative(func(__e *ControlFlow) {
W5585 := __e.Get(1)
_ = W5585
tmp15447 := PrimIsPair(W5584)

if True == tmp15447 {
tmp15406 := MakeNative(func(__e *ControlFlow) {
W5588 := __e.Get(1)
_ = W5588
tmp15407 := MakeNative(func(__e *ControlFlow) {
W5589 := __e.Get(1)
_ = W5589
tmp15408 := MakeNative(func(__e *ControlFlow) {
W5590 := __e.Get(1)
_ = W5590
tmp15428 := PrimIsPair(W5589)

if True == tmp15428 {
tmp15409 := MakeNative(func(__e *ControlFlow) {
W5592 := __e.Get(1)
_ = W5592
tmp15410 := MakeNative(func(__e *ControlFlow) {
W5593 := __e.Get(1)
_ = W5593
tmp15411 := MakeNative(func(__e *ControlFlow) {
W5594 := __e.Get(1)
_ = W5594
tmp15415 := PrimEqual(W5593, Nil)

if True == tmp15415 {
__e.TailApply(PrimFunc(symthaw), W5594)
return
} else {
tmp15413 := Call(__e, PrimFunc(symshen_4pvar_2), W5593)


if True == tmp15413 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5593, Nil, V5560, W5594)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15416 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5590, W5592)
return
}, 0)

__e.TailApply(tmp15411, tmp15416)
return


}, 1)

tmp15417 := PrimTail(W5589)

tmp15418 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15417, V5560)


__e.TailApply(tmp15410, tmp15418)
return


}, 1)

tmp15419 := PrimHead(W5589)

__e.TailApply(tmp15409, tmp15419)
return


} else {
tmp15426 := Call(__e, PrimFunc(symshen_4pvar_2), W5589)


if True == tmp15426 {
tmp15420 := MakeNative(func(__e *ControlFlow) {
W5595 := __e.Get(1)
_ = W5595
tmp15421 := PrimCons(W5595, Nil)

tmp15422 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(W5590, W5595)
return
}, 0)

tmp15423 := Call(__e, PrimFunc(symshen_4bind_b), W5589, tmp15421, V5560, tmp15422)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15423)
return


}, 1)

tmp15424 := Call(__e, PrimFunc(symshen_4newpv), V5560)


__e.TailApply(tmp15420, tmp15424)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15429 := MakeNative(func(__e *ControlFlow) {
Z5591 := __e.Get(1)
_ = Z5591
tmp15430 := Call(__e, W5585, W5588)


__e.TailApply(tmp15430, Z5591)
return


}, 1)

__e.TailApply(tmp15408, tmp15429)
return


}, 1)

tmp15431 := PrimTail(W5584)

tmp15432 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15431, V5560)


__e.TailApply(tmp15407, tmp15432)
return


}, 1)

tmp15433 := PrimHead(W5584)

__e.TailApply(tmp15406, tmp15433)
return


} else {
tmp15445 := Call(__e, PrimFunc(symshen_4pvar_2), W5584)


if True == tmp15445 {
tmp15434 := MakeNative(func(__e *ControlFlow) {
W5596 := __e.Get(1)
_ = W5596
tmp15435 := MakeNative(func(__e *ControlFlow) {
W5597 := __e.Get(1)
_ = W5597
tmp15436 := PrimCons(W5597, Nil)

tmp15437 := PrimCons(W5596, tmp15436)

tmp15438 := MakeNative(func(__e *ControlFlow) {
tmp15439 := Call(__e, W5585, W5596)


__e.TailApply(tmp15439, W5597)
return


}, 0)

tmp15440 := Call(__e, PrimFunc(symshen_4bind_b), W5584, tmp15437, V5560, tmp15438)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15440)
return


}, 1)

tmp15441 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15442 := Call(__e, tmp15435, tmp15441)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15442)
return


}, 1)

tmp15443 := Call(__e, PrimFunc(symshen_4newpv), V5560)


__e.TailApply(tmp15434, tmp15443)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15448 := MakeNative(func(__e *ControlFlow) {
Z5586 := __e.Get(1)
_ = Z5586
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5587 := __e.Get(1)
_ = Z5587
tmp15449 := Call(__e, W5578, W5583)


tmp15450 := Call(__e, tmp15449, Z5586)


__e.TailApply(tmp15450, Z5587)
return


}, 1))
return
}, 1)

__e.TailApply(tmp15405, tmp15448)
return


}, 1)

tmp15451 := PrimTail(W5577)

tmp15452 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15451, V5560)


__e.TailApply(tmp15404, tmp15452)
return


}, 1)

tmp15453 := PrimHead(W5577)

__e.TailApply(tmp15403, tmp15453)
return


} else {
tmp15470 := Call(__e, PrimFunc(symshen_4pvar_2), W5577)


if True == tmp15470 {
tmp15454 := MakeNative(func(__e *ControlFlow) {
W5598 := __e.Get(1)
_ = W5598
tmp15455 := MakeNative(func(__e *ControlFlow) {
W5599 := __e.Get(1)
_ = W5599
tmp15456 := MakeNative(func(__e *ControlFlow) {
W5600 := __e.Get(1)
_ = W5600
tmp15457 := PrimCons(W5600, Nil)

tmp15458 := PrimCons(W5599, tmp15457)

tmp15459 := PrimCons(W5598, tmp15458)

tmp15460 := MakeNative(func(__e *ControlFlow) {
tmp15461 := Call(__e, W5578, W5598)


tmp15462 := Call(__e, tmp15461, W5599)


__e.TailApply(tmp15462, W5600)
return


}, 0)

tmp15463 := Call(__e, PrimFunc(symshen_4bind_b), W5577, tmp15459, V5560, tmp15460)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15463)
return


}, 1)

tmp15464 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15465 := Call(__e, tmp15456, tmp15464)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15465)
return


}, 1)

tmp15466 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15467 := Call(__e, tmp15455, tmp15466)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15467)
return


}, 1)

tmp15468 := Call(__e, PrimFunc(symshen_4newpv), V5560)


__e.TailApply(tmp15454, tmp15468)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15473 := MakeNative(func(__e *ControlFlow) {
Z5579 := __e.Get(1)
_ = Z5579
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5580 := __e.Get(1)
_ = Z5580
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5581 := __e.Get(1)
_ = Z5581
tmp15474 := MakeNative(func(__e *ControlFlow) {
W5582 := __e.Get(1)
_ = W5582
tmp15475 := Call(__e, W5572, Z5579)


tmp15476 := Call(__e, tmp15475, Z5580)


tmp15477 := Call(__e, tmp15476, Z5581)


__e.TailApply(tmp15477, W5582)
return


}, 1)

tmp15478 := PrimTail(W5571)

__e.TailApply(tmp15474, tmp15478)
return


}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp15402, tmp15473)
return


}, 1)

tmp15479 := PrimHead(W5571)

tmp15480 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15479, V5560)


__e.TailApply(tmp15401, tmp15480)
return


} else {
tmp15502 := Call(__e, PrimFunc(symshen_4pvar_2), W5571)


if True == tmp15502 {
tmp15481 := MakeNative(func(__e *ControlFlow) {
W5601 := __e.Get(1)
_ = W5601
tmp15482 := MakeNative(func(__e *ControlFlow) {
W5602 := __e.Get(1)
_ = W5602
tmp15483 := MakeNative(func(__e *ControlFlow) {
W5603 := __e.Get(1)
_ = W5603
tmp15484 := MakeNative(func(__e *ControlFlow) {
W5604 := __e.Get(1)
_ = W5604
tmp15485 := PrimCons(W5603, Nil)

tmp15486 := PrimCons(W5602, tmp15485)

tmp15487 := PrimCons(W5601, tmp15486)

tmp15488 := PrimCons(tmp15487, W5604)

tmp15489 := MakeNative(func(__e *ControlFlow) {
tmp15490 := Call(__e, W5572, W5601)


tmp15491 := Call(__e, tmp15490, W5602)


tmp15492 := Call(__e, tmp15491, W5603)


__e.TailApply(tmp15492, W5604)
return


}, 0)

tmp15493 := Call(__e, PrimFunc(symshen_4bind_b), W5571, tmp15488, V5560, tmp15489)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15493)
return


}, 1)

tmp15494 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15495 := Call(__e, tmp15484, tmp15494)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15495)
return


}, 1)

tmp15496 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15497 := Call(__e, tmp15483, tmp15496)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15497)
return


}, 1)

tmp15498 := Call(__e, PrimFunc(symshen_4newpv), V5560)


tmp15499 := Call(__e, tmp15482, tmp15498)


__e.TailApply(PrimFunc(symshen_4gc), V5560, tmp15499)
return


}, 1)

tmp15500 := Call(__e, PrimFunc(symshen_4newpv), V5560)


__e.TailApply(tmp15481, tmp15500)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp15505 := MakeNative(func(__e *ControlFlow) {
Z5573 := __e.Get(1)
_ = Z5573
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5574 := __e.Get(1)
_ = Z5574
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5575 := __e.Get(1)
_ = Z5575
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5576 := __e.Get(1)
_ = Z5576
tmp15506 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15506

tmp15507 := MakeNative(func(__e *ControlFlow) {
tmp15508 := PrimIntern(MakeString(":"))

tmp15509 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4p_1hyps), W5570, Z5576, V5560, V5561, V5562, V5563)
return
}, 0)

__e.TailApply(PrimFunc(symbind), Z5574, tmp15508, V5560, V5561, V5562, tmp15509)
return


}, 0)

__e.TailApply(PrimFunc(symbind), Z5573, W5569, V5560, V5561, V5562, tmp15507)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

__e.TailApply(tmp15400, tmp15505)
return


}, 1)

tmp15510 := Call(__e, PrimFunc(symshen_4lazyderef), V5559, V5560)


__e.TailApply(tmp15399, tmp15510)
return


}, 1)

tmp15511 := PrimTail(W5568)

__e.TailApply(tmp15398, tmp15511)
return


}, 1)

tmp15512 := PrimHead(W5568)

__e.TailApply(tmp15397, tmp15512)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15515 := Call(__e, PrimFunc(symshen_4lazyderef), V5558, V5560)


__e.TailApply(tmp15396, tmp15515)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5564)
return
}


}, 1)

tmp15535 := Call(__e, PrimFunc(symshen_4unlocked_2), V5561)


var ifres15520 Obj

if True == tmp15535 {
tmp15521 := MakeNative(func(__e *ControlFlow) {
W5565 := __e.Get(1)
_ = W5565
tmp15532 := PrimEqual(W5565, Nil)

if True == tmp15532 {
tmp15522 := MakeNative(func(__e *ControlFlow) {
W5566 := __e.Get(1)
_ = W5566
tmp15523 := MakeNative(func(__e *ControlFlow) {
W5567 := __e.Get(1)
_ = W5567
tmp15527 := PrimEqual(W5566, Nil)

if True == tmp15527 {
__e.TailApply(PrimFunc(symthaw), W5567)
return
} else {
tmp15525 := Call(__e, PrimFunc(symshen_4pvar_2), W5566)


if True == tmp15525 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5566, Nil, V5560, W5567)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp15528 := MakeNative(func(__e *ControlFlow) {
tmp15529 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15529

__e.TailApply(PrimFunc(symthaw), V5563)
return


}, 0)

__e.TailApply(tmp15523, tmp15528)
return


}, 1)

tmp15530 := Call(__e, PrimFunc(symshen_4lazyderef), V5559, V5560)


__e.TailApply(tmp15522, tmp15530)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15533 := Call(__e, PrimFunc(symshen_4lazyderef), V5558, V5560)


tmp15534 := Call(__e, tmp15521, tmp15533)


ifres15520 = tmp15534


} else {
ifres15520 = False


}

__e.TailApply(tmp15395, ifres15520)
return


}, 6)

tmp15536 := Call(__e, ns2_1set, symshen_4p_1hyps, tmp15394)


_ = tmp15536

tmp15537 := MakeNative(func(__e *ControlFlow) {
V5605 := __e.Get(1)
_ = V5605
V5606 := __e.Get(2)
_ = V5606
V5607 := __e.Get(3)
_ = V5607
V5608 := __e.Get(4)
_ = V5608
V5609 := __e.Get(5)
_ = V5609
V5610 := __e.Get(6)
_ = V5610
V5611 := __e.Get(7)
_ = V5611
tmp15538 := MakeNative(func(__e *ControlFlow) {
W5612 := __e.Get(1)
_ = W5612
tmp15539 := MakeNative(func(__e *ControlFlow) {
W5613 := __e.Get(1)
_ = W5613
tmp15549 := PrimEqual(W5613, False)

if True == tmp15549 {
tmp15540 := MakeNative(func(__e *ControlFlow) {
W5622 := __e.Get(1)
_ = W5622
tmp15542 := PrimEqual(W5622, False)

if True == tmp15542 {
__e.TailApply(PrimFunc(symshen_4unlock), V5609, W5612)
return
} else {
__e.Return(W5622)
return
}


}, 1)

tmp15547 := Call(__e, PrimFunc(symshen_4unlocked_2), V5609)


var ifres15543 Obj

if True == tmp15547 {
tmp15544 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15544

tmp15545 := Call(__e, PrimFunc(symshen_4curry), V5605)


tmp15546 := Call(__e, PrimFunc(symshen_4system_1S_1h), tmp15545, V5606, V5607, V5608, V5609, W5612, V5611)


ifres15543 = tmp15546


} else {
ifres15543 = False


}

__e.TailApply(tmp15540, ifres15543)
return


} else {
__e.Return(W5613)
return
}


}, 1)

tmp15594 := Call(__e, PrimFunc(symshen_4unlocked_2), V5609)


var ifres15550 Obj

if True == tmp15594 {
tmp15551 := MakeNative(func(__e *ControlFlow) {
W5614 := __e.Get(1)
_ = W5614
tmp15591 := PrimIsPair(W5614)

if True == tmp15591 {
tmp15552 := MakeNative(func(__e *ControlFlow) {
W5615 := __e.Get(1)
_ = W5615
tmp15587 := PrimEqual(W5615, symwhere)

if True == tmp15587 {
tmp15553 := MakeNative(func(__e *ControlFlow) {
W5616 := __e.Get(1)
_ = W5616
tmp15583 := PrimIsPair(W5616)

if True == tmp15583 {
tmp15554 := MakeNative(func(__e *ControlFlow) {
W5617 := __e.Get(1)
_ = W5617
tmp15555 := MakeNative(func(__e *ControlFlow) {
W5618 := __e.Get(1)
_ = W5618
tmp15578 := PrimIsPair(W5618)

if True == tmp15578 {
tmp15556 := MakeNative(func(__e *ControlFlow) {
W5619 := __e.Get(1)
_ = W5619
tmp15557 := MakeNative(func(__e *ControlFlow) {
W5620 := __e.Get(1)
_ = W5620
tmp15573 := PrimEqual(W5620, Nil)

if True == tmp15573 {
tmp15558 := MakeNative(func(__e *ControlFlow) {
W5621 := __e.Get(1)
_ = W5621
tmp15559 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15559

tmp15560 := MakeNative(func(__e *ControlFlow) {
tmp15561 := Call(__e, PrimFunc(symshen_4curry), W5617)


tmp15562 := MakeNative(func(__e *ControlFlow) {
tmp15563 := MakeNative(func(__e *ControlFlow) {
tmp15564 := MakeNative(func(__e *ControlFlow) {
tmp15565 := PrimIntern(MakeString(":"))

tmp15566 := PrimCons(symverified, Nil)

tmp15567 := PrimCons(tmp15565, tmp15566)

tmp15568 := PrimCons(W5621, tmp15567)

tmp15569 := PrimCons(tmp15568, V5607)

__e.TailApply(PrimFunc(symshen_4t_d_1correct), W5619, V5606, tmp15569, V5608, V5609, W5612, V5611)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4cut), V5608, V5609, W5612, tmp15564)
return


}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5621, symboolean, V5607, V5608, V5609, W5612, tmp15563)
return


}, 0)

__e.TailApply(PrimFunc(symbind), W5621, tmp15561, V5608, V5609, W5612, tmp15562)
return


}, 0)

tmp15570 := Call(__e, PrimFunc(symshen_4cut), V5608, V5609, W5612, tmp15560)


__e.TailApply(PrimFunc(symshen_4gc), V5608, tmp15570)
return


}, 1)

tmp15571 := Call(__e, PrimFunc(symshen_4newpv), V5608)


__e.TailApply(tmp15558, tmp15571)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15574 := PrimTail(W5618)

tmp15575 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15574, V5608)


__e.TailApply(tmp15557, tmp15575)
return


}, 1)

tmp15576 := PrimHead(W5618)

__e.TailApply(tmp15556, tmp15576)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15579 := PrimTail(W5616)

tmp15580 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15579, V5608)


__e.TailApply(tmp15555, tmp15580)
return


}, 1)

tmp15581 := PrimHead(W5616)

__e.TailApply(tmp15554, tmp15581)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15584 := PrimTail(W5614)

tmp15585 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15584, V5608)


__e.TailApply(tmp15553, tmp15585)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15588 := PrimHead(W5614)

tmp15589 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15588, V5608)


__e.TailApply(tmp15552, tmp15589)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15592 := Call(__e, PrimFunc(symshen_4lazyderef), V5605, V5608)


tmp15593 := Call(__e, tmp15551, tmp15592)


ifres15550 = tmp15593


} else {
ifres15550 = False


}

__e.TailApply(tmp15539, ifres15550)
return


}, 1)

tmp15595 := PrimNumberAdd(V5610, MakeNumber(1))

__e.TailApply(tmp15538, tmp15595)
return


}, 7)

tmp15596 := Call(__e, ns2_1set, symshen_4t_d_1correct, tmp15537)


_ = tmp15596

tmp15597 := MakeNative(func(__e *ControlFlow) {
V5623 := __e.Get(1)
_ = V5623
V5624 := __e.Get(2)
_ = V5624
V5625 := __e.Get(3)
_ = V5625
V5626 := __e.Get(4)
_ = V5626
V5627 := __e.Get(5)
_ = V5627
V5628 := __e.Get(6)
_ = V5628
V5629 := __e.Get(7)
_ = V5629
V5630 := __e.Get(8)
_ = V5630
tmp15598 := MakeNative(func(__e *ControlFlow) {
W5631 := __e.Get(1)
_ = W5631
tmp15640 := PrimEqual(W5631, False)

if True == tmp15640 {
tmp15638 := Call(__e, PrimFunc(symshen_4unlocked_2), V5628)


if True == tmp15638 {
tmp15599 := MakeNative(func(__e *ControlFlow) {
W5633 := __e.Get(1)
_ = W5633
tmp15635 := PrimIsPair(W5633)

if True == tmp15635 {
tmp15600 := MakeNative(func(__e *ControlFlow) {
W5634 := __e.Get(1)
_ = W5634
tmp15601 := MakeNative(func(__e *ControlFlow) {
W5635 := __e.Get(1)
_ = W5635
tmp15602 := MakeNative(func(__e *ControlFlow) {
W5636 := __e.Get(1)
_ = W5636
tmp15630 := PrimIsPair(W5636)

if True == tmp15630 {
tmp15603 := MakeNative(func(__e *ControlFlow) {
W5637 := __e.Get(1)
_ = W5637
tmp15604 := MakeNative(func(__e *ControlFlow) {
W5638 := __e.Get(1)
_ = W5638
tmp15625 := PrimIsPair(W5638)

if True == tmp15625 {
tmp15605 := MakeNative(func(__e *ControlFlow) {
W5639 := __e.Get(1)
_ = W5639
tmp15621 := PrimEqual(W5639, sym_1_1_6)

if True == tmp15621 {
tmp15606 := MakeNative(func(__e *ControlFlow) {
W5640 := __e.Get(1)
_ = W5640
tmp15617 := PrimIsPair(W5640)

if True == tmp15617 {
tmp15607 := MakeNative(func(__e *ControlFlow) {
W5641 := __e.Get(1)
_ = W5641
tmp15608 := MakeNative(func(__e *ControlFlow) {
W5642 := __e.Get(1)
_ = W5642
tmp15612 := PrimEqual(W5642, Nil)

if True == tmp15612 {
tmp15609 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15609

tmp15610 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4t_d_1integrity), W5635, W5641, V5625, V5626, V5627, V5628, V5629, V5630)
return
}, 0)

__e.TailApply(PrimFunc(symshen_4system_1S_1h), W5634, W5637, V5625, V5627, V5628, V5629, tmp15610)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15613 := PrimTail(W5640)

tmp15614 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15613, V5627)


__e.TailApply(tmp15608, tmp15614)
return


}, 1)

tmp15615 := PrimHead(W5640)

__e.TailApply(tmp15607, tmp15615)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15618 := PrimTail(W5638)

tmp15619 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15618, V5627)


__e.TailApply(tmp15606, tmp15619)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15622 := PrimHead(W5638)

tmp15623 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15622, V5627)


__e.TailApply(tmp15605, tmp15623)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15626 := PrimTail(W5636)

tmp15627 := Call(__e, PrimFunc(symshen_4lazyderef), tmp15626, V5627)


__e.TailApply(tmp15604, tmp15627)
return


}, 1)

tmp15628 := PrimHead(W5636)

__e.TailApply(tmp15603, tmp15628)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15631 := Call(__e, PrimFunc(symshen_4lazyderef), V5624, V5627)


__e.TailApply(tmp15602, tmp15631)
return


}, 1)

tmp15632 := PrimTail(W5633)

__e.TailApply(tmp15601, tmp15632)
return


}, 1)

tmp15633 := PrimHead(W5633)

__e.TailApply(tmp15600, tmp15633)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15636 := Call(__e, PrimFunc(symshen_4lazyderef), V5623, V5627)


__e.TailApply(tmp15599, tmp15636)
return


} else {
__e.Return(False)
return
}


} else {
__e.Return(W5631)
return
}


}, 1)

tmp15648 := Call(__e, PrimFunc(symshen_4unlocked_2), V5628)


var ifres15641 Obj

if True == tmp15648 {
tmp15642 := MakeNative(func(__e *ControlFlow) {
W5632 := __e.Get(1)
_ = W5632
tmp15645 := PrimEqual(W5632, Nil)

if True == tmp15645 {
tmp15643 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp15643

__e.TailApply(PrimFunc(symis_b), V5624, V5626, V5627, V5628, V5629, V5630)
return


} else {
__e.Return(False)
return
}


}, 1)

tmp15646 := Call(__e, PrimFunc(symshen_4lazyderef), V5623, V5627)


tmp15647 := Call(__e, tmp15642, tmp15646)


ifres15641 = tmp15647


} else {
ifres15641 = False


}

__e.TailApply(tmp15598, ifres15641)
return


}, 8)

tmp15649 := Call(__e, ns2_1set, symshen_4t_d_1integrity, tmp15597)


_ = tmp15649

tmp15650 := MakeNative(func(__e *ControlFlow) {
V5643 := __e.Get(1)
_ = V5643
tmp15659 := PrimIsVector(V5643)

if True == tmp15659 {
tmp15656 := PrimIsString(V5643)

tmp15657 := PrimNot(tmp15656)

var ifres15652 Obj

if True == tmp15657 {
tmp15654 := PrimVectorGet(V5643, MakeNumber(0))

tmp15655 := PrimEqual(tmp15654, symshen_4print_1freshterm)

var ifres15653 Obj

if True == tmp15655 {
ifres15653 = True


} else {
ifres15653 = False


}

ifres15652 = ifres15653


} else {
ifres15652 = False


}

if True == ifres15652 {
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

__e.TailApply(ns2_1set, symshen_4freshterm_2, tmp15650)
return




}, 0)

