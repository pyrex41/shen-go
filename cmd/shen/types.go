package main

import . "github.com/pyrex41/shen-go/kl"

var TypesMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp17741 := MakeNative2(func(__e *ControlFlow, V6028, V6029 Obj) Obj {
tmp17742 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17743 := Call1(__e, SymFn(symshen_4prolog_1abstraction), V6029)


return TailCall1(__e, SymFn(symeval_1kl), tmp17743)


panic("unreachable")
})

return TailCall3(__e, SymFn(symshen_4declare_1with_1signature), V6028, V6029, tmp17742)


panic("unreachable")
})

tmp17744 := Call2(__e, ns2_1set, symdeclare, tmp17741)


_ = tmp17744

tmp17745 := MakeNative3(func(__e *ControlFlow, V6030, V6031, V6032 Obj) Obj {
tmp17746 := Call1(__e, SymFn(symshen_4rectify_1type), V6031)


let__17712 := tmp17746
_ = let__17712

tmp17747 := MakeNative1(func(__e *ControlFlow, Z6035 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6036 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6037 Obj) Obj {
return MakeNative1(func(__e *ControlFlow, Z6038 Obj) Obj {
tmp17748 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17748

tmp17749 := Call2(__e, SymFn(symshen_4deref), V6030, Z6035)


tmp17750 := Call1(__e, SymFn(symreceive), tmp17749)


tmp17751 := Call2(__e, SymFn(symshen_4deref), let__17712, Z6035)


tmp17752 := Call1(__e, SymFn(symreceive), tmp17751)


return TailCall6(__e, SymFn(symshen_4variancy), tmp17750, tmp17752, Z6035, Z6036, Z6037, Z6038)


panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})
panic("unreachable")
})

tmp17753 := Call0(__e, SymFn(symshen_4prolog_1vector))


tmp17754 := Call1(__e, tmp17747, tmp17753)


tmp17755 := Call1(__e, SymFn(symvector), MakeInteger(0))


tmp17756 := Call2(__e, SymFn(sym_8v), MakeInteger(0), tmp17755)


tmp17757 := Call2(__e, SymFn(sym_8v), True, tmp17756)


tmp17758 := Call1(__e, tmp17754, tmp17757)


tmp17759 := Call1(__e, tmp17758, MakeInteger(0))


tmp17760 := MakeNative0(func(__e *ControlFlow) Obj {
return True
panic("unreachable")
})

tmp17761 := Call1(__e, tmp17759, tmp17760)


let__17713 := tmp17761
_ = let__17713

tmp17762 := Call1(__e, SymFn(symthaw), V6032)


let__17714 := tmp17762
_ = let__17714

tmp17763 := GValue(__e, symshen_4_dsigf_d)

tmp17764 := Call3(__e, SymFn(symshen_4assoc_1_6), V6030, let__17714, tmp17763)


tmp17765 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsigf_d, tmp17764)

let__17715 := tmp17765
_ = let__17715

return V6030


panic("unreachable")
})

tmp17766 := Call2(__e, ns2_1set, symshen_4declare_1with_1signature, tmp17745)


_ = tmp17766

tmp17767 := MakeNative6(func(__e *ControlFlow, V6041, V6042, V6043, V6044, V6045, V6046 Obj) Obj {
tmp17778 := Call1(__e, SymFn(symshen_4unlocked_2), V6044)


if True == tmp17778 {
tmp17768 := Call1(__e, SymFn(symshen_4newpv), V6043)


let__17716 := tmp17768
_ = let__17716

tmp17769 := Call1(__e, SymFn(symshen_4newpv), V6043)


let__17717 := tmp17769
_ = let__17717

tmp17770 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17770

tmp17771 := GCons(__e, V6041, Nil)

tmp17772 := GCons(__e, symfn, tmp17771)

tmp17773 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17774 := MakeNative0(func(__e *ControlFlow) Obj {
__e.TailApply(SymFn(symshen_4variants_2), V6041, let__17717, V6042, V6043, V6044, V6045, V6046)
return nil
panic("unreachable")
})

__e.TailApply(SymFn(symshen_4variancy_1signature), V6041, let__17716, let__17717, V6043, V6044, V6045, tmp17774)
return nil


panic("unreachable")
})

tmp17775 := Call(__e, SymFn(symshen_4system_1S_1h), tmp17772, let__17716, Nil, V6043, V6044, V6045, tmp17773)


tmp17776 := Call2(__e, SymFn(symshen_4gc), V6043, tmp17775)


return TailCall2(__e, SymFn(symshen_4gc), V6043, tmp17776)


} else {
return False
}


panic("unreachable")
})

tmp17779 := Call2(__e, ns2_1set, symshen_4variancy, tmp17767)


_ = tmp17779

tmp17780 := MakeNative(func(__e *ControlFlow) {
V6049 := __e.Get(1)
V6050 := __e.Get(2)
V6051 := __e.Get(3)
V6052 := __e.Get(4)
V6053 := __e.Get(5)
V6054 := __e.Get(6)
V6055 := __e.Get(7)
__e.ReturnDirect(func() Obj {
let__17718 := GAdd(__e, V6054, MakeInteger(1))
_ = let__17718

tmp17829 := Call1(__e, SymFn(symshen_4unlocked_2), V6053)


var ifres17782 Obj

if True == tmp17829 {
tmp17783 := Call2(__e, SymFn(symshen_4lazyderef), V6051, V6052)


let__17720 := tmp17783
_ = let__17720

tmp17784 := MakeNative1(func(__e *ControlFlow, Z6060 Obj) Obj {
tmp17785 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17785

tmp17786 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17787 := Call2(__e, SymFn(symshen_4deref), V6049, V6052)


tmp17788 := Call1(__e, SymFn(symarity), tmp17787)


tmp17789 := GEq(__e, tmp17788, MakeInteger(0))

tmp17790 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall4(__e, SymFn(symshen_4cut), V6052, V6053, let__17718, V6055)
panic("unreachable")
})

return TailCall5(__e, SymFn(symwhen), tmp17789, V6052, V6053, let__17718, tmp17790)


panic("unreachable")
})

return TailCall6(__e, SymFn(symis_b), V6050, Z6060, V6052, V6053, let__17718, tmp17786)


panic("unreachable")
})

let__17721 := tmp17784
_ = let__17721

tmp17828 := GConsP(__e, let__17720)

var ifres17791 Obj

if True == tmp17828 {
tmp17792 := GHd(__e, let__17720)

tmp17793 := Call2(__e, SymFn(symshen_4lazyderef), tmp17792, V6052)


let__17722 := tmp17793
_ = let__17722

tmp17794 := MakeNative0(func(__e *ControlFlow) Obj {
tmp17795 := GTl(__e, let__17720)

tmp17796 := Call2(__e, SymFn(symshen_4lazyderef), tmp17795, V6052)


let__17724 := tmp17796
_ = let__17724

tmp17797 := MakeNative1(func(__e *ControlFlow, Z6065 Obj) Obj {
return TailCall1(__e, let__17721, Z6065)
panic("unreachable")
})

let__17725 := tmp17797
_ = let__17725

tmp17813 := GConsP(__e, let__17724)

if True == tmp17813 {
tmp17798 := GHd(__e, let__17724)

let__17726 := tmp17798
_ = let__17726

tmp17799 := GTl(__e, let__17724)

tmp17800 := Call2(__e, SymFn(symshen_4lazyderef), tmp17799, V6052)


let__17727 := tmp17800
_ = let__17727

tmp17801 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17725, let__17726)
panic("unreachable")
})

let__17728 := tmp17801
_ = let__17728

tmp17805 := GEq(__e, let__17727, Nil)

if True == tmp17805 {
return TailCall1(__e, SymFn(symthaw), let__17728)
} else {
tmp17803 := Call1(__e, SymFn(symshen_4pvar_2), let__17727)


if True == tmp17803 {
return TailCall4(__e, SymFn(symshen_4bind_b), let__17727, Nil, V6052, let__17728)
} else {
return False
}


}


} else {
tmp17811 := Call1(__e, SymFn(symshen_4pvar_2), let__17724)


if True == tmp17811 {
tmp17806 := Call1(__e, SymFn(symshen_4newpv), V6052)


let__17729 := tmp17806
_ = let__17729

tmp17807 := GCons(__e, let__17729, Nil)

tmp17808 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17725, let__17729)
panic("unreachable")
})

tmp17809 := Call4(__e, SymFn(symshen_4bind_b), let__17724, tmp17807, V6052, tmp17808)


return TailCall2(__e, SymFn(symshen_4gc), V6052, tmp17809)


} else {
return False
}


}


panic("unreachable")
})

let__17723 := tmp17794
_ = let__17723

tmp17819 := GEq(__e, let__17722, sym_1_1_6)

var ifres17814 Obj

if True == tmp17819 {
tmp17815 := Call1(__e, SymFn(symthaw), let__17723)


ifres17814 = tmp17815


} else {
tmp17818 := Call1(__e, SymFn(symshen_4pvar_2), let__17722)


var ifres17816 Obj

if True == tmp17818 {
tmp17817 := Call4(__e, SymFn(symshen_4bind_b), let__17722, sym_1_1_6, V6052, let__17723)


ifres17816 = tmp17817


} else {
ifres17816 = False


}

ifres17814 = ifres17816


}

ifres17791 = ifres17814


} else {
tmp17827 := Call1(__e, SymFn(symshen_4pvar_2), let__17720)


var ifres17820 Obj

if True == tmp17827 {
tmp17821 := Call1(__e, SymFn(symshen_4newpv), V6052)


let__17730 := tmp17821
_ = let__17730

tmp17822 := GCons(__e, let__17730, Nil)

tmp17823 := GCons(__e, sym_1_1_6, tmp17822)

tmp17824 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, let__17721, let__17730)
panic("unreachable")
})

tmp17825 := Call4(__e, SymFn(symshen_4bind_b), let__17720, tmp17823, V6052, tmp17824)


tmp17826 := Call2(__e, SymFn(symshen_4gc), V6052, tmp17825)


ifres17820 = tmp17826


} else {
ifres17820 = False


}

ifres17791 = ifres17820


}

ifres17782 = ifres17791


} else {
ifres17782 = False


}

let__17719 := ifres17782
_ = let__17719

tmp17837 := GEq(__e, let__17719, False)

if True == tmp17837 {
tmp17833 := Call1(__e, SymFn(symshen_4unlocked_2), V6053)


var ifres17830 Obj

if True == tmp17833 {
tmp17831 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17831

tmp17832 := Call6(__e, SymFn(symis_b), V6050, V6051, V6052, V6053, let__17718, V6055)


ifres17830 = tmp17832


} else {
ifres17830 = False


}

let__17731 := ifres17830
_ = let__17731

tmp17835 := GEq(__e, let__17731, False)

if True == tmp17835 {
return TailCall2(__e, SymFn(symshen_4unlock), V6053, let__17718)
} else {
return let__17731
}


} else {
return let__17719
}


panic("unreachable")
}())
}, 7)

tmp17838 := Call2(__e, ns2_1set, symshen_4variancy_1signature, tmp17780)


_ = tmp17838

tmp17839 := MakeNative(func(__e *ControlFlow) {
V6072 := __e.Get(1)
V6073 := __e.Get(2)
V6074 := __e.Get(3)
V6075 := __e.Get(4)
V6076 := __e.Get(5)
V6077 := __e.Get(6)
V6078 := __e.Get(7)
__e.ReturnDirect(func() Obj {
tmp17843 := Call1(__e, SymFn(symshen_4unlocked_2), V6076)


var ifres17840 Obj

if True == tmp17843 {
tmp17841 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17841

tmp17842 := Call6(__e, SymFn(symis_b), V6073, V6074, V6075, V6076, V6077, V6078)


ifres17840 = tmp17842


} else {
ifres17840 = False


}

let__17732 := ifres17840
_ = let__17732

tmp17855 := GEq(__e, let__17732, False)

if True == tmp17855 {
tmp17853 := Call1(__e, SymFn(symshen_4unlocked_2), V6076)


if True == tmp17853 {
tmp17844 := Call1(__e, SymFn(symshen_4newpv), V6075)


let__17733 := tmp17844
_ = let__17733

tmp17845 := Call0(__e, SymFn(symshen_4incinfs))


_ = tmp17845

tmp17846 := Call2(__e, SymFn(symshen_4deref), V6072, V6075)


tmp17847 := Call3(__e, SymFn(symshen_4app), tmp17846, MakeString(" may create errors\n"), symshen_4a)


tmp17848 := PrimGuard2(__e, symcn, PrimStringConcat, MakeString("warning: changing the type of "), tmp17847)

tmp17849 := Call0(__e, SymFn(symstoutput))


tmp17850 := Call2(__e, SymFn(sympr), tmp17848, tmp17849)


tmp17851 := Call6(__e, SymFn(symis), let__17733, tmp17850, V6075, V6076, V6077, V6078)


return TailCall2(__e, SymFn(symshen_4gc), V6075, tmp17851)


} else {
return False
}


} else {
return let__17732
}


panic("unreachable")
}())
}, 7)

tmp17856 := Call2(__e, ns2_1set, symshen_4variants_2, tmp17839)


_ = tmp17856

tmp17857 := MakeNative1(func(__e *ControlFlow, V6081 Obj) Obj {
tmp17858 := Call1(__e, SymFn(symgensym), symB)


let__17734 := tmp17858
_ = let__17734

tmp17859 := Call1(__e, SymFn(symgensym), symL)


let__17735 := tmp17859
_ = let__17735

tmp17860 := Call1(__e, SymFn(symgensym), symKey)


let__17736 := tmp17860
_ = let__17736

tmp17861 := Call1(__e, SymFn(symgensym), symC)


let__17737 := tmp17861
_ = let__17737

tmp17862 := Call1(__e, SymFn(symgensym), symV)


let__17738 := tmp17862
_ = let__17738

tmp17863 := Call1(__e, SymFn(symshen_4extract_1vars), V6081)


let__17739 := tmp17863
_ = let__17739

tmp17864 := Call1(__e, SymFn(symshen_4rcons__form), V6081)


tmp17865 := GCons(__e, let__17737, Nil)

tmp17866 := GCons(__e, let__17736, tmp17865)

tmp17867 := GCons(__e, let__17735, tmp17866)

tmp17868 := GCons(__e, let__17734, tmp17867)

tmp17869 := GCons(__e, tmp17864, tmp17868)

tmp17870 := GCons(__e, let__17738, tmp17869)

tmp17871 := GCons(__e, symis_b, tmp17870)

tmp17872 := Call3(__e, SymFn(symshen_4stpart), let__17739, tmp17871, let__17734)


tmp17873 := GCons(__e, tmp17872, Nil)

tmp17874 := GCons(__e, let__17737, tmp17873)

tmp17875 := GCons(__e, symlambda, tmp17874)

tmp17876 := GCons(__e, tmp17875, Nil)

tmp17877 := GCons(__e, let__17736, tmp17876)

tmp17878 := GCons(__e, symlambda, tmp17877)

tmp17879 := GCons(__e, tmp17878, Nil)

tmp17880 := GCons(__e, let__17735, tmp17879)

tmp17881 := GCons(__e, symlambda, tmp17880)

tmp17882 := GCons(__e, tmp17881, Nil)

tmp17883 := GCons(__e, let__17734, tmp17882)

tmp17884 := GCons(__e, symlambda, tmp17883)

tmp17885 := GCons(__e, tmp17884, Nil)

tmp17886 := GCons(__e, let__17738, tmp17885)

return GCons(__e, symlambda, tmp17886)


panic("unreachable")
})

tmp17887 := Call2(__e, ns2_1set, symshen_4prolog_1abstraction, tmp17857)


_ = tmp17887

tmp17888 := MakeNative1(func(__e *ControlFlow, V6088 Obj) Obj {
tmp17889 := GValue(__e, symshen_4_ddemodulation_1function_d)

let__17740 := tmp17889
_ = let__17740

return TailCall1(__e, let__17740, V6088)


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symshen_4demod, tmp17888)




panic("unreachable")
})

