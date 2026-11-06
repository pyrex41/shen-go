package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp17721 := MakeNative2(func(__e *ControlFlow, V6010, V6011 Obj) Obj {
tmp17722 := Call1(__e, SymFn(symshen_4rectify_1type), V6011)


let__17692 := tmp17722
_ = let__17692

tmp17723 := MakeNative1(func(__e *ControlFlow, Z6014 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6015 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6016 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6017 Obj) Obj {
tmp17724 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17724

tmp17725 := Call2(__e, SymFn(symshen_4deref), V6010, Z6014)


tmp17726 := Call1(__e, SymFn(symreceive), tmp17725)


tmp17727 := Call2(__e, SymFn(symshen_4deref), let__17692, Z6014)


tmp17728 := Call1(__e, SymFn(symreceive), tmp17727)


return TailCall6(__e, SymFn(symshen_4variancy), tmp17726, tmp17728, Z6014, Z6015, Z6016, Z6017)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp17729 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp17730 := Call1(__e, tmp17723, tmp17729)


tmp17731 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp17732 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp17731)


tmp17733 := Call2(__e, SymFn(sym_8v), True, tmp17732)


tmp17734 := Call1(__e, tmp17730, tmp17733)


tmp17735 := Call1(__e, tmp17734, MakeInteger(0))


tmp17736 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

tmp17737 := Call1(__e, tmp17735, tmp17736)


let__17693 := tmp17737
_ = let__17693

tmp17738 := Call1(__e, SymFn(symshen_4prolog_1abstraction), V6011)


tmp17739 := Call1(__e, SymFn(symeval_1kl), tmp17738)


let__17694 := tmp17739
_ = let__17694

tmp17740 := GValue(__e, symshen_4_dsigf_d)

tmp17741 := Call3(__e, SymFn(symshen_4assoc_1_6), V6010, let__17694, tmp17740)


tmp17742 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp17741)

let__17695 := tmp17742
_ = let__17695

return V6010


panic("unreachable")
})

tmp17743 := Call2(__e, ns2_1set, symdeclare, tmp17721)


_ = tmp17743

tmp17744 := MakeNative6(func(__e *ControlFlow, V6020, V6021, V6022, V6023, V6024, V6025 Obj) Obj {
tmp17755 := Call1(__e, SymFn(symshen_4unlocked_2), V6023)


if True == tmp17755 {
tmp17745 := Call1(__e, SymFn(symshen_4newpv), V6022)


let__17696 := tmp17745
_ = let__17696

tmp17746 := Call1(__e, SymFn(symshen_4newpv), V6022)


let__17697 := tmp17746
_ = let__17697

tmp17747 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17747

tmp17748 := GCons(__e, V6020, Nil)

tmp17749 := GCons(__e, symfn, tmp17748)

tmp17750 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17751 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4variants_2), V6020, let__17697, V6021, V6022, V6023, V6024, V6025)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4variancy_1signature), V6020, let__17696, let__17697, V6022, V6023, V6024, tmp17751)
return nil


panic("unreachable")
})

tmp17752 := Call(__e, SymFn(symshen_4system_1S_1h), tmp17749, let__17696, Nil, V6022, V6023, V6024, tmp17750)


tmp17753 := Call2(__e, SymFn(symshen_4gc), V6022, tmp17752)


return TailCall2(__e, SymFn(symshen_4gc), V6022, tmp17753)


} else {
return False
}


panic("unreachable")
})

tmp17756 := Call2(__e, ns2_1set, symshen_4variancy, tmp17744)


_ = tmp17756

tmp17757 := MakeNative(func(__e *ControlFlow) {
V6028 := __e.Get(1)
V6029 := __e.Get(2)
V6030 := __e.Get(3)
V6031 := __e.Get(4)
V6032 := __e.Get(5)
V6033 := __e.Get(6)
V6034 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__17698 := GAdd(__e, V6033, MakeInteger(1))
_ = let__17698

tmp17806 := Call1(__e, SymFn(symshen_4unlocked_2), V6032)


var ifres17759 Obj

if True == tmp17806 {
tmp17760 := Call2(__e, SymFn(symshen_4lazyderef), V6030, V6031)


let__17700 := tmp17760
_ = let__17700

tmp17761 := MakeNative1(func(__e *ControlFlow, Z6039 Obj) Obj {
tmp17762 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17762

tmp17763 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17764 := Call2(__e, SymFn(symshen_4deref), V6028, V6031)


tmp17765 := Call1(__e, SymFn(symarity), tmp17764)


tmp17766 := GEq(__e, tmp17765, MakeInteger(0))

tmp17767 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall4(__e, SymFn(symshen_4cut), V6031, V6032, let__17698, V6034)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp17766, V6031, V6032, let__17698, tmp17767)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), V6029, Z6039, V6031, V6032, let__17698, tmp17763)


panic("unreachable")
})

let__17701 := tmp17761
_ = let__17701

tmp17805 := GConsP(__e, let__17700)

var ifres17768 Obj

if True == tmp17805 {
tmp17769 := GHd(__e, let__17700)

tmp17770 := Call2(__e, SymFn(symshen_4lazyderef), tmp17769, V6031)


let__17702 := tmp17770
_ = let__17702

tmp17771 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17772 := GTl(__e, let__17700)

tmp17773 := Call2(__e, SymFn(symshen_4lazyderef), tmp17772, V6031)


let__17704 := tmp17773
_ = let__17704

tmp17774 := MakeNative1(func(__e *ControlFlow, Z6044 Obj) Obj {
return TailCall1(__e, let__17701, Z6044)
panic("unreachable")
})

let__17705 := tmp17774
_ = let__17705

tmp17790 := GConsP(__e, let__17704)

if True == tmp17790 {
tmp17775 := GHd(__e, let__17704)

let__17706 := tmp17775
_ = let__17706

tmp17776 := GTl(__e, let__17704)

tmp17777 := Call2(__e, SymFn(symshen_4lazyderef), tmp17776, V6031)


let__17707 := tmp17777
_ = let__17707

tmp17778 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17705, let__17706)
panic("unreachable")
})

let__17708 := tmp17778
_ = let__17708

tmp17782 := GEq(__e, let__17707, Nil)

if True == tmp17782 {
return TailCall1(__e, SymFn(symthaw), let__17708)
} else {
tmp17780 := Call1(__e, SymFn(symshen_4pvar_2), let__17707)


if True == tmp17780 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__17707, Nil, V6031, let__17708)
} else {
return False
}


}


} else {
tmp17788 := Call1(__e, SymFn(symshen_4pvar_2), let__17704)


if True == tmp17788 {
tmp17783 := Call1(__e, SymFn(symshen_4newpv), V6031)


let__17709 := tmp17783
_ = let__17709

tmp17784 := GCons(__e, let__17709, Nil)

tmp17785 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17705, let__17709)
panic("unreachable")
})

tmp17786 := Call4(__e, SymFn(symshen_4bind_b), let__17704, tmp17784, V6031, tmp17785)


return TailCall2(__e, SymFn(symshen_4gc), V6031, tmp17786)


} else {
return False
}


}


panic("unreachable")
})

let__17703 := tmp17771
_ = let__17703

tmp17796 := GEq(__e, let__17702, sym_1_1_6)

var ifres17791 Obj

if True == tmp17796 {
tmp17792 := Call1(__e, SymFn(symthaw), let__17703)


ifres17791 = tmp17792


} else {
tmp17795 := Call1(__e, SymFn(symshen_4pvar_2), let__17702)


var ifres17793 Obj

if True == tmp17795 {
tmp17794 := Call4(__e, SymFn(symshen_4bind_b), let__17702, sym_1_1_6, V6031, let__17703)


ifres17793 = tmp17794


} else {
ifres17793 = False


}

ifres17791 = ifres17793


}

ifres17768 = ifres17791


} else {
tmp17804 := Call1(__e, SymFn(symshen_4pvar_2), let__17700)


var ifres17797 Obj

if True == tmp17804 {
tmp17798 := Call1(__e, SymFn(symshen_4newpv), V6031)


let__17710 := tmp17798
_ = let__17710

tmp17799 := GCons(__e, let__17710, Nil)

tmp17800 := GCons(__e, sym_1_1_6, tmp17799)

tmp17801 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17701, let__17710)
panic("unreachable")
})

tmp17802 := Call4(__e, SymFn(symshen_4bind_b), let__17700, tmp17800, V6031, tmp17801)


tmp17803 := Call2(__e, SymFn(symshen_4gc), V6031, tmp17802)


ifres17797 = tmp17803


} else {
ifres17797 = False


}

ifres17768 = ifres17797


}

ifres17759 = ifres17768


} else {
ifres17759 = False


}

let__17699 := ifres17759
_ = let__17699

tmp17814 := GEq(__e, let__17699, False)

if True == tmp17814 {
tmp17810 := Call1(__e, SymFn(symshen_4unlocked_2), V6032)


var ifres17807 Obj

if True == tmp17810 {
tmp17808 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17808

tmp17809 := Call6(__e, SymFn(symis_b), V6029, V6030, V6031, V6032, let__17698, V6034)


ifres17807 = tmp17809


} else {
ifres17807 = False


}

let__17711 := ifres17807
_ = let__17711

tmp17812 := GEq(__e, let__17711, False)

if True == tmp17812 {
return TailCall2(__e, SymFn(symshen_4unlock), V6032, let__17698)
} else {
return let__17711
}


} else {
return let__17699
}


panic("unreachable")
}())
}, 7)

tmp17815 := Call2(__e, ns2_1set, symshen_4variancy_1signature, tmp17757)


_ = tmp17815

tmp17816 := MakeNative(func(__e *ControlFlow) {
V6051 := __e.Get(1)
V6052 := __e.Get(2)
V6053 := __e.Get(3)
V6054 := __e.Get(4)
V6055 := __e.Get(5)
V6056 := __e.Get(6)
V6057 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp17820 := Call1(__e, SymFn(symshen_4unlocked_2), V6055)


var ifres17817 Obj

if True == tmp17820 {
tmp17818 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17818

tmp17819 := Call6(__e, SymFn(symis_b), V6052, V6053, V6054, V6055, V6056, V6057)


ifres17817 = tmp17819


} else {
ifres17817 = False


}

let__17712 := ifres17817
_ = let__17712

tmp17832 := GEq(__e, let__17712, False)

if True == tmp17832 {
tmp17830 := Call1(__e, SymFn(symshen_4unlocked_2), V6055)


if True == tmp17830 {
tmp17821 := Call1(__e, SymFn(symshen_4newpv), V6054)


let__17713 := tmp17821
_ = let__17713

tmp17822 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17822

tmp17823 := Call2(__e, SymFn(symshen_4deref), V6051, V6054)


tmp17824 := Call3(__e, SymFn(symshen_4app), tmp17823, MakeString(" may create errors\n"), symshen_4a)


tmp17825 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("warning: changing the type of "), tmp17824)

tmp17826 := Call0(__e, SymFn(symstoutput))


tmp17827 := Call2(__e, SymFn(sympr), tmp17825, tmp17826)


tmp17828 := Call6(__e, SymFn(symis), let__17713, tmp17827, V6054, V6055, V6056, V6057)


return TailCall2(__e, SymFn(symshen_4gc), V6054, tmp17828)


} else {
return False
}


} else {
return let__17712
}


panic("unreachable")
}())
}, 7)

tmp17833 := Call2(__e, ns2_1set, symshen_4variants_2, tmp17816)


_ = tmp17833

tmp17834 := MakeNative1(func(__e *ControlFlow, V6060 Obj) Obj {
tmp17835 := Call1(__e, SymFn(symgensym), symB)


let__17714 := tmp17835
_ = let__17714

tmp17836 := Call1(__e, SymFn(symgensym), symL)


let__17715 := tmp17836
_ = let__17715

tmp17837 := Call1(__e, SymFn(symgensym), symKey)


let__17716 := tmp17837
_ = let__17716

tmp17838 := Call1(__e, SymFn(symgensym), symC)


let__17717 := tmp17838
_ = let__17717

tmp17839 := Call1(__e, SymFn(symgensym), symV)


let__17718 := tmp17839
_ = let__17718

tmp17840 := Call1(__e, SymFn(symshen_4extract_1vars), V6060)


let__17719 := tmp17840
_ = let__17719

tmp17841 := Call1(__e, SymFn(symshen_4rcons__form), V6060)


tmp17842 := GCons(__e, let__17717, Nil)

tmp17843 := GCons(__e, let__17716, tmp17842)

tmp17844 := GCons(__e, let__17715, tmp17843)

tmp17845 := GCons(__e, let__17714, tmp17844)

tmp17846 := GCons(__e, tmp17841, tmp17845)

tmp17847 := GCons(__e, let__17718, tmp17846)

tmp17848 := GCons(__e, symis_b, tmp17847)

tmp17849 := Call3(__e, SymFn(symshen_4stpart), let__17719, tmp17848, let__17714)


tmp17850 := GCons(__e, tmp17849, Nil)

tmp17851 := GCons(__e, let__17717, tmp17850)

tmp17852 := GCons(__e, symlambda, tmp17851)

tmp17853 := GCons(__e, tmp17852, Nil)

tmp17854 := GCons(__e, let__17716, tmp17853)

tmp17855 := GCons(__e, symlambda, tmp17854)

tmp17856 := GCons(__e, tmp17855, Nil)

tmp17857 := GCons(__e, let__17715, tmp17856)

tmp17858 := GCons(__e, symlambda, tmp17857)

tmp17859 := GCons(__e, tmp17858, Nil)

tmp17860 := GCons(__e, let__17714, tmp17859)

tmp17861 := GCons(__e, symlambda, tmp17860)

tmp17862 := GCons(__e, tmp17861, Nil)

tmp17863 := GCons(__e, let__17718, tmp17862)

return GCons(__e, symlambda, tmp17863)


panic("unreachable")
})

tmp17864 := Call2(__e, ns2_1set, symshen_4prolog_1abstraction, tmp17834)


_ = tmp17864

tmp17865 := MakeNative1(func(__e *ControlFlow, V6067 Obj) Obj {
tmp17866 := GValue(__e, symshen_4_ddemodulation_1function_d)

let__17720 := tmp17866
_ = let__17720

return TailCall1(__e, let__17720, V6067)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4demod, tmp17865)




panic("unreachable")
})

