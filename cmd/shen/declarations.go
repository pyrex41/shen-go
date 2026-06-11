package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp6644 := MakeNative(func(__e *ControlFlow) {
tmp6645 := MakeNative(func(__e *ControlFlow) {
Z904 := __e.Get(1)
_ = Z904
__e.TailApply(PrimFunc(symshen_4typename), Z904)
return
}, 1)

tmp6646 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp6645, tmp6646)
return


}, 0)

tmp6647 := Call(__e, ns2_1set, symdatatypes, tmp6644)


_ = tmp6647

tmp6648 := MakeNative(func(__e *ControlFlow) {
tmp6649 := MakeNative(func(__e *ControlFlow) {
Z905 := __e.Get(1)
_ = Z905
__e.TailApply(PrimFunc(symshen_4typename), Z905)
return
}, 1)

tmp6650 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp6649, tmp6650)
return


}, 0)

tmp6651 := Call(__e, ns2_1set, symshen_4included, tmp6648)


_ = tmp6651

tmp6652 := MakeNative(func(__e *ControlFlow) {
V908 := __e.Get(1)
_ = V908
tmp6657 := PrimIsPair(V908)

if True == tmp6657 {
tmp6653 := PrimHead(V908)

tmp6654 := PrimStr(tmp6653)

tmp6655 := Call(__e, PrimFunc(symshen_4typename_1h), tmp6654)


__e.Return(PrimIntern(tmp6655))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename)
return
}


}, 1)

tmp6658 := Call(__e, ns2_1set, symshen_4typename, tmp6652)


_ = tmp6658

tmp6659 := MakeNative(func(__e *ControlFlow) {
V909 := __e.Get(1)
_ = V909
tmp6666 := PrimEqual(MakeString("#type"), V909)

if True == tmp6666 {
__e.Return(MakeString(""))
return
} else {
tmp6664 := Call(__e, PrimFunc(symshen_4_7string_2), V909)


if True == tmp6664 {
tmp6660 := Call(__e, PrimFunc(symhdstr), V909)


tmp6661 := PrimTailString(V909)

tmp6662 := Call(__e, PrimFunc(symshen_4typename_1h), tmp6661)


__e.Return(PrimStringConcat(tmp6660, tmp6662))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4typename_1h)
return
}


}


}, 1)

tmp6667 := Call(__e, ns2_1set, symshen_4typename_1h, tmp6659)


_ = tmp6667

tmp6668 := MakeNative(func(__e *ControlFlow) {
V910 := __e.Get(1)
_ = V910
tmp6672 := PrimLessThan(V910, MakeNumber(0))

if True == tmp6672 {
__e.Return(PrimValue(symshen_4_dprolog_1memory_d))
return
} else {
tmp6670 := PrimIsInteger(V910)

if True == tmp6670 {
__e.Return(PrimSet(symshen_4_dprolog_1memory_d, V910))
return
} else {
__e.Return(PrimSimpleError(MakeString("prolog memory expects an integer value\n")))
return
}


}


}, 1)

tmp6673 := Call(__e, ns2_1set, symprolog_1memory, tmp6668)


_ = tmp6673

tmp6674 := MakeNative(func(__e *ControlFlow) {
V911 := __e.Get(1)
_ = V911
tmp6675 := MakeNative(func(__e *ControlFlow) {
tmp6676 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V911, symarity, tmp6676)
return


}, 0)

tmp6677 := MakeNative(func(__e *ControlFlow) {
Z912 := __e.Get(1)
_ = Z912
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp6675, tmp6677)
return


}, 1)

tmp6678 := Call(__e, ns2_1set, symarity, tmp6674)


_ = tmp6678

tmp6679 := MakeNative(func(__e *ControlFlow) {
V915 := __e.Get(1)
_ = V915
tmp6695 := PrimEqual(Nil, V915)

if True == tmp6695 {
__e.Return(Nil)
return
} else {
tmp6693 := PrimIsPair(V915)

var ifres6689 Obj

if True == tmp6693 {
tmp6691 := PrimTail(V915)

tmp6692 := PrimIsPair(tmp6691)

var ifres6690 Obj

if True == tmp6692 {
ifres6690 = True


} else {
ifres6690 = False


}

ifres6689 = ifres6690


} else {
ifres6689 = False


}

if True == ifres6689 {
tmp6680 := MakeNative(func(__e *ControlFlow) {
W916 := __e.Get(1)
_ = W916
tmp6681 := PrimTail(V915)

tmp6682 := PrimTail(tmp6681)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp6682)
return


}, 1)

tmp6683 := PrimHead(V915)

tmp6684 := PrimTail(V915)

tmp6685 := PrimHead(tmp6684)

tmp6686 := PrimValue(sym_dproperty_1vector_d)

tmp6687 := Call(__e, PrimFunc(symput), tmp6683, symarity, tmp6685, tmp6686)


__e.TailApply(tmp6680, tmp6687)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp6696 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp6679)


_ = tmp6696

tmp6697 := MakeNative(func(__e *ControlFlow) {
V917 := __e.Get(1)
_ = V917
tmp6698 := MakeNative(func(__e *ControlFlow) {
W918 := __e.Get(1)
_ = W918
tmp6699 := MakeNative(func(__e *ControlFlow) {
W919 := __e.Get(1)
_ = W919
__e.Return(V917)
return
}, 1)

tmp6700 := Call(__e, PrimFunc(symadjoin), V917, W918)


tmp6701 := PrimValue(sym_dproperty_1vector_d)

tmp6702 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp6700, tmp6701)


__e.TailApply(tmp6699, tmp6702)
return


}, 1)

tmp6703 := PrimValue(sym_dproperty_1vector_d)

tmp6704 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6703)


__e.TailApply(tmp6698, tmp6704)
return


}, 1)

tmp6705 := Call(__e, ns2_1set, symsystemf, tmp6697)


_ = tmp6705

tmp6706 := MakeNative(func(__e *ControlFlow) {
V920 := __e.Get(1)
_ = V920
V921 := __e.Get(2)
_ = V921
tmp6708 := Call(__e, PrimFunc(symelement_2), V920, V921)


if True == tmp6708 {
__e.Return(V921)
return
} else {
__e.Return(PrimCons(V920, V921))
return
}


}, 2)

tmp6709 := Call(__e, ns2_1set, symadjoin, tmp6706)


_ = tmp6709

tmp6710 := MakeNative(func(__e *ControlFlow) {
V922 := __e.Get(1)
_ = V922
tmp6711 := MakeNative(func(__e *ControlFlow) {
W923 := __e.Get(1)
_ = W923
tmp6719 := PrimEqual(W923, MakeNumber(-1))

var ifres6716 Obj

if True == tmp6719 {
ifres6716 = True


} else {
tmp6718 := PrimEqual(W923, MakeNumber(0))

var ifres6717 Obj

if True == tmp6718 {
ifres6717 = True


} else {
ifres6717 = False


}

ifres6716 = ifres6717


}

if True == ifres6716 {
__e.Return(Nil)
return
} else {
tmp6712 := PrimCons(V922, Nil)

tmp6713 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp6712, W923)


tmp6714 := Call(__e, PrimFunc(symeval_1kl), tmp6713)


__e.Return(PrimCons(V922, tmp6714))
return


}


}, 1)

tmp6720 := Call(__e, PrimFunc(symarity), V922)


__e.TailApply(tmp6711, tmp6720)
return


}, 1)

tmp6721 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp6710)


_ = tmp6721

tmp6722 := MakeNative(func(__e *ControlFlow) {
V924 := __e.Get(1)
_ = V924
tmp6727 := PrimIsPair(V924)

if True == tmp6727 {
tmp6723 := PrimHead(V924)

tmp6724 := PrimTail(V924)

tmp6725 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), tmp6723, symshen_4lambda_1form, tmp6724, tmp6725)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4set_1lambda_1form_1entry)
return
}


}, 1)

tmp6728 := Call(__e, ns2_1set, symshen_4set_1lambda_1form_1entry, tmp6722)


_ = tmp6728

tmp6729 := MakeNative(func(__e *ControlFlow) {
V925 := __e.Get(1)
_ = V925
tmp6730 := MakeNative(func(__e *ControlFlow) {
W926 := __e.Get(1)
_ = W926
tmp6731 := MakeNative(func(__e *ControlFlow) {
Z928 := __e.Get(1)
_ = Z928
__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), Z928)
return
}, 1)

tmp6732 := MakeNative(func(__e *ControlFlow) {
Z929 := __e.Get(1)
_ = Z929
__e.TailApply(PrimFunc(symshen_4tuple), Z929)
return
}, 1)

tmp6733 := PrimCons(symshen_4tuple, tmp6732)

tmp6734 := MakeNative(func(__e *ControlFlow) {
Z930 := __e.Get(1)
_ = Z930
__e.TailApply(PrimFunc(symshen_4pvar), Z930)
return
}, 1)

tmp6735 := PrimCons(symshen_4pvar, tmp6734)

tmp6736 := MakeNative(func(__e *ControlFlow) {
Z931 := __e.Get(1)
_ = Z931
__e.TailApply(PrimFunc(symshen_4dictionary), Z931)
return
}, 1)

tmp6737 := PrimCons(symshen_4dictionary, tmp6736)

tmp6738 := MakeNative(func(__e *ControlFlow) {
Z932 := __e.Get(1)
_ = Z932
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z932)
return
}, 1)

tmp6739 := PrimCons(symshen_4print_1prolog_1vector, tmp6738)

tmp6740 := MakeNative(func(__e *ControlFlow) {
Z933 := __e.Get(1)
_ = Z933
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z933)
return
}, 1)

tmp6741 := PrimCons(symshen_4print_1freshterm, tmp6740)

tmp6742 := MakeNative(func(__e *ControlFlow) {
Z934 := __e.Get(1)
_ = Z934
__e.TailApply(PrimFunc(symshen_4printF), Z934)
return
}, 1)

tmp6743 := PrimCons(symshen_4printF, tmp6742)

tmp6744 := PrimCons(tmp6743, W926)

tmp6745 := PrimCons(tmp6741, tmp6744)

tmp6746 := PrimCons(tmp6739, tmp6745)

tmp6747 := PrimCons(tmp6737, tmp6746)

tmp6748 := PrimCons(tmp6735, tmp6747)

tmp6749 := PrimCons(tmp6733, tmp6748)

__e.TailApply(PrimFunc(symshen_4for_1each), tmp6731, tmp6749)
return


}, 1)

tmp6750 := MakeNative(func(__e *ControlFlow) {
Z927 := __e.Get(1)
_ = Z927
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z927)
return
}, 1)

tmp6751 := Call(__e, PrimFunc(symmap), tmp6750, V925)


__e.TailApply(tmp6730, tmp6751)
return


}, 1)

__e.TailApply(ns2_1set, symshen_4build_1lambda_1table, tmp6729)
return




}, 0)

