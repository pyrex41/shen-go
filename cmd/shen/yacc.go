package main

import . "github.com/pyrex41/shen-go/kl"

var YaccMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp17518 := MakeNative2(func(__e *ControlFlow, V112, V113 Obj) Obj {
tmp17519 := Call1(__e, V112, V113)


let__17380 := tmp17519
_ = let__17380

tmp17526 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17380)


if True == tmp17526 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("parse failure\n"))
} else {
tmp17524 := Call1(__e, SymFn(symshen_4partial_1parse_1failure_2), let__17380)


if True == tmp17524 {
tmp17520 := Call1(__e, SymFn(symshen_4in_1_6), let__17380)


tmp17521 := PrimGuard2(__e, symset, PrimSet, symshen_4_dresidue_d, tmp17520)

_ = tmp17521

tmp17522 := GValue(__e, symshen_4_dresidue_d)

return TailCall1(__e, SymFn(symshen_4raise_1syntax_1error), tmp17522)


} else {
return TailCall1(__e, SymFn(symshen_4_5_1out), let__17380)
}


}


panic("unreachable")
})

tmp17527 := Call2(__e, ns2_1set, symcompile, tmp17518)


_ = tmp17527

tmp17528 := MakeNative1(func(__e *ControlFlow, V115 Obj) Obj {
tmp17529 := GValue(__e, sym_dmaximum_1print_1sequence_1size_d)

tmp17530 := Call3(__e, SymFn(symshen_4syntax_1error_1message), tmp17529, MakeInteger(0), V115)


tmp17532 := Call1(__e, SymFn(symshen_4proc_1nl), PrimGuard2(__e, symcn, PrimStringConcat, MakeString("syntax error here: "), tmp17530))


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp17532)


panic("unreachable")
})

tmp17533 := Call2(__e, ns2_1set, symshen_4raise_1syntax_1error, tmp17528)


_ = tmp17533

tmp17534 := MakeNative3(func(__e *ControlFlow, V123, V124, V125 Obj) Obj {
tmp17545 := GEq(__e, Nil, V125)

if True == tmp17545 {
return MakeString("\n")
} else {
tmp17543 := GEq(__e, V123, V124)

if True == tmp17543 {
return MakeString("...etc \n")
} else {
tmp17541 := GConsP(__e, V125)

if True == tmp17541 {
tmp17535 := GHd(__e, V125)

tmp17536 := Call3(__e, SymFn(symshen_4app), tmp17535, MakeString(" "), symshen_4s)


tmp17537 := GAdd(__e, V124, MakeInteger(1))

tmp17538 := GTl(__e, V125)

tmp17539 := Call3(__e, SymFn(symshen_4syntax_1error_1message), V123, tmp17537, tmp17538)


return PrimGuard2(__e, symcn, PrimStringConcat, tmp17536, tmp17539)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4syntax_1error_1message)
}


}


}


panic("unreachable")
})

tmp17546 := Call2(__e, ns2_1set, symshen_4syntax_1error_1message, tmp17534)


_ = tmp17546

tmp17547 := MakeNative1(func(__e *ControlFlow, V126 Obj) Obj {
tmp17548 := Call0(__e, SymFn(symfail))


return GEq(__e, V126, tmp17548)


panic("unreachable")
})

tmp17549 := Call2(__e, ns2_1set, symshen_4parse_1failure_2, tmp17547)


_ = tmp17549

tmp17550 := MakeNative1(func(__e *ControlFlow, V127 Obj) Obj {
tmp17551 := Call1(__e, SymFn(symshen_4in_1_6), V127)


return GConsP(__e, tmp17551)


panic("unreachable")
})

tmp17552 := Call2(__e, ns2_1set, symshen_4partial_1parse_1failure_2, tmp17550)


_ = tmp17552

tmp17553 := MakeNative1(func(__e *ControlFlow, V130 Obj) Obj {
tmp17566 := GConsP(__e, V130)

var ifres17557 Obj

if True == tmp17566 {
tmp17564 := GTl(__e, V130)

tmp17565 := GConsP(__e, tmp17564)

var ifres17559 Obj

if True == tmp17565 {
tmp17561 := GTl(__e, V130)

tmp17562 := GTl(__e, tmp17561)

tmp17563 := GEq(__e, Nil, tmp17562)

var ifres17560 Obj

if True == tmp17563 {
ifres17560 = True


} else {
ifres17560 = False


}

ifres17559 = ifres17560


} else {
ifres17559 = False


}

var ifres17558 Obj

if True == ifres17559 {
ifres17558 = True


} else {
ifres17558 = False


}

ifres17557 = ifres17558


} else {
ifres17557 = False


}

if True == ifres17557 {
tmp17554 := GTl(__e, V130)

return GHd(__e, tmp17554)


} else {
tmp17555 := Call3(__e, SymFn(symshen_4app), V130, MakeString(" is not a YACC stream\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp17555)


}


panic("unreachable")
})

tmp17567 := Call2(__e, ns2_1set, symshen_4objectcode, tmp17553)


_ = tmp17567

tmp17568 := MakeNative1(func(__e *ControlFlow, V131 Obj) Obj {
tmp17569 := MakeNative1(func(__e *ControlFlow, Z132 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5yacc_6), Z132)
panic("unreachable")
})

return TailCall2(__e, SymFn(symcompile), tmp17569, V131)


panic("unreachable")
})

tmp17570 := Call2(__e, ns2_1set, symshen_4yacc_1_6shen, tmp17568)


_ = tmp17570

tmp17571 := MakeNative1(func(__e *ControlFlow, V133 Obj) Obj {
tmp17598 := GConsP(__e, V133)

var ifres17572 Obj

if True == tmp17598 {
tmp17573 := Call1(__e, SymFn(symhead), V133)


let__17382 := tmp17573
_ = let__17382

tmp17574 := Call1(__e, SymFn(symtail), V133)


let__17383 := tmp17574
_ = let__17383

tmp17575 := Call1(__e, SymFn(symshen_4_5yaccsig_6), let__17383)


let__17384 := tmp17575
_ = let__17384

tmp17596 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17384)


var ifres17576 Obj

if True == tmp17596 {
tmp17577 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17576 = tmp17577


} else {
tmp17578 := Call1(__e, SymFn(symshen_4_5_1out), let__17384)


let__17385 := tmp17578
_ = let__17385

tmp17579 := Call1(__e, SymFn(symshen_4in_1_6), let__17384)


let__17386 := tmp17579
_ = let__17386

tmp17580 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__17386)


let__17387 := tmp17580
_ = let__17387

tmp17595 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17387)


var ifres17581 Obj

if True == tmp17595 {
tmp17582 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17581 = tmp17582


} else {
tmp17583 := Call1(__e, SymFn(symshen_4_5_1out), let__17387)


let__17388 := tmp17583
_ = let__17388

tmp17584 := Call1(__e, SymFn(symshen_4in_1_6), let__17387)


let__17389 := tmp17584
_ = let__17389

tmp17585 := Call1(__e, SymFn(symgensym), symS)


let__17390 := tmp17585
_ = let__17390

tmp17586 := GCons(__e, let__17382, Nil)

tmp17587 := GCons(__e, symdefine, tmp17586)

tmp17588 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), let__17385, let__17390, let__17388)


tmp17589 := GCons(__e, tmp17588, Nil)

tmp17590 := GCons(__e, sym_1_6, tmp17589)

tmp17591 := GCons(__e, let__17390, tmp17590)

tmp17592 := Call2(__e, SymFn(symappend), let__17385, tmp17591)


tmp17593 := Call2(__e, SymFn(symappend), tmp17587, tmp17592)


let__17391 := tmp17593
_ = let__17391

tmp17594 := Call2(__e, SymFn(symshen_4comb), let__17389, let__17391)


ifres17581 = tmp17594


}

ifres17576 = ifres17581


}

ifres17572 = ifres17576


} else {
tmp17597 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17572 = tmp17597


}

let__17381 := ifres17572
_ = let__17381

tmp17600 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17381)


if True == tmp17600 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17381
}


panic("unreachable")
})

tmp17601 := Call2(__e, ns2_1set, symshen_4_5yacc_6, tmp17571)


_ = tmp17601

tmp17602 := MakeNative1(func(__e *ControlFlow, V145 Obj) Obj {
tmp17659 := GConsP(__e, V145)

var ifres17603 Obj

if True == tmp17659 {
tmp17604 := Call1(__e, SymFn(symhead), V145)


let__17393 := tmp17604
_ = let__17393

tmp17605 := Call1(__e, SymFn(symtail), V145)


let__17394 := tmp17605
_ = let__17394

tmp17657 := Call1(__e, SymFn(symshen_4ccons_2), let__17394)


var ifres17606 Obj

if True == tmp17657 {
tmp17607 := Call1(__e, SymFn(symhead), let__17394)


let__17395 := tmp17607
_ = let__17395

tmp17608 := Call1(__e, SymFn(symtail), let__17394)


let__17396 := tmp17608
_ = let__17396

tmp17655 := Call2(__e, SymFn(symshen_4hds_a_2), let__17395, symlist)


var ifres17609 Obj

if True == tmp17655 {
tmp17610 := Call1(__e, SymFn(symtail), let__17395)


let__17397 := tmp17610
_ = let__17397

tmp17653 := GConsP(__e, let__17397)

var ifres17611 Obj

if True == tmp17653 {
tmp17612 := Call1(__e, SymFn(symhead), let__17397)


let__17398 := tmp17612
_ = let__17398

tmp17613 := Call1(__e, SymFn(symtail), let__17397)


let__17399 := tmp17613
_ = let__17399

tmp17614 := Call1(__e, SymFn(sym_5end_6), let__17399)


let__17400 := tmp17614
_ = let__17400

tmp17651 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17400)


var ifres17615 Obj

if True == tmp17651 {
tmp17616 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17615 = tmp17616


} else {
tmp17617 := Call1(__e, SymFn(symshen_4in_1_6), let__17400)


let__17401 := tmp17617
_ = let__17401

tmp17650 := Call2(__e, SymFn(symshen_4hds_a_2), let__17396, sym_a_a_6)


var ifres17618 Obj

if True == tmp17650 {
tmp17619 := Call1(__e, SymFn(symtail), let__17396)


let__17402 := tmp17619
_ = let__17402

tmp17648 := GConsP(__e, let__17402)

var ifres17620 Obj

if True == tmp17648 {
tmp17621 := Call1(__e, SymFn(symhead), let__17402)


let__17403 := tmp17621
_ = let__17403

tmp17622 := Call1(__e, SymFn(symtail), let__17402)


let__17404 := tmp17622
_ = let__17404

tmp17646 := GConsP(__e, let__17404)

var ifres17623 Obj

if True == tmp17646 {
tmp17624 := Call1(__e, SymFn(symhead), let__17404)


let__17405 := tmp17624
_ = let__17405

tmp17625 := Call1(__e, SymFn(symtail), let__17404)


let__17406 := tmp17625
_ = let__17406

tmp17644 := GEq(__e, sym_i, let__17393)

var ifres17641 Obj

if True == tmp17644 {
tmp17643 := GEq(__e, sym_j, let__17405)

var ifres17642 Obj

if True == tmp17643 {
ifres17642 = True


} else {
ifres17642 = False


}

ifres17641 = ifres17642


} else {
ifres17641 = False


}

var ifres17626 Obj

if True == ifres17641 {
tmp17627 := GCons(__e, let__17398, Nil)

tmp17628 := GCons(__e, symlist, tmp17627)

tmp17629 := GCons(__e, let__17398, Nil)

tmp17630 := GCons(__e, symlist, tmp17629)

tmp17631 := GCons(__e, let__17403, Nil)

tmp17632 := GCons(__e, tmp17630, tmp17631)

tmp17633 := GCons(__e, symstr, tmp17632)

tmp17634 := GCons(__e, sym_j, Nil)

tmp17635 := GCons(__e, tmp17633, tmp17634)

tmp17636 := GCons(__e, sym_1_1_6, tmp17635)

tmp17637 := GCons(__e, tmp17628, tmp17636)

tmp17638 := GCons(__e, sym_i, tmp17637)

tmp17639 := Call2(__e, SymFn(symshen_4comb), let__17406, tmp17638)


ifres17626 = tmp17639


} else {
tmp17640 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17626 = tmp17640


}

ifres17623 = ifres17626


} else {
tmp17645 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17623 = tmp17645


}

ifres17620 = ifres17623


} else {
tmp17647 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17620 = tmp17647


}

ifres17618 = ifres17620


} else {
tmp17649 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17618 = tmp17649


}

ifres17615 = ifres17618


}

ifres17611 = ifres17615


} else {
tmp17652 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17611 = tmp17652


}

ifres17609 = ifres17611


} else {
tmp17654 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17609 = tmp17654


}

ifres17606 = ifres17609


} else {
tmp17656 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17606 = tmp17656


}

ifres17603 = ifres17606


} else {
tmp17658 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17603 = tmp17658


}

let__17392 := ifres17603
_ = let__17392

tmp17669 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17392)


if True == tmp17669 {
tmp17660 := Call1(__e, SymFn(sym_5e_6), V145)


let__17408 := tmp17660
_ = let__17408

tmp17665 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17408)


var ifres17661 Obj

if True == tmp17665 {
tmp17662 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17661 = tmp17662


} else {
tmp17663 := Call1(__e, SymFn(symshen_4in_1_6), let__17408)


let__17409 := tmp17663
_ = let__17409

tmp17664 := Call2(__e, SymFn(symshen_4comb), let__17409, Nil)


ifres17661 = tmp17664


}

let__17407 := ifres17661
_ = let__17407

tmp17667 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17407)


if True == tmp17667 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17407
}


} else {
return let__17392
}


panic("unreachable")
})

tmp17670 := Call2(__e, ns2_1set, symshen_4_5yaccsig_6, tmp17602)


_ = tmp17670

tmp17671 := MakeNative1(func(__e *ControlFlow, V164 Obj) Obj {
tmp17672 := Call1(__e, SymFn(symshen_4_5c_1rule_6), V164)


let__17411 := tmp17672
_ = let__17411

tmp17685 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17411)


var ifres17673 Obj

if True == tmp17685 {
tmp17674 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17673 = tmp17674


} else {
tmp17675 := Call1(__e, SymFn(symshen_4_5_1out), let__17411)


let__17412 := tmp17675
_ = let__17412

tmp17676 := Call1(__e, SymFn(symshen_4in_1_6), let__17411)


let__17413 := tmp17676
_ = let__17413

tmp17677 := Call1(__e, SymFn(symshen_4_5c_1rules_6), let__17413)


let__17414 := tmp17677
_ = let__17414

tmp17684 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17414)


var ifres17678 Obj

if True == tmp17684 {
tmp17679 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17678 = tmp17679


} else {
tmp17680 := Call1(__e, SymFn(symshen_4_5_1out), let__17414)


let__17415 := tmp17680
_ = let__17415

tmp17681 := Call1(__e, SymFn(symshen_4in_1_6), let__17414)


let__17416 := tmp17681
_ = let__17416

tmp17682 := GCons(__e, let__17412, let__17415)

tmp17683 := Call2(__e, SymFn(symshen_4comb), let__17416, tmp17682)


ifres17678 = tmp17683


}

ifres17673 = ifres17678


}

let__17410 := ifres17673
_ = let__17410

tmp17701 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17410)


if True == tmp17701 {
tmp17686 := Call1(__e, SymFn(sym_5_b_6), V164)


let__17418 := tmp17686
_ = let__17418

tmp17697 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17418)


var ifres17687 Obj

if True == tmp17697 {
tmp17688 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17687 = tmp17688


} else {
tmp17689 := Call1(__e, SymFn(symshen_4_5_1out), let__17418)


let__17419 := tmp17689
_ = let__17419

tmp17690 := Call1(__e, SymFn(symshen_4in_1_6), let__17418)


let__17420 := tmp17690
_ = let__17420

tmp17695 := Call1(__e, SymFn(symempty_2), let__17419)


var ifres17691 Obj

if True == tmp17695 {
ifres17691 = Nil


} else {
tmp17692 := Call3(__e, SymFn(symshen_4app), let__17419, MakeString("\n ..."), symshen_4r)


tmp17694 := (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symsimple_1error) {
return PrimSimpleError(PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp17692))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("YACC syntax error here:\n "), tmp17692)
return Call1(__e, SymFn(symsimple_1error), __typedArg0)
})()

ifres17691 = tmp17694


}

tmp17696 := Call2(__e, SymFn(symshen_4comb), let__17420, ifres17691)


ifres17687 = tmp17696


}

let__17417 := ifres17687
_ = let__17417

tmp17699 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17417)


if True == tmp17699 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17417
}


} else {
return let__17410
}


panic("unreachable")
})

tmp17702 := Call2(__e, ns2_1set, symshen_4_5c_1rules_6, tmp17671)


_ = tmp17702

tmp17703 := MakeNative1(func(__e *ControlFlow, V176 Obj) Obj {
tmp17704 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__17422 := tmp17704
_ = let__17422

tmp17723 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17422)


var ifres17705 Obj

if True == tmp17723 {
tmp17706 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17705 = tmp17706


} else {
tmp17707 := Call1(__e, SymFn(symshen_4_5_1out), let__17422)


let__17423 := tmp17707
_ = let__17423

tmp17708 := Call1(__e, SymFn(symshen_4in_1_6), let__17422)


let__17424 := tmp17708
_ = let__17424

tmp17709 := Call1(__e, SymFn(symshen_4_5semantics_6), let__17424)


let__17425 := tmp17709
_ = let__17425

tmp17722 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17425)


var ifres17710 Obj

if True == tmp17722 {
tmp17711 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17710 = tmp17711


} else {
tmp17712 := Call1(__e, SymFn(symshen_4_5_1out), let__17425)


let__17426 := tmp17712
_ = let__17426

tmp17713 := Call1(__e, SymFn(symshen_4in_1_6), let__17425)


let__17427 := tmp17713
_ = let__17427

tmp17714 := Call1(__e, SymFn(symshen_4_5sc_6), let__17427)


let__17428 := tmp17714
_ = let__17428

tmp17721 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17428)


var ifres17715 Obj

if True == tmp17721 {
tmp17716 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17715 = tmp17716


} else {
tmp17717 := Call1(__e, SymFn(symshen_4in_1_6), let__17428)


let__17429 := tmp17717
_ = let__17429

tmp17718 := GCons(__e, let__17426, Nil)

tmp17719 := GCons(__e, let__17423, tmp17718)

tmp17720 := Call2(__e, SymFn(symshen_4comb), let__17429, tmp17719)


ifres17715 = tmp17720


}

ifres17710 = ifres17715


}

ifres17705 = ifres17710


}

let__17421 := ifres17705
_ = let__17421

tmp17742 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17421)


if True == tmp17742 {
tmp17724 := Call1(__e, SymFn(symshen_4_5syntax_6), V176)


let__17431 := tmp17724
_ = let__17431

tmp17738 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17431)


var ifres17725 Obj

if True == tmp17738 {
tmp17726 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17725 = tmp17726


} else {
tmp17727 := Call1(__e, SymFn(symshen_4_5_1out), let__17431)


let__17432 := tmp17727
_ = let__17432

tmp17728 := Call1(__e, SymFn(symshen_4in_1_6), let__17431)


let__17433 := tmp17728
_ = let__17433

tmp17729 := Call1(__e, SymFn(symshen_4_5sc_6), let__17433)


let__17434 := tmp17729
_ = let__17434

tmp17737 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17434)


var ifres17730 Obj

if True == tmp17737 {
tmp17731 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17730 = tmp17731


} else {
tmp17732 := Call1(__e, SymFn(symshen_4in_1_6), let__17434)


let__17435 := tmp17732
_ = let__17435

tmp17733 := Call1(__e, SymFn(symshen_4autocomplete), let__17432)


tmp17734 := GCons(__e, tmp17733, Nil)

tmp17735 := GCons(__e, let__17432, tmp17734)

tmp17736 := Call2(__e, SymFn(symshen_4comb), let__17435, tmp17735)


ifres17730 = tmp17736


}

ifres17725 = ifres17730


}

let__17430 := ifres17725
_ = let__17430

tmp17740 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17430)


if True == tmp17740 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17430
}


} else {
return let__17421
}


panic("unreachable")
})

tmp17743 := Call2(__e, ns2_1set, symshen_4_5c_1rule_6, tmp17703)


_ = tmp17743

tmp17744 := MakeNative1(func(__e *ControlFlow, V192 Obj) Obj {
tmp17773 := GConsP(__e, V192)

var ifres17765 Obj

if True == tmp17773 {
tmp17771 := GTl(__e, V192)

tmp17772 := GEq(__e, Nil, tmp17771)

var ifres17767 Obj

if True == tmp17772 {
tmp17769 := GHd(__e, V192)

tmp17770 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17769)


var ifres17768 Obj

if True == tmp17770 {
ifres17768 = True


} else {
ifres17768 = False


}

ifres17767 = ifres17768


} else {
ifres17767 = False


}

var ifres17766 Obj

if True == ifres17767 {
ifres17766 = True


} else {
ifres17766 = False


}

ifres17765 = ifres17766


} else {
ifres17765 = False


}

if True == ifres17765 {
return GHd(__e, V192)
} else {
tmp17763 := GConsP(__e, V192)

var ifres17759 Obj

if True == tmp17763 {
tmp17761 := GHd(__e, V192)

tmp17762 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp17761)


var ifres17760 Obj

if True == tmp17762 {
ifres17760 = True


} else {
ifres17760 = False


}

ifres17759 = ifres17760


} else {
ifres17759 = False


}

if True == ifres17759 {
tmp17745 := GHd(__e, V192)

tmp17746 := GTl(__e, V192)

tmp17747 := Call1(__e, SymFn(symshen_4autocomplete), tmp17746)


tmp17748 := GCons(__e, tmp17747, Nil)

tmp17749 := GCons(__e, tmp17745, tmp17748)

return GCons(__e, symappend, tmp17749)


} else {
tmp17757 := GConsP(__e, V192)

if True == tmp17757 {
tmp17750 := GHd(__e, V192)

tmp17751 := Call1(__e, SymFn(symshen_4autocomplete), tmp17750)


tmp17752 := GTl(__e, V192)

tmp17753 := Call1(__e, SymFn(symshen_4autocomplete), tmp17752)


tmp17754 := GCons(__e, tmp17753, Nil)

tmp17755 := GCons(__e, tmp17751, tmp17754)

return GCons(__e, symcons, tmp17755)


} else {
return V192
}


}


}


panic("unreachable")
})

tmp17774 := Call2(__e, ns2_1set, symshen_4autocomplete, tmp17744)


_ = tmp17774

tmp17775 := MakeNative1(func(__e *ControlFlow, V193 Obj) Obj {
tmp17781 := PrimGuard1(__e, symsymbol_2, PrimIsSymbol, V193)

if True == tmp17781 {
tmp17777 := Call1(__e, SymFn(symexplode), V193)


let__17436 := tmp17777
_ = let__17436

tmp17778 := MakeNative1(func(__e *ControlFlow, Z195 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5non_1terminal_2_6), Z195)
panic("unreachable")
})

tmp17779 := Call2(__e, SymFn(symcompile), tmp17778, let__17436)


if True == tmp17779 {
return True
} else {
return False
}


} else {
return False
}


panic("unreachable")
})

tmp17782 := Call2(__e, ns2_1set, symshen_4non_1terminal_2, tmp17775)


_ = tmp17782

tmp17783 := MakeNative1(func(__e *ControlFlow, V196 Obj) Obj {
tmp17784 := Call1(__e, SymFn(symshen_4_5packagenames_6), V196)


let__17438 := tmp17784
_ = let__17438

tmp17794 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17438)


var ifres17785 Obj

if True == tmp17794 {
tmp17786 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17785 = tmp17786


} else {
tmp17787 := Call1(__e, SymFn(symshen_4in_1_6), let__17438)


let__17439 := tmp17787
_ = let__17439

tmp17788 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), let__17439)


let__17440 := tmp17788
_ = let__17440

tmp17793 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17440)


var ifres17789 Obj

if True == tmp17793 {
tmp17790 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17789 = tmp17790


} else {
tmp17791 := Call1(__e, SymFn(symshen_4in_1_6), let__17440)


let__17441 := tmp17791
_ = let__17441

tmp17792 := Call2(__e, SymFn(symshen_4comb), let__17441, True)


ifres17789 = tmp17792


}

ifres17785 = ifres17789


}

let__17437 := ifres17785
_ = let__17437

tmp17812 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17437)


if True == tmp17812 {
tmp17795 := Call1(__e, SymFn(symshen_4_5non_1terminal_1name_6), V196)


let__17443 := tmp17795
_ = let__17443

tmp17800 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17443)


var ifres17796 Obj

if True == tmp17800 {
tmp17797 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17796 = tmp17797


} else {
tmp17798 := Call1(__e, SymFn(symshen_4in_1_6), let__17443)


let__17444 := tmp17798
_ = let__17444

tmp17799 := Call2(__e, SymFn(symshen_4comb), let__17444, True)


ifres17796 = tmp17799


}

let__17442 := ifres17796
_ = let__17442

tmp17810 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17442)


if True == tmp17810 {
tmp17801 := Call1(__e, SymFn(sym_5_b_6), V196)


let__17446 := tmp17801
_ = let__17446

tmp17806 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17446)


var ifres17802 Obj

if True == tmp17806 {
tmp17803 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17802 = tmp17803


} else {
tmp17804 := Call1(__e, SymFn(symshen_4in_1_6), let__17446)


let__17447 := tmp17804
_ = let__17447

tmp17805 := Call2(__e, SymFn(symshen_4comb), let__17447, False)


ifres17802 = tmp17805


}

let__17445 := ifres17802
_ = let__17445

tmp17808 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17445)


if True == tmp17808 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17445
}


} else {
return let__17442
}


} else {
return let__17437
}


panic("unreachable")
})

tmp17813 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_2_6, tmp17783)


_ = tmp17813

tmp17814 := MakeNative1(func(__e *ControlFlow, V208 Obj) Obj {
tmp17815 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__17449 := tmp17815
_ = let__17449

tmp17829 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17449)


var ifres17816 Obj

if True == tmp17829 {
tmp17817 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17816 = tmp17817


} else {
tmp17818 := Call1(__e, SymFn(symshen_4in_1_6), let__17449)


let__17450 := tmp17818
_ = let__17450

tmp17828 := Call2(__e, SymFn(symshen_4hds_a_2), let__17450, MakeString("."))


var ifres17819 Obj

if True == tmp17828 {
tmp17820 := Call1(__e, SymFn(symtail), let__17450)


let__17451 := tmp17820
_ = let__17451

tmp17821 := Call1(__e, SymFn(symshen_4_5packagenames_6), let__17451)


let__17452 := tmp17821
_ = let__17452

tmp17826 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17452)


var ifres17822 Obj

if True == tmp17826 {
tmp17823 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17822 = tmp17823


} else {
tmp17824 := Call1(__e, SymFn(symshen_4in_1_6), let__17452)


let__17453 := tmp17824
_ = let__17453

tmp17825 := Call2(__e, SymFn(symshen_4comb), let__17453, symshen_4skip)


ifres17822 = tmp17825


}

ifres17819 = ifres17822


} else {
tmp17827 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17819 = tmp17827


}

ifres17816 = ifres17819


}

let__17448 := ifres17816
_ = let__17448

tmp17843 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17448)


if True == tmp17843 {
tmp17830 := Call1(__e, SymFn(symshen_4_5packagename_6), V208)


let__17455 := tmp17830
_ = let__17455

tmp17839 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17455)


var ifres17831 Obj

if True == tmp17839 {
tmp17832 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17831 = tmp17832


} else {
tmp17833 := Call1(__e, SymFn(symshen_4in_1_6), let__17455)


let__17456 := tmp17833
_ = let__17456

tmp17838 := Call2(__e, SymFn(symshen_4hds_a_2), let__17456, MakeString("."))


var ifres17834 Obj

if True == tmp17838 {
tmp17835 := Call1(__e, SymFn(symtail), let__17456)


let__17457 := tmp17835
_ = let__17457

tmp17836 := Call2(__e, SymFn(symshen_4comb), let__17457, symshen_4skip)


ifres17834 = tmp17836


} else {
tmp17837 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17834 = tmp17837


}

ifres17831 = ifres17834


}

let__17454 := ifres17831
_ = let__17454

tmp17841 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17454)


if True == tmp17841 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17454
}


} else {
return let__17448
}


panic("unreachable")
})

tmp17844 := Call2(__e, ns2_1set, symshen_4_5packagenames_6, tmp17814)


_ = tmp17844

tmp17845 := MakeNative1(func(__e *ControlFlow, V219 Obj) Obj {
tmp17846 := Call1(__e, SymFn(symshen_4_5packagechar_6), V219)


let__17459 := tmp17846
_ = let__17459

tmp17856 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17459)


var ifres17847 Obj

if True == tmp17856 {
tmp17848 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17847 = tmp17848


} else {
tmp17849 := Call1(__e, SymFn(symshen_4in_1_6), let__17459)


let__17460 := tmp17849
_ = let__17460

tmp17850 := Call1(__e, SymFn(symshen_4_5packagename_6), let__17460)


let__17461 := tmp17850
_ = let__17461

tmp17855 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17461)


var ifres17851 Obj

if True == tmp17855 {
tmp17852 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17851 = tmp17852


} else {
tmp17853 := Call1(__e, SymFn(symshen_4in_1_6), let__17461)


let__17462 := tmp17853
_ = let__17462

tmp17854 := Call2(__e, SymFn(symshen_4comb), let__17462, symshen_4skip)


ifres17851 = tmp17854


}

ifres17847 = ifres17851


}

let__17458 := ifres17847
_ = let__17458

tmp17866 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17458)


if True == tmp17866 {
tmp17857 := Call1(__e, SymFn(sym_5e_6), V219)


let__17464 := tmp17857
_ = let__17464

tmp17862 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17464)


var ifres17858 Obj

if True == tmp17862 {
tmp17859 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17858 = tmp17859


} else {
tmp17860 := Call1(__e, SymFn(symshen_4in_1_6), let__17464)


let__17465 := tmp17860
_ = let__17465

tmp17861 := Call2(__e, SymFn(symshen_4comb), let__17465, symshen_4skip)


ifres17858 = tmp17861


}

let__17463 := ifres17858
_ = let__17463

tmp17864 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17463)


if True == tmp17864 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17463
}


} else {
return let__17458
}


panic("unreachable")
})

tmp17867 := Call2(__e, ns2_1set, symshen_4_5packagename_6, tmp17845)


_ = tmp17867

tmp17868 := MakeNative1(func(__e *ControlFlow, V228 Obj) Obj {
tmp17878 := GConsP(__e, V228)

var ifres17869 Obj

if True == tmp17878 {
tmp17870 := Call1(__e, SymFn(symhead), V228)


let__17467 := tmp17870
_ = let__17467

tmp17871 := Call1(__e, SymFn(symtail), V228)


let__17468 := tmp17871
_ = let__17468

tmp17875 := GEq(__e, let__17467, MakeString("."))

tmp17876 := GNot(__e, tmp17875)

var ifres17872 Obj

if True == tmp17876 {
tmp17873 := Call2(__e, SymFn(symshen_4comb), let__17468, symshen_4skip)


ifres17872 = tmp17873


} else {
tmp17874 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17872 = tmp17874


}

ifres17869 = ifres17872


} else {
tmp17877 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17869 = tmp17877


}

let__17466 := ifres17869
_ = let__17466

tmp17880 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17466)


if True == tmp17880 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17466
}


panic("unreachable")
})

tmp17881 := Call2(__e, ns2_1set, symshen_4_5packagechar_6, tmp17868)


_ = tmp17881

tmp17882 := MakeNative1(func(__e *ControlFlow, V232 Obj) Obj {
tmp17901 := Call2(__e, SymFn(symshen_4hds_a_2), V232, MakeString("<"))


var ifres17883 Obj

if True == tmp17901 {
tmp17884 := Call1(__e, SymFn(symtail), V232)


let__17470 := tmp17884
_ = let__17470

tmp17885 := Call1(__e, SymFn(sym_5_b_6), let__17470)


let__17471 := tmp17885
_ = let__17471

tmp17899 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17471)


var ifres17886 Obj

if True == tmp17899 {
tmp17887 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17886 = tmp17887


} else {
tmp17888 := Call1(__e, SymFn(symshen_4_5_1out), let__17471)


let__17472 := tmp17888
_ = let__17472

tmp17889 := Call1(__e, SymFn(symshen_4in_1_6), let__17471)


let__17473 := tmp17889
_ = let__17473

tmp17893 := Call1(__e, SymFn(symreverse), let__17472)


let__17474 := tmp17893
_ = let__17474

tmp17898 := GConsP(__e, let__17474)

var ifres17894 Obj

if True == tmp17898 {
tmp17896 := GHd(__e, let__17474)

tmp17897 := GEq(__e, tmp17896, MakeString(">"))

var ifres17895 Obj

if True == tmp17897 {
ifres17895 = True


} else {
ifres17895 = False


}

ifres17894 = ifres17895


} else {
ifres17894 = False


}

var ifres17890 Obj

if True == ifres17894 {
tmp17891 := Call2(__e, SymFn(symshen_4comb), let__17473, symshen_4skip)


ifres17890 = tmp17891


} else {
tmp17892 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17890 = tmp17892


}

ifres17886 = ifres17890


}

ifres17883 = ifres17886


} else {
tmp17900 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17883 = tmp17900


}

let__17469 := ifres17883
_ = let__17469

tmp17903 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17469)


if True == tmp17903 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17469
}


panic("unreachable")
})

tmp17904 := Call2(__e, ns2_1set, symshen_4_5non_1terminal_1name_6, tmp17882)


_ = tmp17904

tmp17905 := MakeNative1(func(__e *ControlFlow, V239 Obj) Obj {
tmp17906 := PrimGuard1(__e, symintern, PrimIntern, MakeString(";"))

return GEq(__e, V239, tmp17906)


panic("unreachable")
})

tmp17907 := Call2(__e, ns2_1set, symshen_4semicolon_2, tmp17905)


_ = tmp17907

tmp17908 := MakeNative1(func(__e *ControlFlow, V240 Obj) Obj {
tmp17917 := GConsP(__e, V240)

var ifres17909 Obj

if True == tmp17917 {
tmp17910 := Call1(__e, SymFn(symhead), V240)


let__17476 := tmp17910
_ = let__17476

tmp17911 := Call1(__e, SymFn(symtail), V240)


let__17477 := tmp17911
_ = let__17477

tmp17915 := Call1(__e, SymFn(symshen_4colon_1equal_2), let__17476)


var ifres17912 Obj

if True == tmp17915 {
tmp17913 := Call2(__e, SymFn(symshen_4comb), let__17477, symshen_4skip)


ifres17912 = tmp17913


} else {
tmp17914 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17912 = tmp17914


}

ifres17909 = ifres17912


} else {
tmp17916 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17909 = tmp17916


}

let__17475 := ifres17909
_ = let__17475

tmp17919 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17475)


if True == tmp17919 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17475
}


panic("unreachable")
})

tmp17920 := Call2(__e, ns2_1set, symshen_4_5colon_1equal_6, tmp17908)


_ = tmp17920

tmp17921 := MakeNative1(func(__e *ControlFlow, V244 Obj) Obj {
tmp17922 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":="))

return GEq(__e, tmp17922, V244)


panic("unreachable")
})

tmp17923 := Call2(__e, ns2_1set, symshen_4colon_1equal_2, tmp17921)


_ = tmp17923

tmp17924 := MakeNative1(func(__e *ControlFlow, V245 Obj) Obj {
tmp17925 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__17479 := tmp17925
_ = let__17479

tmp17938 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17479)


var ifres17926 Obj

if True == tmp17938 {
tmp17927 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17926 = tmp17927


} else {
tmp17928 := Call1(__e, SymFn(symshen_4_5_1out), let__17479)


let__17480 := tmp17928
_ = let__17480

tmp17929 := Call1(__e, SymFn(symshen_4in_1_6), let__17479)


let__17481 := tmp17929
_ = let__17481

tmp17930 := Call1(__e, SymFn(symshen_4_5syntax_6), let__17481)


let__17482 := tmp17930
_ = let__17482

tmp17937 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17482)


var ifres17931 Obj

if True == tmp17937 {
tmp17932 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17931 = tmp17932


} else {
tmp17933 := Call1(__e, SymFn(symshen_4_5_1out), let__17482)


let__17483 := tmp17933
_ = let__17483

tmp17934 := Call1(__e, SymFn(symshen_4in_1_6), let__17482)


let__17484 := tmp17934
_ = let__17484

tmp17935 := GCons(__e, let__17480, let__17483)

tmp17936 := Call2(__e, SymFn(symshen_4comb), let__17484, tmp17935)


ifres17931 = tmp17936


}

ifres17926 = ifres17931


}

let__17478 := ifres17926
_ = let__17478

tmp17950 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17478)


if True == tmp17950 {
tmp17939 := Call1(__e, SymFn(symshen_4_5syntax_1item_6), V245)


let__17486 := tmp17939
_ = let__17486

tmp17946 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17486)


var ifres17940 Obj

if True == tmp17946 {
tmp17941 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17940 = tmp17941


} else {
tmp17942 := Call1(__e, SymFn(symshen_4_5_1out), let__17486)


let__17487 := tmp17942
_ = let__17487

tmp17943 := Call1(__e, SymFn(symshen_4in_1_6), let__17486)


let__17488 := tmp17943
_ = let__17488

tmp17944 := GCons(__e, let__17487, Nil)

tmp17945 := Call2(__e, SymFn(symshen_4comb), let__17488, tmp17944)


ifres17940 = tmp17945


}

let__17485 := ifres17940
_ = let__17485

tmp17948 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17485)


if True == tmp17948 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17485
}


} else {
return let__17478
}


panic("unreachable")
})

tmp17951 := Call2(__e, ns2_1set, symshen_4_5syntax_6, tmp17924)


_ = tmp17951

tmp17952 := MakeNative1(func(__e *ControlFlow, V257 Obj) Obj {
tmp17961 := GConsP(__e, V257)

var ifres17953 Obj

if True == tmp17961 {
tmp17954 := Call1(__e, SymFn(symhead), V257)


let__17490 := tmp17954
_ = let__17490

tmp17955 := Call1(__e, SymFn(symtail), V257)


let__17491 := tmp17955
_ = let__17491

tmp17959 := Call1(__e, SymFn(symshen_4syntax_1item_2), let__17490)


var ifres17956 Obj

if True == tmp17959 {
tmp17957 := Call2(__e, SymFn(symshen_4comb), let__17491, let__17490)


ifres17956 = tmp17957


} else {
tmp17958 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17956 = tmp17958


}

ifres17953 = ifres17956


} else {
tmp17960 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres17953 = tmp17960


}

let__17489 := ifres17953
_ = let__17489

tmp17963 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17489)


if True == tmp17963 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17489
}


panic("unreachable")
})

tmp17964 := Call2(__e, ns2_1set, symshen_4_5syntax_1item_6, tmp17952)


_ = tmp17964

tmp17965 := MakeNative1(func(__e *ControlFlow, V263 Obj) Obj {
tmp18001 := Call1(__e, SymFn(symshen_4colon_1equal_2), V263)


if True == tmp18001 {
return False
} else {
tmp17999 := Call1(__e, SymFn(symshen_4semicolon_2), V263)


if True == tmp17999 {
return False
} else {
tmp17997 := Call1(__e, SymFn(symatom_2), V263)


if True == tmp17997 {
return True
} else {
tmp17995 := GConsP(__e, V263)

var ifres17976 Obj

if True == tmp17995 {
tmp17993 := GHd(__e, V263)

tmp17994 := GEq(__e, symcons, tmp17993)

var ifres17978 Obj

if True == tmp17994 {
tmp17991 := GTl(__e, V263)

tmp17992 := GConsP(__e, tmp17991)

var ifres17980 Obj

if True == tmp17992 {
tmp17988 := GTl(__e, V263)

tmp17989 := GTl(__e, tmp17988)

tmp17990 := GConsP(__e, tmp17989)

var ifres17982 Obj

if True == tmp17990 {
tmp17984 := GTl(__e, V263)

tmp17985 := GTl(__e, tmp17984)

tmp17986 := GTl(__e, tmp17985)

tmp17987 := GEq(__e, Nil, tmp17986)

var ifres17983 Obj

if True == tmp17987 {
ifres17983 = True


} else {
ifres17983 = False


}

ifres17982 = ifres17983


} else {
ifres17982 = False


}

var ifres17981 Obj

if True == ifres17982 {
ifres17981 = True


} else {
ifres17981 = False


}

ifres17980 = ifres17981


} else {
ifres17980 = False


}

var ifres17979 Obj

if True == ifres17980 {
ifres17979 = True


} else {
ifres17979 = False


}

ifres17978 = ifres17979


} else {
ifres17978 = False


}

var ifres17977 Obj

if True == ifres17978 {
ifres17977 = True


} else {
ifres17977 = False


}

ifres17976 = ifres17977


} else {
ifres17976 = False


}

if True == ifres17976 {
tmp17972 := GTl(__e, V263)

tmp17973 := GHd(__e, tmp17972)

tmp17974 := Call1(__e, SymFn(symshen_4syntax_1item_2), tmp17973)


if True == tmp17974 {
tmp17967 := GTl(__e, V263)

tmp17968 := GTl(__e, tmp17967)

tmp17969 := GHd(__e, tmp17968)

tmp17970 := Call1(__e, SymFn(symshen_4syntax_1item_2), tmp17969)


if True == tmp17970 {
return True
} else {
return False
}


} else {
return False
}


} else {
return False
}


}


}


}


panic("unreachable")
})

tmp18002 := Call2(__e, ns2_1set, symshen_4syntax_1item_2, tmp17965)


_ = tmp18002

tmp18003 := MakeNative1(func(__e *ControlFlow, V264 Obj) Obj {
tmp18004 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__17493 := tmp18004
_ = let__17493

tmp18030 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17493)


var ifres18005 Obj

if True == tmp18030 {
tmp18006 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18005 = tmp18006


} else {
tmp18007 := Call1(__e, SymFn(symshen_4in_1_6), let__17493)


let__17494 := tmp18007
_ = let__17494

tmp18029 := GConsP(__e, let__17494)

var ifres18008 Obj

if True == tmp18029 {
tmp18009 := Call1(__e, SymFn(symhead), let__17494)


let__17495 := tmp18009
_ = let__17495

tmp18010 := Call1(__e, SymFn(symtail), let__17494)


let__17496 := tmp18010
_ = let__17496

tmp18027 := Call2(__e, SymFn(symshen_4hds_a_2), let__17496, symwhere)


var ifres18011 Obj

if True == tmp18027 {
tmp18012 := Call1(__e, SymFn(symtail), let__17496)


let__17497 := tmp18012
_ = let__17497

tmp18025 := GConsP(__e, let__17497)

var ifres18013 Obj

if True == tmp18025 {
tmp18014 := Call1(__e, SymFn(symhead), let__17497)


let__17498 := tmp18014
_ = let__17498

tmp18015 := Call1(__e, SymFn(symtail), let__17497)


let__17499 := tmp18015
_ = let__17499

tmp18022 := Call1(__e, SymFn(symshen_4semicolon_2), let__17495)


tmp18023 := GNot(__e, tmp18022)

var ifres18016 Obj

if True == tmp18023 {
tmp18017 := GCons(__e, let__17495, Nil)

tmp18018 := GCons(__e, let__17498, tmp18017)

tmp18019 := GCons(__e, symwhere, tmp18018)

tmp18020 := Call2(__e, SymFn(symshen_4comb), let__17499, tmp18019)


ifres18016 = tmp18020


} else {
tmp18021 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18016 = tmp18021


}

ifres18013 = ifres18016


} else {
tmp18024 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18013 = tmp18024


}

ifres18011 = ifres18013


} else {
tmp18026 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18011 = tmp18026


}

ifres18008 = ifres18011


} else {
tmp18028 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18008 = tmp18028


}

ifres18005 = ifres18008


}

let__17492 := ifres18005
_ = let__17492

tmp18049 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17492)


if True == tmp18049 {
tmp18031 := Call1(__e, SymFn(symshen_4_5colon_1equal_6), V264)


let__17501 := tmp18031
_ = let__17501

tmp18045 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17501)


var ifres18032 Obj

if True == tmp18045 {
tmp18033 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18032 = tmp18033


} else {
tmp18034 := Call1(__e, SymFn(symshen_4in_1_6), let__17501)


let__17502 := tmp18034
_ = let__17502

tmp18044 := GConsP(__e, let__17502)

var ifres18035 Obj

if True == tmp18044 {
tmp18036 := Call1(__e, SymFn(symhead), let__17502)


let__17503 := tmp18036
_ = let__17503

tmp18037 := Call1(__e, SymFn(symtail), let__17502)


let__17504 := tmp18037
_ = let__17504

tmp18041 := Call1(__e, SymFn(symshen_4semicolon_2), let__17503)


tmp18042 := GNot(__e, tmp18041)

var ifres18038 Obj

if True == tmp18042 {
tmp18039 := Call2(__e, SymFn(symshen_4comb), let__17504, let__17503)


ifres18038 = tmp18039


} else {
tmp18040 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18038 = tmp18040


}

ifres18035 = ifres18038


} else {
tmp18043 := Call0(__e, SymFn(symshen_4parse_1failure))


ifres18035 = tmp18043


}

ifres18032 = ifres18035


}

let__17500 := ifres18032
_ = let__17500

tmp18047 := Call1(__e, SymFn(symshen_4parse_1failure_2), let__17500)


if True == tmp18047 {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
} else {
return let__17500
}


} else {
return let__17492
}


panic("unreachable")
})

tmp18050 := Call2(__e, ns2_1set, symshen_4_5semantics_6, tmp18003)


_ = tmp18050

tmp18051 := MakeNative3(func(__e *ControlFlow, V286, V287, V288 Obj) Obj {
tmp18059 := GEq(__e, Nil, V288)

if True == tmp18059 {
return GCons(__e, symshen_4parse_1failure, Nil)
} else {
tmp18057 := GConsP(__e, V288)

if True == tmp18057 {
tmp18052 := GHd(__e, V288)

tmp18053 := Call3(__e, SymFn(symshen_4c_1rule_1_6shen), V286, tmp18052, V287)


tmp18054 := GTl(__e, V288)

tmp18055 := Call3(__e, SymFn(symshen_4c_1rules_1_6shen), V286, V287, tmp18054)


return TailCall2(__e, SymFn(symshen_4combine_1c_1code), tmp18053, tmp18055)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rules->shen\n"))
}


}


panic("unreachable")
})

tmp18060 := Call2(__e, ns2_1set, symshen_4c_1rules_1_6shen, tmp18051)


_ = tmp18060

tmp18061 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall0(__e, SymFn(symfail))
panic("unreachable")
})

tmp18062 := Call2(__e, ns2_1set, symshen_4parse_1failure, tmp18061)


_ = tmp18062

tmp18063 := MakeNative2(func(__e *ControlFlow, V289, V290 Obj) Obj {
tmp18064 := GCons(__e, symResult, Nil)

tmp18065 := GCons(__e, symshen_4parse_1failure_2, tmp18064)

tmp18066 := GCons(__e, symResult, Nil)

tmp18067 := GCons(__e, V290, tmp18066)

tmp18068 := GCons(__e, tmp18065, tmp18067)

tmp18069 := GCons(__e, symif, tmp18068)

tmp18070 := GCons(__e, tmp18069, Nil)

tmp18071 := GCons(__e, V289, tmp18070)

tmp18072 := GCons(__e, symResult, tmp18071)

return GCons(__e, symlet, tmp18072)


panic("unreachable")
})

tmp18073 := Call2(__e, ns2_1set, symshen_4combine_1c_1code, tmp18063)


_ = tmp18073

tmp18074 := MakeNative3(func(__e *ControlFlow, V297, V298, V299 Obj) Obj {
tmp18088 := GConsP(__e, V298)

var ifres18079 Obj

if True == tmp18088 {
tmp18086 := GTl(__e, V298)

tmp18087 := GConsP(__e, tmp18086)

var ifres18081 Obj

if True == tmp18087 {
tmp18083 := GTl(__e, V298)

tmp18084 := GTl(__e, tmp18083)

tmp18085 := GEq(__e, Nil, tmp18084)

var ifres18082 Obj

if True == tmp18085 {
ifres18082 = True


} else {
ifres18082 = False


}

ifres18081 = ifres18082


} else {
ifres18081 = False


}

var ifres18080 Obj

if True == ifres18081 {
ifres18080 = True


} else {
ifres18080 = False


}

ifres18079 = ifres18080


} else {
ifres18079 = False


}

if True == ifres18079 {
tmp18075 := GHd(__e, V298)

tmp18076 := GTl(__e, V298)

tmp18077 := GHd(__e, tmp18076)

return TailCall4(__e, SymFn(symshen_4yacc_1syntax), V297, V299, tmp18075, tmp18077)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.c-rule->shen\n"))
}


panic("unreachable")
})

tmp18089 := Call2(__e, ns2_1set, symshen_4c_1rule_1_6shen, tmp18074)


_ = tmp18089

tmp18090 := MakeNative4(func(__e *ControlFlow, V308, V309, V310, V311 Obj) Obj {
tmp18154 := GEq(__e, Nil, V310)

var ifres18132 Obj

if True == tmp18154 {
tmp18153 := GConsP(__e, V311)

var ifres18134 Obj

if True == tmp18153 {
tmp18151 := GHd(__e, V311)

tmp18152 := GEq(__e, symwhere, tmp18151)

var ifres18136 Obj

if True == tmp18152 {
tmp18149 := GTl(__e, V311)

tmp18150 := GConsP(__e, tmp18149)

var ifres18138 Obj

if True == tmp18150 {
tmp18146 := GTl(__e, V311)

tmp18147 := GTl(__e, tmp18146)

tmp18148 := GConsP(__e, tmp18147)

var ifres18140 Obj

if True == tmp18148 {
tmp18142 := GTl(__e, V311)

tmp18143 := GTl(__e, tmp18142)

tmp18144 := GTl(__e, tmp18143)

tmp18145 := GEq(__e, Nil, tmp18144)

var ifres18141 Obj

if True == tmp18145 {
ifres18141 = True


} else {
ifres18141 = False


}

ifres18140 = ifres18141


} else {
ifres18140 = False


}

var ifres18139 Obj

if True == ifres18140 {
ifres18139 = True


} else {
ifres18139 = False


}

ifres18138 = ifres18139


} else {
ifres18138 = False


}

var ifres18137 Obj

if True == ifres18138 {
ifres18137 = True


} else {
ifres18137 = False


}

ifres18136 = ifres18137


} else {
ifres18136 = False


}

var ifres18135 Obj

if True == ifres18136 {
ifres18135 = True


} else {
ifres18135 = False


}

ifres18134 = ifres18135


} else {
ifres18134 = False


}

var ifres18133 Obj

if True == ifres18134 {
ifres18133 = True


} else {
ifres18133 = False


}

ifres18132 = ifres18133


} else {
ifres18132 = False


}

if True == ifres18132 {
tmp18091 := GTl(__e, V311)

tmp18092 := GHd(__e, tmp18091)

tmp18093 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), tmp18092)


tmp18094 := GTl(__e, V311)

tmp18095 := GTl(__e, tmp18094)

tmp18096 := GHd(__e, tmp18095)

tmp18097 := Call4(__e, SymFn(symshen_4yacc_1syntax), V308, V309, Nil, tmp18096)


tmp18098 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18099 := GCons(__e, tmp18098, Nil)

tmp18100 := GCons(__e, tmp18097, tmp18099)

tmp18101 := GCons(__e, tmp18093, tmp18100)

return GCons(__e, symif, tmp18101)


} else {
tmp18130 := GEq(__e, Nil, V310)

if True == tmp18130 {
return TailCall3(__e, SymFn(symshen_4yacc_1semantics), V308, V309, V311)
} else {
tmp18128 := GConsP(__e, V310)

if True == tmp18128 {
tmp18125 := GHd(__e, V310)

tmp18126 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp18125)


if True == tmp18126 {
tmp18102 := GHd(__e, V310)

tmp18103 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4non_1terminalcode), V308, V309, tmp18102, tmp18103, V311)


} else {
tmp18122 := GHd(__e, V310)

tmp18123 := PrimGuard1(__e, symvariable_2, PrimIsVariable, tmp18122)

if True == tmp18123 {
tmp18104 := GHd(__e, V310)

tmp18105 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4variablecode), V308, V309, tmp18104, tmp18105, V311)


} else {
tmp18119 := GHd(__e, V310)

tmp18120 := GEq(__e, sym__, tmp18119)

if True == tmp18120 {
tmp18106 := GHd(__e, V310)

tmp18107 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4wildcardcode), V308, V309, tmp18106, tmp18107, V311)


} else {
tmp18116 := GHd(__e, V310)

tmp18117 := Call1(__e, SymFn(symatom_2), tmp18116)


if True == tmp18117 {
tmp18108 := GHd(__e, V310)

tmp18109 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4terminalcode), V308, V309, tmp18108, tmp18109, V311)


} else {
tmp18113 := GHd(__e, V310)

tmp18114 := GConsP(__e, tmp18113)

if True == tmp18114 {
tmp18110 := GHd(__e, V310)

tmp18111 := GTl(__e, V310)

return TailCall5(__e, SymFn(symshen_4conscode), V308, V309, tmp18110, tmp18111, V311)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.yacc-syntax\n"))
}


}


}


}


}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.yacc-syntax\n"))
}


}


}


panic("unreachable")
})

tmp18155 := Call2(__e, ns2_1set, symshen_4yacc_1syntax, tmp18090)


_ = tmp18155

tmp18156 := MakeNative5(func(__e *ControlFlow, V312, V313, V314, V315, V316 Obj) Obj {
tmp18157 := Call2(__e, SymFn(symconcat), symParse, V314)


let__17505 := tmp18157
_ = let__17505

tmp18158 := Call2(__e, SymFn(symconcat), symAction, V314)


let__17506 := tmp18158
_ = let__17506

tmp18159 := Call2(__e, SymFn(symconcat), symRemainder, V314)


let__17507 := tmp18159
_ = let__17507

tmp18160 := GCons(__e, V313, Nil)

tmp18161 := GCons(__e, V314, tmp18160)

tmp18162 := GCons(__e, let__17505, Nil)

tmp18163 := GCons(__e, symshen_4parse_1failure_2, tmp18162)

tmp18164 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18165 := GCons(__e, let__17505, Nil)

tmp18166 := GCons(__e, symshen_4in_1_6, tmp18165)

tmp18167 := Call4(__e, SymFn(symshen_4yacc_1syntax), V312, let__17507, V315, V316)


tmp18168 := GCons(__e, tmp18167, Nil)

tmp18169 := GCons(__e, tmp18166, tmp18168)

tmp18170 := GCons(__e, let__17507, tmp18169)

tmp18171 := GCons(__e, symlet, tmp18170)

let__17508 := tmp18171
_ = let__17508

tmp18182 := Call2(__e, SymFn(symshen_4occurs_1check_2), V314, V316)


var ifres18179 Obj

if True == tmp18182 {
ifres18179 = True


} else {
tmp18181 := Call2(__e, SymFn(symshen_4occurs_1check_2), let__17506, V316)


var ifres18180 Obj

if True == tmp18181 {
ifres18180 = True


} else {
ifres18180 = False


}

ifres18179 = ifres18180


}

var ifres18172 Obj

if True == ifres18179 {
tmp18173 := GCons(__e, let__17505, Nil)

tmp18174 := GCons(__e, symshen_4_5_1out, tmp18173)

tmp18175 := GCons(__e, let__17508, Nil)

tmp18176 := GCons(__e, tmp18174, tmp18175)

tmp18177 := GCons(__e, let__17506, tmp18176)

tmp18178 := GCons(__e, symlet, tmp18177)

ifres18172 = tmp18178


} else {
ifres18172 = let__17508


}

tmp18183 := GCons(__e, ifres18172, Nil)

tmp18184 := GCons(__e, tmp18164, tmp18183)

tmp18185 := GCons(__e, tmp18163, tmp18184)

tmp18186 := GCons(__e, symif, tmp18185)

tmp18187 := GCons(__e, tmp18186, Nil)

tmp18188 := GCons(__e, tmp18161, tmp18187)

tmp18189 := GCons(__e, let__17505, tmp18188)

return GCons(__e, symlet, tmp18189)


panic("unreachable")
})

tmp18190 := Call2(__e, ns2_1set, symshen_4non_1terminalcode, tmp18156)


_ = tmp18190

tmp18191 := MakeNative5(func(__e *ControlFlow, V321, V322, V323, V324, V325 Obj) Obj {
tmp18192 := Call1(__e, SymFn(symgensym), symRemainder)


let__17509 := tmp18192
_ = let__17509

tmp18193 := GCons(__e, V322, Nil)

tmp18194 := GCons(__e, symcons_2, tmp18193)

tmp18195 := GCons(__e, V322, Nil)

tmp18196 := GCons(__e, symtail, tmp18195)

tmp18197 := Call4(__e, SymFn(symshen_4yacc_1syntax), V321, let__17509, V324, V325)


tmp18198 := GCons(__e, tmp18197, Nil)

tmp18199 := GCons(__e, tmp18196, tmp18198)

tmp18200 := GCons(__e, let__17509, tmp18199)

tmp18201 := GCons(__e, symlet, tmp18200)

let__17510 := tmp18201
_ = let__17510

tmp18209 := Call2(__e, SymFn(symshen_4occurs_1check_2), V323, V325)


var ifres18202 Obj

if True == tmp18209 {
tmp18203 := GCons(__e, V322, Nil)

tmp18204 := GCons(__e, symhead, tmp18203)

tmp18205 := GCons(__e, let__17510, Nil)

tmp18206 := GCons(__e, tmp18204, tmp18205)

tmp18207 := GCons(__e, V323, tmp18206)

tmp18208 := GCons(__e, symlet, tmp18207)

ifres18202 = tmp18208


} else {
ifres18202 = let__17510


}

tmp18210 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18211 := GCons(__e, tmp18210, Nil)

tmp18212 := GCons(__e, ifres18202, tmp18211)

tmp18213 := GCons(__e, tmp18194, tmp18212)

return GCons(__e, symif, tmp18213)


panic("unreachable")
})

tmp18214 := Call2(__e, ns2_1set, symshen_4variablecode, tmp18191)


_ = tmp18214

tmp18215 := MakeNative5(func(__e *ControlFlow, V328, V329, V330, V331, V332 Obj) Obj {
tmp18216 := Call1(__e, SymFn(symgensym), symRemainder)


let__17511 := tmp18216
_ = let__17511

tmp18217 := GCons(__e, V329, Nil)

tmp18218 := GCons(__e, symcons_2, tmp18217)

tmp18219 := GCons(__e, V329, Nil)

tmp18220 := GCons(__e, symtail, tmp18219)

tmp18221 := Call4(__e, SymFn(symshen_4yacc_1syntax), V328, let__17511, V331, V332)


tmp18222 := GCons(__e, tmp18221, Nil)

tmp18223 := GCons(__e, tmp18220, tmp18222)

tmp18224 := GCons(__e, let__17511, tmp18223)

tmp18225 := GCons(__e, symlet, tmp18224)

tmp18226 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18227 := GCons(__e, tmp18226, Nil)

tmp18228 := GCons(__e, tmp18225, tmp18227)

tmp18229 := GCons(__e, tmp18218, tmp18228)

return GCons(__e, symif, tmp18229)


panic("unreachable")
})

tmp18230 := Call2(__e, ns2_1set, symshen_4wildcardcode, tmp18215)


_ = tmp18230

tmp18231 := MakeNative5(func(__e *ControlFlow, V334, V335, V336, V337, V338 Obj) Obj {
tmp18232 := Call1(__e, SymFn(symgensym), symRemainder)


let__17512 := tmp18232
_ = let__17512

tmp18233 := GCons(__e, V336, Nil)

tmp18234 := GCons(__e, V335, tmp18233)

tmp18235 := GCons(__e, symshen_4hds_a_2, tmp18234)

tmp18236 := GCons(__e, V335, Nil)

tmp18237 := GCons(__e, symtail, tmp18236)

tmp18238 := Call4(__e, SymFn(symshen_4yacc_1syntax), V334, let__17512, V337, V338)


tmp18239 := GCons(__e, tmp18238, Nil)

tmp18240 := GCons(__e, tmp18237, tmp18239)

tmp18241 := GCons(__e, let__17512, tmp18240)

tmp18242 := GCons(__e, symlet, tmp18241)

tmp18243 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18244 := GCons(__e, tmp18243, Nil)

tmp18245 := GCons(__e, tmp18242, tmp18244)

tmp18246 := GCons(__e, tmp18235, tmp18245)

return GCons(__e, symif, tmp18246)


panic("unreachable")
})

tmp18247 := Call2(__e, ns2_1set, symshen_4terminalcode, tmp18231)


_ = tmp18247

tmp18248 := MakeNative2(func(__e *ControlFlow, V347, V348 Obj) Obj {
tmp18254 := GConsP(__e, V347)

var ifres18250 Obj

if True == tmp18254 {
tmp18252 := GHd(__e, V347)

tmp18253 := GEq(__e, tmp18252, V348)

var ifres18251 Obj

if True == tmp18253 {
ifres18251 = True


} else {
ifres18251 = False


}

ifres18250 = ifres18251


} else {
ifres18250 = False


}

if True == ifres18250 {
return True
} else {
return False
}


panic("unreachable")
})

tmp18255 := Call2(__e, ns2_1set, symshen_4hds_a_2, tmp18248)


_ = tmp18255

tmp18256 := MakeNative5(func(__e *ControlFlow, V349, V350, V351, V352, V353 Obj) Obj {
tmp18257 := Call1(__e, SymFn(symgensym), symRemainder)


let__17513 := tmp18257
_ = let__17513

tmp18258 := Call1(__e, SymFn(symgensym), symHd)


let__17514 := tmp18258
_ = let__17514

tmp18259 := Call1(__e, SymFn(symgensym), symTl)


let__17515 := tmp18259
_ = let__17515

tmp18260 := GCons(__e, V350, Nil)

tmp18261 := GCons(__e, symshen_4ccons_2, tmp18260)

tmp18262 := GCons(__e, V350, Nil)

tmp18263 := GCons(__e, symhead, tmp18262)

tmp18264 := GCons(__e, V350, Nil)

tmp18265 := GCons(__e, symtail, tmp18264)

tmp18266 := Call1(__e, SymFn(symshen_4decons), V351)


tmp18267 := GCons(__e, sym_5end_6, Nil)

tmp18268 := Call2(__e, SymFn(symappend), tmp18266, tmp18267)


tmp18269 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__17515, V352, V353)


tmp18270 := GCons(__e, tmp18269, Nil)

tmp18271 := GCons(__e, symshen_4processed, tmp18270)

tmp18272 := Call4(__e, SymFn(symshen_4yacc_1syntax), V349, let__17514, tmp18268, tmp18271)


tmp18273 := GCons(__e, tmp18272, Nil)

tmp18274 := GCons(__e, tmp18265, tmp18273)

tmp18275 := GCons(__e, let__17515, tmp18274)

tmp18276 := GCons(__e, tmp18263, tmp18275)

tmp18277 := GCons(__e, let__17514, tmp18276)

tmp18278 := GCons(__e, symlet, tmp18277)

tmp18279 := GCons(__e, symshen_4parse_1failure, Nil)

tmp18280 := GCons(__e, tmp18279, Nil)

tmp18281 := GCons(__e, tmp18278, tmp18280)

tmp18282 := GCons(__e, tmp18261, tmp18281)

return GCons(__e, symif, tmp18282)


panic("unreachable")
})

tmp18283 := Call2(__e, ns2_1set, symshen_4conscode, tmp18256)


_ = tmp18283

tmp18284 := MakeNative1(func(__e *ControlFlow, V367 Obj) Obj {
tmp18296 := GConsP(__e, V367)

var ifres18292 Obj

if True == tmp18296 {
tmp18294 := GHd(__e, V367)

tmp18295 := GConsP(__e, tmp18294)

var ifres18293 Obj

if True == tmp18295 {
ifres18293 = True


} else {
ifres18293 = False


}

ifres18292 = ifres18293


} else {
ifres18292 = False


}

if True == ifres18292 {
return True
} else {
tmp18290 := GConsP(__e, V367)

var ifres18286 Obj

if True == tmp18290 {
tmp18288 := GHd(__e, V367)

tmp18289 := GEq(__e, Nil, tmp18288)

var ifres18287 Obj

if True == tmp18289 {
ifres18287 = True


} else {
ifres18287 = False


}

ifres18286 = ifres18287


} else {
ifres18286 = False


}

if True == ifres18286 {
return True
} else {
return False
}


}


panic("unreachable")
})

tmp18297 := Call2(__e, ns2_1set, symshen_4ccons_2, tmp18284)


_ = tmp18297

tmp18298 := MakeNative1(func(__e *ControlFlow, V368 Obj) Obj {
tmp18325 := GConsP(__e, V368)

var ifres18306 Obj

if True == tmp18325 {
tmp18323 := GHd(__e, V368)

tmp18324 := GEq(__e, symcons, tmp18323)

var ifres18308 Obj

if True == tmp18324 {
tmp18321 := GTl(__e, V368)

tmp18322 := GConsP(__e, tmp18321)

var ifres18310 Obj

if True == tmp18322 {
tmp18318 := GTl(__e, V368)

tmp18319 := GTl(__e, tmp18318)

tmp18320 := GConsP(__e, tmp18319)

var ifres18312 Obj

if True == tmp18320 {
tmp18314 := GTl(__e, V368)

tmp18315 := GTl(__e, tmp18314)

tmp18316 := GTl(__e, tmp18315)

tmp18317 := GEq(__e, Nil, tmp18316)

var ifres18313 Obj

if True == tmp18317 {
ifres18313 = True


} else {
ifres18313 = False


}

ifres18312 = ifres18313


} else {
ifres18312 = False


}

var ifres18311 Obj

if True == ifres18312 {
ifres18311 = True


} else {
ifres18311 = False


}

ifres18310 = ifres18311


} else {
ifres18310 = False


}

var ifres18309 Obj

if True == ifres18310 {
ifres18309 = True


} else {
ifres18309 = False


}

ifres18308 = ifres18309


} else {
ifres18308 = False


}

var ifres18307 Obj

if True == ifres18308 {
ifres18307 = True


} else {
ifres18307 = False


}

ifres18306 = ifres18307


} else {
ifres18306 = False


}

if True == ifres18306 {
tmp18299 := GTl(__e, V368)

tmp18300 := GHd(__e, tmp18299)

tmp18301 := GTl(__e, V368)

tmp18302 := GTl(__e, tmp18301)

tmp18303 := GHd(__e, tmp18302)

tmp18304 := Call1(__e, SymFn(symshen_4decons), tmp18303)


return GCons(__e, tmp18300, tmp18304)


} else {
return V368
}


panic("unreachable")
})

tmp18326 := Call2(__e, ns2_1set, symshen_4decons, tmp18298)


_ = tmp18326

tmp18327 := MakeNative2(func(__e *ControlFlow, V369, V370 Obj) Obj {
tmp18328 := GCons(__e, V370, Nil)

return GCons(__e, V369, tmp18328)


panic("unreachable")
})

tmp18329 := Call2(__e, ns2_1set, symshen_4comb, tmp18327)


_ = tmp18329

tmp18330 := MakeNative3(func(__e *ControlFlow, V375, V376, V377 Obj) Obj {
tmp18350 := GConsP(__e, V377)

var ifres18337 Obj

if True == tmp18350 {
tmp18348 := GHd(__e, V377)

tmp18349 := GEq(__e, symshen_4processed, tmp18348)

var ifres18339 Obj

if True == tmp18349 {
tmp18346 := GTl(__e, V377)

tmp18347 := GConsP(__e, tmp18346)

var ifres18341 Obj

if True == tmp18347 {
tmp18343 := GTl(__e, V377)

tmp18344 := GTl(__e, tmp18343)

tmp18345 := GEq(__e, Nil, tmp18344)

var ifres18342 Obj

if True == tmp18345 {
ifres18342 = True


} else {
ifres18342 = False


}

ifres18341 = ifres18342


} else {
ifres18341 = False


}

var ifres18340 Obj

if True == ifres18341 {
ifres18340 = True


} else {
ifres18340 = False


}

ifres18339 = ifres18340


} else {
ifres18339 = False


}

var ifres18338 Obj

if True == ifres18339 {
ifres18338 = True


} else {
ifres18338 = False


}

ifres18337 = ifres18338


} else {
ifres18337 = False


}

if True == ifres18337 {
tmp18331 := GTl(__e, V377)

return GHd(__e, tmp18331)


} else {
tmp18332 := Call1(__e, SymFn(symshen_4process_1yacc_1semantics), V377)


let__17516 := tmp18332
_ = let__17516

tmp18333 := Call2(__e, SymFn(symshen_4use_1type_1info), V375, let__17516)


let__17517 := tmp18333
_ = let__17517

tmp18334 := GCons(__e, let__17517, Nil)

tmp18335 := GCons(__e, V376, tmp18334)

return GCons(__e, symshen_4comb, tmp18335)


}


panic("unreachable")
})

tmp18351 := Call2(__e, ns2_1set, symshen_4yacc_1semantics, tmp18330)


_ = tmp18351

tmp18352 := MakeNative2(func(__e *ControlFlow, V383, V384 Obj) Obj {
tmp18540 := GConsP(__e, V383)

var ifres18361 Obj

if True == tmp18540 {
tmp18538 := GHd(__e, V383)

tmp18539 := GEq(__e, sym_i, tmp18538)

var ifres18363 Obj

if True == tmp18539 {
tmp18536 := GTl(__e, V383)

tmp18537 := GConsP(__e, tmp18536)

var ifres18365 Obj

if True == tmp18537 {
tmp18533 := GTl(__e, V383)

tmp18534 := GHd(__e, tmp18533)

tmp18535 := GConsP(__e, tmp18534)

var ifres18367 Obj

if True == tmp18535 {
tmp18529 := GTl(__e, V383)

tmp18530 := GHd(__e, tmp18529)

tmp18531 := GHd(__e, tmp18530)

tmp18532 := GEq(__e, symlist, tmp18531)

var ifres18369 Obj

if True == tmp18532 {
tmp18525 := GTl(__e, V383)

tmp18526 := GHd(__e, tmp18525)

tmp18527 := GTl(__e, tmp18526)

tmp18528 := GConsP(__e, tmp18527)

var ifres18371 Obj

if True == tmp18528 {
tmp18520 := GTl(__e, V383)

tmp18521 := GHd(__e, tmp18520)

tmp18522 := GTl(__e, tmp18521)

tmp18523 := GTl(__e, tmp18522)

tmp18524 := GEq(__e, Nil, tmp18523)

var ifres18373 Obj

if True == tmp18524 {
tmp18517 := GTl(__e, V383)

tmp18518 := GTl(__e, tmp18517)

tmp18519 := GConsP(__e, tmp18518)

var ifres18375 Obj

if True == tmp18519 {
tmp18513 := GTl(__e, V383)

tmp18514 := GTl(__e, tmp18513)

tmp18515 := GHd(__e, tmp18514)

tmp18516 := GEq(__e, sym_1_1_6, tmp18515)

var ifres18377 Obj

if True == tmp18516 {
tmp18509 := GTl(__e, V383)

tmp18510 := GTl(__e, tmp18509)

tmp18511 := GTl(__e, tmp18510)

tmp18512 := GConsP(__e, tmp18511)

var ifres18379 Obj

if True == tmp18512 {
tmp18504 := GTl(__e, V383)

tmp18505 := GTl(__e, tmp18504)

tmp18506 := GTl(__e, tmp18505)

tmp18507 := GHd(__e, tmp18506)

tmp18508 := GConsP(__e, tmp18507)

var ifres18381 Obj

if True == tmp18508 {
tmp18498 := GTl(__e, V383)

tmp18499 := GTl(__e, tmp18498)

tmp18500 := GTl(__e, tmp18499)

tmp18501 := GHd(__e, tmp18500)

tmp18502 := GHd(__e, tmp18501)

tmp18503 := GEq(__e, symstr, tmp18502)

var ifres18383 Obj

if True == tmp18503 {
tmp18492 := GTl(__e, V383)

tmp18493 := GTl(__e, tmp18492)

tmp18494 := GTl(__e, tmp18493)

tmp18495 := GHd(__e, tmp18494)

tmp18496 := GTl(__e, tmp18495)

tmp18497 := GConsP(__e, tmp18496)

var ifres18385 Obj

if True == tmp18497 {
tmp18485 := GTl(__e, V383)

tmp18486 := GTl(__e, tmp18485)

tmp18487 := GTl(__e, tmp18486)

tmp18488 := GHd(__e, tmp18487)

tmp18489 := GTl(__e, tmp18488)

tmp18490 := GHd(__e, tmp18489)

tmp18491 := GConsP(__e, tmp18490)

var ifres18387 Obj

if True == tmp18491 {
tmp18477 := GTl(__e, V383)

tmp18478 := GTl(__e, tmp18477)

tmp18479 := GTl(__e, tmp18478)

tmp18480 := GHd(__e, tmp18479)

tmp18481 := GTl(__e, tmp18480)

tmp18482 := GHd(__e, tmp18481)

tmp18483 := GHd(__e, tmp18482)

tmp18484 := GEq(__e, symlist, tmp18483)

var ifres18389 Obj

if True == tmp18484 {
tmp18469 := GTl(__e, V383)

tmp18470 := GTl(__e, tmp18469)

tmp18471 := GTl(__e, tmp18470)

tmp18472 := GHd(__e, tmp18471)

tmp18473 := GTl(__e, tmp18472)

tmp18474 := GHd(__e, tmp18473)

tmp18475 := GTl(__e, tmp18474)

tmp18476 := GConsP(__e, tmp18475)

var ifres18391 Obj

if True == tmp18476 {
tmp18460 := GTl(__e, V383)

tmp18461 := GTl(__e, tmp18460)

tmp18462 := GTl(__e, tmp18461)

tmp18463 := GHd(__e, tmp18462)

tmp18464 := GTl(__e, tmp18463)

tmp18465 := GHd(__e, tmp18464)

tmp18466 := GTl(__e, tmp18465)

tmp18467 := GTl(__e, tmp18466)

tmp18468 := GEq(__e, Nil, tmp18467)

var ifres18393 Obj

if True == tmp18468 {
tmp18453 := GTl(__e, V383)

tmp18454 := GTl(__e, tmp18453)

tmp18455 := GTl(__e, tmp18454)

tmp18456 := GHd(__e, tmp18455)

tmp18457 := GTl(__e, tmp18456)

tmp18458 := GTl(__e, tmp18457)

tmp18459 := GConsP(__e, tmp18458)

var ifres18395 Obj

if True == tmp18459 {
tmp18445 := GTl(__e, V383)

tmp18446 := GTl(__e, tmp18445)

tmp18447 := GTl(__e, tmp18446)

tmp18448 := GHd(__e, tmp18447)

tmp18449 := GTl(__e, tmp18448)

tmp18450 := GTl(__e, tmp18449)

tmp18451 := GTl(__e, tmp18450)

tmp18452 := GEq(__e, Nil, tmp18451)

var ifres18397 Obj

if True == tmp18452 {
tmp18440 := GTl(__e, V383)

tmp18441 := GTl(__e, tmp18440)

tmp18442 := GTl(__e, tmp18441)

tmp18443 := GTl(__e, tmp18442)

tmp18444 := GConsP(__e, tmp18443)

var ifres18399 Obj

if True == tmp18444 {
tmp18434 := GTl(__e, V383)

tmp18435 := GTl(__e, tmp18434)

tmp18436 := GTl(__e, tmp18435)

tmp18437 := GTl(__e, tmp18436)

tmp18438 := GHd(__e, tmp18437)

tmp18439 := GEq(__e, sym_j, tmp18438)

var ifres18401 Obj

if True == tmp18439 {
tmp18428 := GTl(__e, V383)

tmp18429 := GTl(__e, tmp18428)

tmp18430 := GTl(__e, tmp18429)

tmp18431 := GTl(__e, tmp18430)

tmp18432 := GTl(__e, tmp18431)

tmp18433 := GEq(__e, Nil, tmp18432)

var ifres18403 Obj

if True == tmp18433 {
tmp18415 := GTl(__e, V383)

tmp18416 := GHd(__e, tmp18415)

tmp18417 := GTl(__e, tmp18416)

tmp18418 := GHd(__e, tmp18417)

tmp18419 := GTl(__e, V383)

tmp18420 := GTl(__e, tmp18419)

tmp18421 := GTl(__e, tmp18420)

tmp18422 := GHd(__e, tmp18421)

tmp18423 := GTl(__e, tmp18422)

tmp18424 := GHd(__e, tmp18423)

tmp18425 := GTl(__e, tmp18424)

tmp18426 := GHd(__e, tmp18425)

tmp18427 := GEq(__e, tmp18418, tmp18426)

var ifres18405 Obj

if True == tmp18427 {
tmp18407 := GTl(__e, V383)

tmp18408 := GTl(__e, tmp18407)

tmp18409 := GTl(__e, tmp18408)

tmp18410 := GHd(__e, tmp18409)

tmp18411 := GTl(__e, tmp18410)

tmp18412 := GTl(__e, tmp18411)

tmp18413 := GHd(__e, tmp18412)

tmp18414 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp18413)


var ifres18406 Obj

if True == tmp18414 {
ifres18406 = True


} else {
ifres18406 = False


}

ifres18405 = ifres18406


} else {
ifres18405 = False


}

var ifres18404 Obj

if True == ifres18405 {
ifres18404 = True


} else {
ifres18404 = False


}

ifres18403 = ifres18404


} else {
ifres18403 = False


}

var ifres18402 Obj

if True == ifres18403 {
ifres18402 = True


} else {
ifres18402 = False


}

ifres18401 = ifres18402


} else {
ifres18401 = False


}

var ifres18400 Obj

if True == ifres18401 {
ifres18400 = True


} else {
ifres18400 = False


}

ifres18399 = ifres18400


} else {
ifres18399 = False


}

var ifres18398 Obj

if True == ifres18399 {
ifres18398 = True


} else {
ifres18398 = False


}

ifres18397 = ifres18398


} else {
ifres18397 = False


}

var ifres18396 Obj

if True == ifres18397 {
ifres18396 = True


} else {
ifres18396 = False


}

ifres18395 = ifres18396


} else {
ifres18395 = False


}

var ifres18394 Obj

if True == ifres18395 {
ifres18394 = True


} else {
ifres18394 = False


}

ifres18393 = ifres18394


} else {
ifres18393 = False


}

var ifres18392 Obj

if True == ifres18393 {
ifres18392 = True


} else {
ifres18392 = False


}

ifres18391 = ifres18392


} else {
ifres18391 = False


}

var ifres18390 Obj

if True == ifres18391 {
ifres18390 = True


} else {
ifres18390 = False


}

ifres18389 = ifres18390


} else {
ifres18389 = False


}

var ifres18388 Obj

if True == ifres18389 {
ifres18388 = True


} else {
ifres18388 = False


}

ifres18387 = ifres18388


} else {
ifres18387 = False


}

var ifres18386 Obj

if True == ifres18387 {
ifres18386 = True


} else {
ifres18386 = False


}

ifres18385 = ifres18386


} else {
ifres18385 = False


}

var ifres18384 Obj

if True == ifres18385 {
ifres18384 = True


} else {
ifres18384 = False


}

ifres18383 = ifres18384


} else {
ifres18383 = False


}

var ifres18382 Obj

if True == ifres18383 {
ifres18382 = True


} else {
ifres18382 = False


}

ifres18381 = ifres18382


} else {
ifres18381 = False


}

var ifres18380 Obj

if True == ifres18381 {
ifres18380 = True


} else {
ifres18380 = False


}

ifres18379 = ifres18380


} else {
ifres18379 = False


}

var ifres18378 Obj

if True == ifres18379 {
ifres18378 = True


} else {
ifres18378 = False


}

ifres18377 = ifres18378


} else {
ifres18377 = False


}

var ifres18376 Obj

if True == ifres18377 {
ifres18376 = True


} else {
ifres18376 = False


}

ifres18375 = ifres18376


} else {
ifres18375 = False


}

var ifres18374 Obj

if True == ifres18375 {
ifres18374 = True


} else {
ifres18374 = False


}

ifres18373 = ifres18374


} else {
ifres18373 = False


}

var ifres18372 Obj

if True == ifres18373 {
ifres18372 = True


} else {
ifres18372 = False


}

ifres18371 = ifres18372


} else {
ifres18371 = False


}

var ifres18370 Obj

if True == ifres18371 {
ifres18370 = True


} else {
ifres18370 = False


}

ifres18369 = ifres18370


} else {
ifres18369 = False


}

var ifres18368 Obj

if True == ifres18369 {
ifres18368 = True


} else {
ifres18368 = False


}

ifres18367 = ifres18368


} else {
ifres18367 = False


}

var ifres18366 Obj

if True == ifres18367 {
ifres18366 = True


} else {
ifres18366 = False


}

ifres18365 = ifres18366


} else {
ifres18365 = False


}

var ifres18364 Obj

if True == ifres18365 {
ifres18364 = True


} else {
ifres18364 = False


}

ifres18363 = ifres18364


} else {
ifres18363 = False


}

var ifres18362 Obj

if True == ifres18363 {
ifres18362 = True


} else {
ifres18362 = False


}

ifres18361 = ifres18362


} else {
ifres18361 = False


}

if True == ifres18361 {
tmp18353 := GTl(__e, V383)

tmp18354 := GTl(__e, tmp18353)

tmp18355 := GTl(__e, tmp18354)

tmp18356 := GHd(__e, tmp18355)

tmp18357 := GTl(__e, tmp18356)

tmp18358 := GTl(__e, tmp18357)

tmp18359 := GCons(__e, V384, tmp18358)

return GCons(__e, symtype, tmp18359)


} else {
return V384
}


panic("unreachable")
})

tmp18541 := Call2(__e, ns2_1set, symshen_4use_1type_1info, tmp18352)


_ = tmp18541

tmp18542 := MakeNative1(func(__e *ControlFlow, V387 Obj) Obj {
tmp18552 := PrimGuard1(__e, symvariable_2, PrimIsVariable, V387)

if True == tmp18552 {
return False
} else {
tmp18550 := GConsP(__e, V387)

if True == tmp18550 {
tmp18547 := GHd(__e, V387)

tmp18548 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp18547)


if True == tmp18548 {
tmp18544 := GTl(__e, V387)

tmp18545 := Call1(__e, SymFn(symshen_4monomorphic_2), tmp18544)


if True == tmp18545 {
return True
} else {
return False
}


} else {
return False
}


} else {
return True
}


}


panic("unreachable")
})

tmp18553 := Call2(__e, ns2_1set, symshen_4monomorphic_2, tmp18542)


_ = tmp18553

tmp18554 := MakeNative1(func(__e *ControlFlow, V388 Obj) Obj {
tmp18580 := GConsP(__e, V388)

var ifres18562 Obj

if True == tmp18580 {
tmp18578 := GHd(__e, V388)

tmp18579 := GEq(__e, symprotect, tmp18578)

var ifres18564 Obj

if True == tmp18579 {
tmp18576 := GTl(__e, V388)

tmp18577 := GConsP(__e, tmp18576)

var ifres18566 Obj

if True == tmp18577 {
tmp18573 := GTl(__e, V388)

tmp18574 := GTl(__e, tmp18573)

tmp18575 := GEq(__e, Nil, tmp18574)

var ifres18568 Obj

if True == tmp18575 {
tmp18570 := GTl(__e, V388)

tmp18571 := GHd(__e, tmp18570)

tmp18572 := Call1(__e, SymFn(symshen_4non_1terminal_2), tmp18571)


var ifres18569 Obj

if True == tmp18572 {
ifres18569 = True


} else {
ifres18569 = False


}

ifres18568 = ifres18569


} else {
ifres18568 = False


}

var ifres18567 Obj

if True == ifres18568 {
ifres18567 = True


} else {
ifres18567 = False


}

ifres18566 = ifres18567


} else {
ifres18566 = False


}

var ifres18565 Obj

if True == ifres18566 {
ifres18565 = True


} else {
ifres18565 = False


}

ifres18564 = ifres18565


} else {
ifres18564 = False


}

var ifres18563 Obj

if True == ifres18564 {
ifres18563 = True


} else {
ifres18563 = False


}

ifres18562 = ifres18563


} else {
ifres18562 = False


}

if True == ifres18562 {
tmp18555 := GTl(__e, V388)

return GHd(__e, tmp18555)


} else {
tmp18560 := GConsP(__e, V388)

if True == tmp18560 {
tmp18556 := MakeNative1(func(__e *ControlFlow, Z389 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4process_1yacc_1semantics), Z389)
panic("unreachable")
})

return TailCall2(__e, SymFn(symmap), tmp18556, V388)


} else {
tmp18558 := Call1(__e, SymFn(symshen_4non_1terminal_2), V388)


if True == tmp18558 {
return TailCall2(__e, SymFn(symconcat), symAction, V388)
} else {
return V388
}


}


}


panic("unreachable")
})

tmp18581 := Call2(__e, ns2_1set, symshen_4process_1yacc_1semantics, tmp18554)


_ = tmp18581

tmp18582 := MakeNative1(func(__e *ControlFlow, V392 Obj) Obj {
tmp18595 := GConsP(__e, V392)

var ifres18586 Obj

if True == tmp18595 {
tmp18593 := GTl(__e, V392)

tmp18594 := GConsP(__e, tmp18593)

var ifres18588 Obj

if True == tmp18594 {
tmp18590 := GTl(__e, V392)

tmp18591 := GTl(__e, tmp18590)

tmp18592 := GEq(__e, Nil, tmp18591)

var ifres18589 Obj

if True == tmp18592 {
ifres18589 = True


} else {
ifres18589 = False


}

ifres18588 = ifres18589


} else {
ifres18588 = False


}

var ifres18587 Obj

if True == ifres18588 {
ifres18587 = True


} else {
ifres18587 = False


}

ifres18586 = ifres18587


} else {
ifres18586 = False


}

if True == ifres18586 {
tmp18583 := GTl(__e, V392)

return GHd(__e, tmp18583)


} else {
tmp18584 := GTl(__e, V392)

return GHd(__e, tmp18584)


}


panic("unreachable")
})

tmp18596 := Call2(__e, ns2_1set, symshen_4_5_1out, tmp18582)


_ = tmp18596

tmp18597 := MakeNative1(func(__e *ControlFlow, V393 Obj) Obj {
return GHd(__e, V393)
panic("unreachable")
})

tmp18598 := Call2(__e, ns2_1set, symshen_4in_1_6, tmp18597)


_ = tmp18598

tmp18599 := MakeNative1(func(__e *ControlFlow, V394 Obj) Obj {
tmp18600 := GCons(__e, V394, Nil)

return GCons(__e, Nil, tmp18600)


panic("unreachable")
})

tmp18601 := Call2(__e, ns2_1set, sym_5_b_6, tmp18599)


_ = tmp18601

tmp18602 := MakeNative1(func(__e *ControlFlow, V395 Obj) Obj {
tmp18603 := GCons(__e, Nil, Nil)

return GCons(__e, V395, tmp18603)


panic("unreachable")
})

tmp18604 := Call2(__e, ns2_1set, sym_5e_6, tmp18602)


_ = tmp18604

tmp18605 := MakeNative1(func(__e *ControlFlow, V398 Obj) Obj {
tmp18608 := GEq(__e, Nil, V398)

if True == tmp18608 {
tmp18606 := GCons(__e, Nil, Nil)

return GCons(__e, Nil, tmp18606)


} else {
return TailCall0(__e, SymFn(symshen_4parse_1failure))
}


panic("unreachable")
})

return TailCall2(__e, ns2_1set, sym_5end_6, tmp18605)




panic("unreachable")
})

