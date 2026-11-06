package main

import . "github.com/pyrex41/shen-go/kl"

var MacrosMain = MakeNative0(func(__e *ControlFlow) Obj {
tmp8622 := MakeNative1(func(__e *ControlFlow, V7129 Obj) Obj {
tmp8623 := MakeNative1(func(__e *ControlFlow, Z7131 Obj) Obj {
return GTl(__e, Z7131)
panic("unreachable")
})

tmp8624 := GValue(__e, sym_dmacros_d)

tmp8625 := Call2(__e, SymFn(symmap), tmp8623, tmp8624)


let__8531 := tmp8625
_ = let__8531

tmp8626 := GValue(__e, symshen_4_dsource_1form_1handlers_d)

return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7129, let__8531, let__8531, tmp8626)


panic("unreachable")
})

tmp8627 := Call2(__e, ns2_1set, symmacroexpand, tmp8622)


_ = tmp8627

tmp8628 := (func() Obj {
var __self Obj
__self = MakeNative4(func(__e *ControlFlow, __self1, __self2, __self3, __self4 Obj) Obj {
__selftop:
V7144 := __self1
_ = V7144
V7145 := __self2
_ = V7145
V7146 := __self3
_ = V7146
V7147 := __self4
_ = V7147
tmp8637 := GEq(__e, Nil, V7145)

if True == tmp8637 {
return V7144
} else {
tmp8635 := GConsP(__e, V7145)

if True == tmp8635 {
tmp8629 := GHd(__e, V7145)

tmp8630 := Call3(__e, SymFn(symshen_4source_1macro_1walk), tmp8629, V7144, V7147)


let__8532 := tmp8630
_ = let__8532

tmp8633 := GEq(__e, V7144, let__8532)

if True == tmp8633 {
tmp8631 := GTl(__e, V7145)

if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = V7144, tmp8631, V7146, V7147
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), V7144, tmp8631, V7146, V7147)


} else {
if SymFn(symshen_4macroexpand_1h) == __self {
__self1, __self2, __self3, __self4 = let__8532, V7146, V7146, V7147
__e.Tick()
goto __selftop
}
return TailCall4(__e, SymFn(symshen_4macroexpand_1h), let__8532, V7146, V7146, V7147)
}


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.macroexpand-h"))
}


}


panic("unreachable")
})
return __self
})()

tmp8638 := Call2(__e, ns2_1set, symshen_4macroexpand_1h, tmp8628)


_ = tmp8638

tmp8639 := MakeNative3(func(__e *ControlFlow, V7152, V7153, V7154 Obj) Obj {
tmp8648 := GEq(__e, Nil, V7154)

if True == tmp8648 {
return TailCall2(__e, SymFn(symshen_4walk), V7152, V7153)
} else {
tmp8645 := Call2(__e, SymFn(symshen_4source_1form_1handler), V7153, V7154)


tmp8646 := GConsP(__e, tmp8645)

if True == tmp8646 {
return V7153
} else {
tmp8643 := GConsP(__e, V7153)

if True == tmp8643 {
tmp8640 := MakeNative1(func(__e *ControlFlow, Z7155 Obj) Obj {
return TailCall3(__e, SymFn(symshen_4source_1macro_1walk), V7152, Z7155, V7154)
panic("unreachable")
})

tmp8641 := Call2(__e, SymFn(symmap), tmp8640, V7153)


return TailCall1(__e, V7152, tmp8641)


} else {
return TailCall1(__e, V7152, V7153)
}


}


}


panic("unreachable")
})

tmp8649 := Call2(__e, ns2_1set, symshen_4source_1macro_1walk, tmp8639)


_ = tmp8649

tmp8650 := MakeNative2(func(__e *ControlFlow, V7156, V7157 Obj) Obj {
tmp8654 := GConsP(__e, V7157)

if True == tmp8654 {
tmp8651 := MakeNative1(func(__e *ControlFlow, Z7158 Obj) Obj {
return TailCall2(__e, SymFn(symshen_4walk), V7156, Z7158)
panic("unreachable")
})

tmp8652 := Call2(__e, SymFn(symmap), tmp8651, V7157)


return TailCall1(__e, V7156, tmp8652)


} else {
return TailCall1(__e, V7156, V7157)
}


panic("unreachable")
})

tmp8655 := Call2(__e, ns2_1set, symshen_4walk, tmp8650)


_ = tmp8655

tmp8656 := MakeNative1(func(__e *ControlFlow, V7159 Obj) Obj {
tmp8657 := MakeNative0(func(__e *ControlFlow) Obj {
return V7159
panic("unreachable")
})

let__8533 := tmp8657
_ = let__8533

tmp8943 := GConsP(__e, V7159)

if True == tmp8943 {
tmp8658 := GHd(__e, V7159)

let__8534 := tmp8658
_ = let__8534

tmp8659 := GTl(__e, V7159)

let__8535 := tmp8659
_ = let__8535

tmp8941 := GEq(__e, symdefmacro, let__8534)

var ifres8938 Obj

if True == tmp8941 {
tmp8940 := GConsP(__e, let__8535)

var ifres8939 Obj

if True == tmp8940 {
ifres8939 = True


} else {
ifres8939 = False


}

ifres8938 = ifres8939


} else {
ifres8938 = False


}

if True == ifres8938 {
tmp8660 := GHd(__e, let__8535)

tmp8661 := GTl(__e, let__8535)

return TailCall2(__e, SymFn(symshen_4process_1def), tmp8660, tmp8661)


} else {
tmp8936 := GEq(__e, symdefcc, let__8534)

if True == tmp8936 {
return TailCall1(__e, SymFn(symshen_4yacc_1_6shen), let__8535)
} else {
tmp8934 := GEq(__e, symu_b, let__8534)

var ifres8927 Obj

if True == tmp8934 {
tmp8933 := GConsP(__e, let__8535)

var ifres8929 Obj

if True == tmp8933 {
tmp8931 := GTl(__e, let__8535)

tmp8932 := GEq(__e, Nil, tmp8931)

var ifres8930 Obj

if True == tmp8932 {
ifres8930 = True


} else {
ifres8930 = False


}

ifres8929 = ifres8930


} else {
ifres8929 = False


}

var ifres8928 Obj

if True == ifres8929 {
ifres8928 = True


} else {
ifres8928 = False


}

ifres8927 = ifres8928


} else {
ifres8927 = False


}

if True == ifres8927 {
tmp8662 := GHd(__e, let__8535)

tmp8663 := Call1(__e, SymFn(symshen_4make_1uppercase), tmp8662)


tmp8664 := GCons(__e, tmp8663, Nil)

return GCons(__e, symprotect, tmp8664)


} else {
tmp8925 := GEq(__e, symerror, let__8534)

var ifres8922 Obj

if True == tmp8925 {
tmp8924 := GConsP(__e, let__8535)

var ifres8923 Obj

if True == tmp8924 {
ifres8923 = True


} else {
ifres8923 = False


}

ifres8922 = ifres8923


} else {
ifres8922 = False


}

if True == ifres8922 {
tmp8665 := GHd(__e, let__8535)

tmp8666 := GTl(__e, let__8535)

tmp8667 := Call2(__e, SymFn(symshen_4mkstr), tmp8665, tmp8666)


tmp8668 := GCons(__e, tmp8667, Nil)

return GCons(__e, symsimple_1error, tmp8668)


} else {
tmp8920 := GEq(__e, symoutput, let__8534)

var ifres8917 Obj

if True == tmp8920 {
tmp8919 := GConsP(__e, let__8535)

var ifres8918 Obj

if True == tmp8919 {
ifres8918 = True


} else {
ifres8918 = False


}

ifres8917 = ifres8918


} else {
ifres8917 = False


}

if True == ifres8917 {
tmp8669 := GHd(__e, let__8535)

tmp8670 := GTl(__e, let__8535)

tmp8671 := Call2(__e, SymFn(symshen_4mkstr), tmp8669, tmp8670)


tmp8672 := GCons(__e, symstoutput, Nil)

tmp8673 := GCons(__e, tmp8672, Nil)

tmp8674 := GCons(__e, tmp8671, tmp8673)

return GCons(__e, sympr, tmp8674)


} else {
tmp8915 := GEq(__e, sympr, let__8534)

var ifres8908 Obj

if True == tmp8915 {
tmp8914 := GConsP(__e, let__8535)

var ifres8910 Obj

if True == tmp8914 {
tmp8912 := GTl(__e, let__8535)

tmp8913 := GEq(__e, Nil, tmp8912)

var ifres8911 Obj

if True == tmp8913 {
ifres8911 = True


} else {
ifres8911 = False


}

ifres8910 = ifres8911


} else {
ifres8910 = False


}

var ifres8909 Obj

if True == ifres8910 {
ifres8909 = True


} else {
ifres8909 = False


}

ifres8908 = ifres8909


} else {
ifres8908 = False


}

if True == ifres8908 {
tmp8675 := GHd(__e, let__8535)

tmp8676 := GCons(__e, symstoutput, Nil)

tmp8677 := GCons(__e, tmp8676, Nil)

tmp8678 := GCons(__e, tmp8675, tmp8677)

return GCons(__e, sympr, tmp8678)


} else {
tmp8906 := GEq(__e, symmake_1string, let__8534)

var ifres8903 Obj

if True == tmp8906 {
tmp8905 := GConsP(__e, let__8535)

var ifres8904 Obj

if True == tmp8905 {
ifres8904 = True


} else {
ifres8904 = False


}

ifres8903 = ifres8904


} else {
ifres8903 = False


}

if True == ifres8903 {
tmp8679 := GHd(__e, let__8535)

tmp8680 := GTl(__e, let__8535)

return TailCall2(__e, SymFn(symshen_4mkstr), tmp8679, tmp8680)


} else {
tmp8901 := GEq(__e, symlineread, let__8534)

var ifres8898 Obj

if True == tmp8901 {
tmp8900 := GEq(__e, Nil, let__8535)

var ifres8899 Obj

if True == tmp8900 {
ifres8899 = True


} else {
ifres8899 = False


}

ifres8898 = ifres8899


} else {
ifres8898 = False


}

if True == ifres8898 {
tmp8681 := GCons(__e, symstinput, Nil)

tmp8682 := GCons(__e, tmp8681, Nil)

return GCons(__e, symlineread, tmp8682)


} else {
tmp8896 := GEq(__e, syminput, let__8534)

var ifres8893 Obj

if True == tmp8896 {
tmp8895 := GEq(__e, Nil, let__8535)

var ifres8894 Obj

if True == tmp8895 {
ifres8894 = True


} else {
ifres8894 = False


}

ifres8893 = ifres8894


} else {
ifres8893 = False


}

if True == ifres8893 {
tmp8683 := GCons(__e, symstinput, Nil)

tmp8684 := GCons(__e, tmp8683, Nil)

return GCons(__e, syminput, tmp8684)


} else {
tmp8891 := GEq(__e, symread, let__8534)

var ifres8888 Obj

if True == tmp8891 {
tmp8890 := GEq(__e, Nil, let__8535)

var ifres8889 Obj

if True == tmp8890 {
ifres8889 = True


} else {
ifres8889 = False


}

ifres8888 = ifres8889


} else {
ifres8888 = False


}

if True == ifres8888 {
tmp8685 := GCons(__e, symstinput, Nil)

tmp8686 := GCons(__e, tmp8685, Nil)

return GCons(__e, symread, tmp8686)


} else {
tmp8886 := GEq(__e, syminput_7, let__8534)

var ifres8883 Obj

if True == tmp8886 {
tmp8885 := GConsP(__e, let__8535)

var ifres8884 Obj

if True == tmp8885 {
ifres8884 = True


} else {
ifres8884 = False


}

ifres8883 = ifres8884


} else {
ifres8883 = False


}

if True == ifres8883 {
return TailCall1(__e, SymFn(symshen_4process_1input_7), V7159)
} else {
tmp8881 := GEq(__e, symread_1byte, let__8534)

var ifres8878 Obj

if True == tmp8881 {
tmp8880 := GEq(__e, Nil, let__8535)

var ifres8879 Obj

if True == tmp8880 {
ifres8879 = True


} else {
ifres8879 = False


}

ifres8878 = ifres8879


} else {
ifres8878 = False


}

if True == ifres8878 {
return TailCall0(__e, SymFn(symshen_4process_1read_1byte))
} else {
tmp8876 := GEq(__e, symprolog_2, let__8534)

if True == tmp8876 {
return TailCall1(__e, SymFn(symshen_4call_1prolog), let__8535)
} else {
tmp8874 := GEq(__e, symdefprolog, let__8534)

var ifres8871 Obj

if True == tmp8874 {
tmp8873 := GConsP(__e, let__8535)

var ifres8872 Obj

if True == tmp8873 {
ifres8872 = True


} else {
ifres8872 = False


}

ifres8871 = ifres8872


} else {
ifres8871 = False


}

if True == ifres8871 {
tmp8687 := GHd(__e, let__8535)

tmp8688 := GTl(__e, let__8535)

return TailCall2(__e, SymFn(symshen_4compile_1prolog), tmp8687, tmp8688)


} else {
tmp8869 := GEq(__e, symdatatype, let__8534)

var ifres8866 Obj

if True == tmp8869 {
tmp8868 := GConsP(__e, let__8535)

var ifres8867 Obj

if True == tmp8868 {
ifres8867 = True


} else {
ifres8867 = False


}

ifres8866 = ifres8867


} else {
ifres8866 = False


}

if True == ifres8866 {
tmp8689 := GHd(__e, let__8535)

tmp8690 := GTl(__e, let__8535)

return TailCall2(__e, SymFn(symshen_4process_1datatype), tmp8689, tmp8690)


} else {
tmp8864 := GEq(__e, sym_8s, let__8534)

if True == tmp8864 {
return TailCall1(__e, SymFn(symshen_4process_1_8s), V7159)
} else {
tmp8862 := GEq(__e, symsynonyms, let__8534)

if True == tmp8862 {
return TailCall1(__e, SymFn(symshen_4process_1synonyms), let__8535)
} else {
tmp8860 := GEq(__e, symnl, let__8534)

var ifres8857 Obj

if True == tmp8860 {
tmp8859 := GEq(__e, Nil, let__8535)

var ifres8858 Obj

if True == tmp8859 {
ifres8858 = True


} else {
ifres8858 = False


}

ifres8857 = ifres8858


} else {
ifres8857 = False


}

if True == ifres8857 {
tmp8691 := GCons(__e, MakeInteger(1), Nil)

return GCons(__e, symnl, tmp8691)


} else {
tmp8855 := GEq(__e, symlet, let__8534)

if True == tmp8855 {
return TailCall1(__e, SymFn(symshen_4process_1let), V7159)
} else {
tmp8853 := GEq(__e, sym_c_4, let__8534)

if True == tmp8853 {
return TailCall1(__e, SymFn(symshen_4process_1lambda), V7159)
} else {
tmp8851 := GEq(__e, symcases, let__8534)

if True == tmp8851 {
return TailCall1(__e, SymFn(symshen_4process_1cases), V7159)
} else {
tmp8849 := GEq(__e, symtime, let__8534)

var ifres8842 Obj

if True == tmp8849 {
tmp8848 := GConsP(__e, let__8535)

var ifres8844 Obj

if True == tmp8848 {
tmp8846 := GTl(__e, let__8535)

tmp8847 := GEq(__e, Nil, tmp8846)

var ifres8845 Obj

if True == tmp8847 {
ifres8845 = True


} else {
ifres8845 = False


}

ifres8844 = ifres8845


} else {
ifres8844 = False


}

var ifres8843 Obj

if True == ifres8844 {
ifres8843 = True


} else {
ifres8843 = False


}

ifres8842 = ifres8843


} else {
ifres8842 = False


}

if True == ifres8842 {
tmp8692 := GHd(__e, let__8535)

return TailCall1(__e, SymFn(symshen_4process_1time), tmp8692)


} else {
tmp8840 := GEq(__e, symput, let__8534)

var ifres8825 Obj

if True == tmp8840 {
tmp8839 := GConsP(__e, let__8535)

var ifres8827 Obj

if True == tmp8839 {
tmp8829 := GTl(__e, let__8535)

let__8536 := tmp8829
_ = let__8536

tmp8838 := GConsP(__e, let__8536)

var ifres8830 Obj

if True == tmp8838 {
tmp8832 := GTl(__e, let__8536)

let__8537 := tmp8832
_ = let__8537

tmp8837 := GConsP(__e, let__8537)

var ifres8833 Obj

if True == tmp8837 {
tmp8835 := GTl(__e, let__8537)

tmp8836 := GEq(__e, Nil, tmp8835)

var ifres8834 Obj

if True == tmp8836 {
ifres8834 = True


} else {
ifres8834 = False


}

ifres8833 = ifres8834


} else {
ifres8833 = False


}

var ifres8831 Obj

if True == ifres8833 {
ifres8831 = True


} else {
ifres8831 = False


}

ifres8830 = ifres8831


} else {
ifres8830 = False


}

var ifres8828 Obj

if True == ifres8830 {
ifres8828 = True


} else {
ifres8828 = False


}

ifres8827 = ifres8828


} else {
ifres8827 = False


}

var ifres8826 Obj

if True == ifres8827 {
ifres8826 = True


} else {
ifres8826 = False


}

ifres8825 = ifres8826


} else {
ifres8825 = False


}

if True == ifres8825 {
tmp8693 := GTl(__e, let__8535)

let__8538 := tmp8693
_ = let__8538

tmp8694 := GHd(__e, let__8535)

tmp8695 := GHd(__e, let__8538)

tmp8696 := GTl(__e, let__8538)

tmp8697 := GHd(__e, tmp8696)

tmp8698 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8699 := GCons(__e, symvalue, tmp8698)

tmp8700 := GCons(__e, tmp8699, Nil)

tmp8701 := GCons(__e, tmp8697, tmp8700)

tmp8702 := GCons(__e, tmp8695, tmp8701)

tmp8703 := GCons(__e, tmp8694, tmp8702)

return GCons(__e, symput, tmp8703)


} else {
tmp8823 := GEq(__e, symget, let__8534)

var ifres8812 Obj

if True == tmp8823 {
tmp8822 := GConsP(__e, let__8535)

var ifres8814 Obj

if True == tmp8822 {
tmp8816 := GTl(__e, let__8535)

let__8539 := tmp8816
_ = let__8539

tmp8821 := GConsP(__e, let__8539)

var ifres8817 Obj

if True == tmp8821 {
tmp8819 := GTl(__e, let__8539)

tmp8820 := GEq(__e, Nil, tmp8819)

var ifres8818 Obj

if True == tmp8820 {
ifres8818 = True


} else {
ifres8818 = False


}

ifres8817 = ifres8818


} else {
ifres8817 = False


}

var ifres8815 Obj

if True == ifres8817 {
ifres8815 = True


} else {
ifres8815 = False


}

ifres8814 = ifres8815


} else {
ifres8814 = False


}

var ifres8813 Obj

if True == ifres8814 {
ifres8813 = True


} else {
ifres8813 = False


}

ifres8812 = ifres8813


} else {
ifres8812 = False


}

if True == ifres8812 {
tmp8704 := GHd(__e, let__8535)

tmp8705 := GTl(__e, let__8535)

tmp8706 := GHd(__e, tmp8705)

tmp8707 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8708 := GCons(__e, symvalue, tmp8707)

tmp8709 := GCons(__e, tmp8708, Nil)

tmp8710 := GCons(__e, tmp8706, tmp8709)

tmp8711 := GCons(__e, tmp8704, tmp8710)

return GCons(__e, symget, tmp8711)


} else {
tmp8810 := GEq(__e, symunput, let__8534)

var ifres8799 Obj

if True == tmp8810 {
tmp8809 := GConsP(__e, let__8535)

var ifres8801 Obj

if True == tmp8809 {
tmp8803 := GTl(__e, let__8535)

let__8540 := tmp8803
_ = let__8540

tmp8808 := GConsP(__e, let__8540)

var ifres8804 Obj

if True == tmp8808 {
tmp8806 := GTl(__e, let__8540)

tmp8807 := GEq(__e, Nil, tmp8806)

var ifres8805 Obj

if True == tmp8807 {
ifres8805 = True


} else {
ifres8805 = False


}

ifres8804 = ifres8805


} else {
ifres8804 = False


}

var ifres8802 Obj

if True == ifres8804 {
ifres8802 = True


} else {
ifres8802 = False


}

ifres8801 = ifres8802


} else {
ifres8801 = False


}

var ifres8800 Obj

if True == ifres8801 {
ifres8800 = True


} else {
ifres8800 = False


}

ifres8799 = ifres8800


} else {
ifres8799 = False


}

if True == ifres8799 {
tmp8712 := GHd(__e, let__8535)

tmp8713 := GTl(__e, let__8535)

tmp8714 := GHd(__e, tmp8713)

tmp8715 := GCons(__e, sym_dproperty_1vector_d, Nil)

tmp8716 := GCons(__e, symvalue, tmp8715)

tmp8717 := GCons(__e, tmp8716, Nil)

tmp8718 := GCons(__e, tmp8714, tmp8717)

tmp8719 := GCons(__e, tmp8712, tmp8718)

return GCons(__e, symunput, tmp8719)


} else {
tmp8797 := GEq(__e, symshen_4_8c, let__8534)

var ifres8790 Obj

if True == tmp8797 {
tmp8796 := GConsP(__e, let__8535)

var ifres8792 Obj

if True == tmp8796 {
tmp8794 := GTl(__e, let__8535)

tmp8795 := GEq(__e, Nil, tmp8794)

var ifres8793 Obj

if True == tmp8795 {
ifres8793 = True


} else {
ifres8793 = False


}

ifres8792 = ifres8793


} else {
ifres8792 = False


}

var ifres8791 Obj

if True == ifres8792 {
ifres8791 = True


} else {
ifres8791 = False


}

ifres8790 = ifres8791


} else {
ifres8790 = False


}

if True == ifres8790 {
tmp8720 := GHd(__e, let__8535)

return TailCall1(__e, SymFn(symshen_4rcons__form), tmp8720)


} else {
tmp8721 := MakeNative0(func(__e *ControlFlow) Obj {
tmp8748 := GConsP(__e, let__8535)

var ifres8729 Obj

if True == tmp8748 {
tmp8731 := GTl(__e, let__8535)

let__8542 := tmp8731
_ = let__8542

tmp8747 := GConsP(__e, let__8542)

var ifres8732 Obj

if True == tmp8747 {
tmp8745 := GTl(__e, let__8542)

tmp8746 := GConsP(__e, tmp8745)

var ifres8734 Obj

if True == tmp8746 {
tmp8736 := GCons(__e, symdo, Nil)

tmp8737 := GCons(__e, sym_d, tmp8736)

tmp8738 := GCons(__e, sym_7, tmp8737)

tmp8739 := GCons(__e, symor, tmp8738)

tmp8740 := GCons(__e, symand, tmp8739)

tmp8741 := GCons(__e, symappend, tmp8740)

tmp8742 := GCons(__e, sym_8v, tmp8741)

tmp8743 := GCons(__e, sym_8p, tmp8742)

tmp8744 := Call2(__e, SymFn(symelement_2), let__8534, tmp8743)


var ifres8735 Obj

if True == tmp8744 {
ifres8735 = True


} else {
ifres8735 = False


}

ifres8734 = ifres8735


} else {
ifres8734 = False


}

var ifres8733 Obj

if True == ifres8734 {
ifres8733 = True


} else {
ifres8733 = False


}

ifres8732 = ifres8733


} else {
ifres8732 = False


}

var ifres8730 Obj

if True == ifres8732 {
ifres8730 = True


} else {
ifres8730 = False


}

ifres8729 = ifres8730


} else {
ifres8729 = False


}

if True == ifres8729 {
tmp8722 := GHd(__e, let__8535)

tmp8723 := GTl(__e, let__8535)

tmp8724 := GCons(__e, let__8534, tmp8723)

tmp8725 := Call1(__e, SymFn(symshen_4process_1assoc), tmp8724)


tmp8726 := GCons(__e, tmp8725, Nil)

tmp8727 := GCons(__e, tmp8722, tmp8726)

return GCons(__e, let__8534, tmp8727)


} else {
return TailCall1(__e, SymFn(symthaw), let__8533)
}


panic("unreachable")
})

let__8541 := tmp8721
_ = let__8541

tmp8788 := GEq(__e, symshen_4_8ch, let__8534)

if True == tmp8788 {
tmp8786 := GConsP(__e, let__8535)

if True == tmp8786 {
tmp8749 := GHd(__e, let__8535)

let__8543 := tmp8749
_ = let__8543

tmp8750 := GTl(__e, let__8535)

let__8544 := tmp8750
_ = let__8544

tmp8784 := GConsP(__e, let__8543)

var ifres8764 Obj

if True == tmp8784 {
tmp8766 := GTl(__e, let__8543)

let__8545 := tmp8766
_ = let__8545

tmp8783 := GConsP(__e, let__8545)

var ifres8767 Obj

if True == tmp8783 {
tmp8769 := GTl(__e, let__8545)

let__8546 := tmp8769
_ = let__8546

tmp8782 := GConsP(__e, let__8546)

var ifres8770 Obj

if True == tmp8782 {
tmp8780 := GTl(__e, let__8546)

tmp8781 := GEq(__e, Nil, tmp8780)

var ifres8772 Obj

if True == tmp8781 {
tmp8779 := GEq(__e, Nil, let__8544)

var ifres8774 Obj

if True == tmp8779 {
tmp8776 := GHd(__e, let__8545)

tmp8777 := PrimGuard1(__e, symintern, PrimIntern, MakeString(":"))

tmp8778 := GEq(__e, tmp8776, tmp8777)

var ifres8775 Obj

if True == tmp8778 {
ifres8775 = True


} else {
ifres8775 = False


}

ifres8774 = ifres8775


} else {
ifres8774 = False


}

var ifres8773 Obj

if True == ifres8774 {
ifres8773 = True


} else {
ifres8773 = False


}

ifres8772 = ifres8773


} else {
ifres8772 = False


}

var ifres8771 Obj

if True == ifres8772 {
ifres8771 = True


} else {
ifres8771 = False


}

ifres8770 = ifres8771


} else {
ifres8770 = False


}

var ifres8768 Obj

if True == ifres8770 {
ifres8768 = True


} else {
ifres8768 = False


}

ifres8767 = ifres8768


} else {
ifres8767 = False


}

var ifres8765 Obj

if True == ifres8767 {
ifres8765 = True


} else {
ifres8765 = False


}

ifres8764 = ifres8765


} else {
ifres8764 = False


}

if True == ifres8764 {
tmp8751 := GTl(__e, let__8543)

let__8547 := tmp8751
_ = let__8547

tmp8752 := GHd(__e, let__8543)

tmp8753 := GHd(__e, let__8547)

tmp8754 := GTl(__e, let__8547)

tmp8755 := GCons(__e, sym_7, tmp8754)

tmp8756 := GCons(__e, tmp8755, Nil)

tmp8757 := GCons(__e, tmp8753, tmp8756)

tmp8758 := GCons(__e, tmp8752, tmp8757)

tmp8759 := GCons(__e, tmp8758, Nil)

tmp8760 := GCons(__e, sym_1, tmp8759)

return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8760)


} else {
tmp8762 := GEq(__e, Nil, let__8544)

if True == tmp8762 {
return TailCall1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8543)
} else {
return TailCall1(__e, SymFn(symthaw), let__8541)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8541)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8541)
}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8533)
}


panic("unreachable")
})

tmp8944 := Call2(__e, ns2_1set, symshen_4macros, tmp8656)


_ = tmp8944

tmp8945 := MakeNative1(func(__e *ControlFlow, V7176 Obj) Obj {
tmp8946 := MakeNative0(func(__e *ControlFlow) Obj {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4process_1input_7)
panic("unreachable")
})

let__8548 := tmp8946
_ = let__8548

tmp8970 := GConsP(__e, V7176)

if True == tmp8970 {
tmp8947 := GTl(__e, V7176)

let__8549 := tmp8947
_ = let__8549

tmp8967 := GHd(__e, V7176)

tmp8968 := GEq(__e, syminput_7, tmp8967)

if True == tmp8968 {
tmp8965 := GConsP(__e, let__8549)

if True == tmp8965 {
tmp8948 := GHd(__e, let__8549)

let__8550 := tmp8948
_ = let__8550

tmp8949 := GTl(__e, let__8549)

let__8551 := tmp8949
_ = let__8551

tmp8963 := GEq(__e, Nil, let__8551)

if True == tmp8963 {
tmp8950 := Call1(__e, SymFn(symshen_4rcons__form), let__8550)


tmp8951 := GCons(__e, symstinput, Nil)

tmp8952 := GCons(__e, tmp8951, Nil)

tmp8953 := GCons(__e, tmp8950, tmp8952)

return GCons(__e, symshen_4input_1h_7, tmp8953)


} else {
tmp8961 := GConsP(__e, let__8551)

var ifres8957 Obj

if True == tmp8961 {
tmp8959 := GTl(__e, let__8551)

tmp8960 := GEq(__e, Nil, tmp8959)

var ifres8958 Obj

if True == tmp8960 {
ifres8958 = True


} else {
ifres8958 = False


}

ifres8957 = ifres8958


} else {
ifres8957 = False


}

if True == ifres8957 {
tmp8954 := Call1(__e, SymFn(symshen_4rcons__form), let__8550)


tmp8955 := GCons(__e, tmp8954, let__8551)

return GCons(__e, symshen_4input_1h_7, tmp8955)


} else {
return TailCall1(__e, SymFn(symthaw), let__8548)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8548)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8548)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8548)
}


panic("unreachable")
})

tmp8971 := Call2(__e, ns2_1set, symshen_4process_1input_7, tmp8945)


_ = tmp8971

tmp8972 := MakeNative1(func(__e *ControlFlow, V7183 Obj) Obj {
tmp8973 := MakeNative0(func(__e *ControlFlow) Obj {
return V7183
panic("unreachable")
})

let__8552 := tmp8973
_ = let__8552

tmp9014 := GConsP(__e, V7183)

if True == tmp9014 {
tmp8974 := GHd(__e, V7183)

let__8553 := tmp8974
_ = let__8553

tmp8975 := GTl(__e, V7183)

let__8554 := tmp8975
_ = let__8554

tmp9012 := GEq(__e, sym_7, let__8553)

var ifres9005 Obj

if True == tmp9012 {
tmp9011 := GConsP(__e, let__8554)

var ifres9007 Obj

if True == tmp9011 {
tmp9009 := GTl(__e, let__8554)

tmp9010 := GEq(__e, Nil, tmp9009)

var ifres9008 Obj

if True == tmp9010 {
ifres9008 = True


} else {
ifres9008 = False


}

ifres9007 = ifres9008


} else {
ifres9007 = False


}

var ifres9006 Obj

if True == ifres9007 {
ifres9006 = True


} else {
ifres9006 = False


}

ifres9005 = ifres9006


} else {
ifres9005 = False


}

if True == ifres9005 {
tmp8976 := GHd(__e, let__8554)

tmp8977 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8976)


tmp8978 := GCons(__e, tmp8977, Nil)

return GCons(__e, sym_7, tmp8978)


} else {
tmp9003 := GEq(__e, sym_1, let__8553)

var ifres8996 Obj

if True == tmp9003 {
tmp9002 := GConsP(__e, let__8554)

var ifres8998 Obj

if True == tmp9002 {
tmp9000 := GTl(__e, let__8554)

tmp9001 := GEq(__e, Nil, tmp9000)

var ifres8999 Obj

if True == tmp9001 {
ifres8999 = True


} else {
ifres8999 = False


}

ifres8998 = ifres8999


} else {
ifres8998 = False


}

var ifres8997 Obj

if True == ifres8998 {
ifres8997 = True


} else {
ifres8997 = False


}

ifres8996 = ifres8997


} else {
ifres8996 = False


}

if True == ifres8996 {
tmp8979 := GHd(__e, let__8554)

tmp8980 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), tmp8979)


tmp8981 := GCons(__e, tmp8980, Nil)

return GCons(__e, sym_1, tmp8981)


} else {
tmp8994 := GEq(__e, symbar_b, let__8553)

var ifres8987 Obj

if True == tmp8994 {
tmp8993 := GConsP(__e, let__8554)

var ifres8989 Obj

if True == tmp8993 {
tmp8991 := GTl(__e, let__8554)

tmp8992 := GEq(__e, Nil, tmp8991)

var ifres8990 Obj

if True == tmp8992 {
ifres8990 = True


} else {
ifres8990 = False


}

ifres8989 = ifres8990


} else {
ifres8989 = False


}

var ifres8988 Obj

if True == ifres8989 {
ifres8988 = True


} else {
ifres8988 = False


}

ifres8987 = ifres8988


} else {
ifres8987 = False


}

if True == ifres8987 {
return GHd(__e, let__8554)
} else {
tmp8982 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8553)


tmp8983 := Call1(__e, SymFn(symshen_4cons_1form_1respect_1modes), let__8554)


tmp8984 := GCons(__e, tmp8983, Nil)

tmp8985 := GCons(__e, tmp8982, tmp8984)

return GCons(__e, symcons, tmp8985)


}


}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8552)
}


panic("unreachable")
})

tmp9015 := Call2(__e, ns2_1set, symshen_4cons_1form_1respect_1modes, tmp8972)


_ = tmp9015

tmp9016 := MakeNative2(func(__e *ControlFlow, V7187, V7188 Obj) Obj {
tmp9017 := GCons(__e, symX, Nil)

tmp9018 := GCons(__e, sym_1_6, tmp9017)

tmp9019 := GCons(__e, symX, tmp9018)

let__8555 := tmp9019
_ = let__8555

tmp9020 := Call2(__e, SymFn(symappend), V7188, let__8555)


tmp9021 := GCons(__e, V7187, tmp9020)

tmp9022 := GCons(__e, symdefine, tmp9021)

tmp9023 := Call1(__e, SymFn(symeval), tmp9022)


let__8556 := tmp9023
_ = let__8556

tmp9024 := Call1(__e, SymFn(symfn), V7187)


tmp9025 := Call2(__e, SymFn(symshen_4record_1macro), V7187, tmp9024)


let__8557 := tmp9025
_ = let__8557

return V7187


panic("unreachable")
})

tmp9026 := Call2(__e, ns2_1set, symshen_4process_1def, tmp9016)


_ = tmp9026

tmp9027 := MakeNative1(func(__e *ControlFlow, V7192 Obj) Obj {
tmp9059 := GConsP(__e, V7192)

var ifres9039 Obj

if True == tmp9059 {
tmp9041 := GTl(__e, V7192)

let__8558 := tmp9041
_ = let__8558

tmp9057 := GHd(__e, V7192)

tmp9058 := GEq(__e, symlet, tmp9057)

var ifres9042 Obj

if True == tmp9058 {
tmp9056 := GConsP(__e, let__8558)

var ifres9044 Obj

if True == tmp9056 {
tmp9046 := GTl(__e, let__8558)

let__8559 := tmp9046
_ = let__8559

tmp9055 := GConsP(__e, let__8559)

var ifres9047 Obj

if True == tmp9055 {
tmp9049 := GTl(__e, let__8559)

let__8560 := tmp9049
_ = let__8560

tmp9054 := GConsP(__e, let__8560)

var ifres9050 Obj

if True == tmp9054 {
tmp9052 := GTl(__e, let__8560)

tmp9053 := GConsP(__e, tmp9052)

var ifres9051 Obj

if True == tmp9053 {
ifres9051 = True


} else {
ifres9051 = False


}

ifres9050 = ifres9051


} else {
ifres9050 = False


}

var ifres9048 Obj

if True == ifres9050 {
ifres9048 = True


} else {
ifres9048 = False


}

ifres9047 = ifres9048


} else {
ifres9047 = False


}

var ifres9045 Obj

if True == ifres9047 {
ifres9045 = True


} else {
ifres9045 = False


}

ifres9044 = ifres9045


} else {
ifres9044 = False


}

var ifres9043 Obj

if True == ifres9044 {
ifres9043 = True


} else {
ifres9043 = False


}

ifres9042 = ifres9043


} else {
ifres9042 = False


}

var ifres9040 Obj

if True == ifres9042 {
ifres9040 = True


} else {
ifres9040 = False


}

ifres9039 = ifres9040


} else {
ifres9039 = False


}

if True == ifres9039 {
tmp9028 := GTl(__e, V7192)

let__8561 := tmp9028
_ = let__8561

tmp9029 := GHd(__e, let__8561)

tmp9030 := GTl(__e, let__8561)

tmp9031 := GHd(__e, tmp9030)

tmp9032 := GTl(__e, let__8561)

tmp9033 := GTl(__e, tmp9032)

tmp9034 := GCons(__e, symlet, tmp9033)

tmp9035 := GCons(__e, tmp9034, Nil)

tmp9036 := GCons(__e, tmp9031, tmp9035)

tmp9037 := GCons(__e, tmp9029, tmp9036)

return GCons(__e, symlet, tmp9037)


} else {
return V7192
}


panic("unreachable")
})

tmp9060 := Call2(__e, ns2_1set, symshen_4process_1let, tmp9027)


_ = tmp9060

tmp9061 := (func() Obj {
var __self Obj
__self = MakeNative1(func(__e *ControlFlow, __self1 Obj) Obj {
__selftop:
V7197 := __self1
_ = V7197
tmp9062 := MakeNative0(func(__e *ControlFlow) Obj {
return V7197
panic("unreachable")
})

let__8562 := tmp9062
_ = let__8562

tmp9092 := GConsP(__e, V7197)

if True == tmp9092 {
tmp9063 := GTl(__e, V7197)

let__8563 := tmp9063
_ = let__8563

tmp9089 := GHd(__e, V7197)

tmp9090 := GEq(__e, sym_8s, tmp9089)

if True == tmp9090 {
tmp9087 := GConsP(__e, let__8563)

if True == tmp9087 {
tmp9064 := GHd(__e, let__8563)

let__8564 := tmp9064
_ = let__8564

tmp9065 := GTl(__e, let__8563)

let__8565 := tmp9065
_ = let__8565

tmp9085 := GConsP(__e, let__8565)

if True == tmp9085 {
tmp9066 := GTl(__e, let__8565)

let__8566 := tmp9066
_ = let__8566

tmp9083 := GConsP(__e, let__8566)

if True == tmp9083 {
tmp9067 := GCons(__e, sym_8s, let__8565)

tmp9068 := Call1(__e, SymFn(symshen_4process_1_8s), tmp9067)


tmp9069 := GCons(__e, tmp9068, Nil)

tmp9070 := GCons(__e, let__8564, tmp9069)

return GCons(__e, sym_8s, tmp9070)


} else {
tmp9081 := GEq(__e, Nil, let__8566)

var ifres9078 Obj

if True == tmp9081 {
tmp9080 := PrimGuard1(__e, symstring_2, PrimIsString, let__8564)

var ifres9079 Obj

if True == tmp9080 {
ifres9079 = True


} else {
ifres9079 = False


}

ifres9078 = ifres9079


} else {
ifres9078 = False


}

if True == ifres9078 {
tmp9071 := Call1(__e, SymFn(symexplode), let__8564)


let__8567 := tmp9071
_ = let__8567

tmp9075 := Call1(__e, SymFn(symlength), let__8567)


if True == GGt(__e, tmp9075, MakeInteger(1)) {
tmp9072 := Call2(__e, SymFn(symappend), let__8567, let__8565)


tmp9073 := GCons(__e, sym_8s, tmp9072)

if SymFn(symshen_4process_1_8s) == __self {
__self1 = tmp9073
__e.Tick()
goto __selftop
}
return TailCall1(__e, SymFn(symshen_4process_1_8s), tmp9073)


} else {
return V7197
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8562)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8562)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8562)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8562)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8562)
}


panic("unreachable")
})
return __self
})()

tmp9093 := Call2(__e, ns2_1set, symshen_4process_1_8s, tmp9061)


_ = tmp9093

tmp9094 := MakeNative2(func(__e *ControlFlow, V7207, V7208 Obj) Obj {
tmp9095 := Call1(__e, SymFn(symshen_4intern_1type), V7207)


let__8568 := tmp9095
_ = let__8568

tmp9096 := MakeNative1(func(__e *ControlFlow, Z7211 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5datatype_6), Z7211)
panic("unreachable")
})

tmp9097 := GCons(__e, let__8568, V7208)

tmp9098 := Call2(__e, SymFn(symcompile), tmp9096, tmp9097)


let__8569 := tmp9098
_ = let__8569

return let__8568


panic("unreachable")
})

tmp9099 := Call2(__e, ns2_1set, symshen_4process_1datatype, tmp9094)


_ = tmp9099

tmp9100 := MakeNative1(func(__e *ControlFlow, V7212 Obj) Obj {
tmp9101 := PrimGuard1(__e, symstr, PrimStr, V7212)

return (func() Obj {
if TypedIREnabled() && HasCanonicalPrimitiveBinding(symintern) {
return PrimIntern(PrimGuard2(__e, symcn, PrimStringConcat, tmp9101, MakeString("#type")))
}
__typedArg0 := PrimGuard2(__e, symcn, PrimStringConcat, tmp9101, MakeString("#type"))
return Call1(__e, SymFn(symintern), __typedArg0)
})()


panic("unreachable")
})

tmp9103 := Call2(__e, ns2_1set, symshen_4intern_1type, tmp9100)


_ = tmp9103

tmp9104 := MakeNative1(func(__e *ControlFlow, V7213 Obj) Obj {
tmp9105 := GValue(__e, symshen_4_dsynonyms_d)

tmp9106 := Call2(__e, SymFn(symappend), V7213, tmp9105)


tmp9107 := PrimGuard2(__e, symset, PrimSet, symshen_4_dsynonyms_d, tmp9106)

return TailCall1(__e, SymFn(symshen_4synonyms_1h), tmp9107)


panic("unreachable")
})

tmp9108 := Call2(__e, ns2_1set, symshen_4process_1synonyms, tmp9104)


_ = tmp9108

tmp9109 := MakeNative1(func(__e *ControlFlow, V7216 Obj) Obj {
tmp9147 := GConsP(__e, V7216)

var ifres9119 Obj

if True == tmp9147 {
tmp9121 := GTl(__e, V7216)

let__8570 := tmp9121
_ = let__8570

tmp9145 := GHd(__e, V7216)

tmp9146 := GEq(__e, symdefun, tmp9145)

var ifres9122 Obj

if True == tmp9146 {
tmp9144 := GConsP(__e, let__8570)

var ifres9124 Obj

if True == tmp9144 {
tmp9126 := GTl(__e, let__8570)

let__8571 := tmp9126
_ = let__8571

tmp9143 := GConsP(__e, let__8571)

var ifres9127 Obj

if True == tmp9143 {
tmp9129 := GHd(__e, let__8571)

let__8572 := tmp9129
_ = let__8572

tmp9130 := GTl(__e, let__8571)

let__8573 := tmp9130
_ = let__8573

tmp9142 := GConsP(__e, let__8572)

var ifres9131 Obj

if True == tmp9142 {
tmp9140 := GTl(__e, let__8572)

tmp9141 := GEq(__e, Nil, tmp9140)

var ifres9133 Obj

if True == tmp9141 {
tmp9139 := GConsP(__e, let__8573)

var ifres9135 Obj

if True == tmp9139 {
tmp9137 := GTl(__e, let__8573)

tmp9138 := GEq(__e, Nil, tmp9137)

var ifres9136 Obj

if True == tmp9138 {
ifres9136 = True


} else {
ifres9136 = False


}

ifres9135 = ifres9136


} else {
ifres9135 = False


}

var ifres9134 Obj

if True == ifres9135 {
ifres9134 = True


} else {
ifres9134 = False


}

ifres9133 = ifres9134


} else {
ifres9133 = False


}

var ifres9132 Obj

if True == ifres9133 {
ifres9132 = True


} else {
ifres9132 = False


}

ifres9131 = ifres9132


} else {
ifres9131 = False


}

var ifres9128 Obj

if True == ifres9131 {
ifres9128 = True


} else {
ifres9128 = False


}

ifres9127 = ifres9128


} else {
ifres9127 = False


}

var ifres9125 Obj

if True == ifres9127 {
ifres9125 = True


} else {
ifres9125 = False


}

ifres9124 = ifres9125


} else {
ifres9124 = False


}

var ifres9123 Obj

if True == ifres9124 {
ifres9123 = True


} else {
ifres9123 = False


}

ifres9122 = ifres9123


} else {
ifres9122 = False


}

var ifres9120 Obj

if True == ifres9122 {
ifres9120 = True


} else {
ifres9120 = False


}

ifres9119 = ifres9120


} else {
ifres9119 = False


}

if True == ifres9119 {
tmp9110 := GTl(__e, V7216)

let__8574 := tmp9110
_ = let__8574

tmp9111 := GTl(__e, let__8574)

tmp9112 := GHd(__e, tmp9111)

tmp9113 := GHd(__e, tmp9112)

tmp9114 := GTl(__e, let__8574)

tmp9115 := GTl(__e, tmp9114)

tmp9116 := GCons(__e, tmp9113, tmp9115)

tmp9117 := GCons(__e, sym_c_4, tmp9116)

return TailCall1(__e, SymFn(symeval), tmp9117)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4lambda_1of_1defun)
}


panic("unreachable")
})

tmp9148 := Call2(__e, ns2_1set, symshen_4lambda_1of_1defun, tmp9109)


_ = tmp9148

tmp9149 := MakeNative1(func(__e *ControlFlow, V7222 Obj) Obj {
tmp9150 := MakeNative1(func(__e *ControlFlow, Z7224 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4curry_1type), Z7224)
panic("unreachable")
})

tmp9151 := Call2(__e, SymFn(symmap), tmp9150, V7222)


let__8575 := tmp9151
_ = let__8575

tmp9152 := Call1(__e, SymFn(symshen_4compile_1synonyms), let__8575)


tmp9153 := Call2(__e, SymFn(symshen_4shendef_1_6kldef), symshen_4demod, tmp9152)


tmp9154 := Call1(__e, SymFn(symshen_4lambda_1of_1defun), tmp9153)


let__8576 := tmp9154
_ = let__8576

tmp9155 := PrimGuard2(__e, symset, PrimSet, symshen_4_ddemodulation_1function_d, let__8576)

let__8577 := tmp9155
_ = let__8577

return symsynonyms


panic("unreachable")
})

tmp9156 := Call2(__e, ns2_1set, symshen_4synonyms_1h, tmp9149)


_ = tmp9156

tmp9157 := MakeNative1(func(__e *ControlFlow, V7229 Obj) Obj {
tmp9177 := GEq(__e, Nil, V7229)

if True == tmp9177 {
tmp9158 := Call1(__e, SymFn(symgensym), symX)


let__8578 := tmp9158
_ = let__8578

tmp9159 := GCons(__e, let__8578, Nil)

tmp9160 := GCons(__e, sym_1_6, tmp9159)

return GCons(__e, let__8578, tmp9160)


} else {
tmp9175 := GConsP(__e, V7229)

var ifres9171 Obj

if True == tmp9175 {
tmp9173 := GTl(__e, V7229)

tmp9174 := GConsP(__e, tmp9173)

var ifres9172 Obj

if True == tmp9174 {
ifres9172 = True


} else {
ifres9172 = False


}

ifres9171 = ifres9172


} else {
ifres9171 = False


}

if True == ifres9171 {
tmp9161 := GTl(__e, V7229)

let__8579 := tmp9161
_ = let__8579

tmp9162 := GHd(__e, V7229)

tmp9163 := Call1(__e, SymFn(symshen_4rcons__form), tmp9162)


tmp9164 := GHd(__e, let__8579)

tmp9165 := Call1(__e, SymFn(symshen_4rcons__form), tmp9164)


tmp9166 := GTl(__e, let__8579)

tmp9167 := Call1(__e, SymFn(symshen_4compile_1synonyms), tmp9166)


tmp9168 := GCons(__e, tmp9165, tmp9167)

tmp9169 := GCons(__e, sym_1_6, tmp9168)

return GCons(__e, tmp9163, tmp9169)


} else {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("synonyms requires an even number of arguments\n"))
}


}


panic("unreachable")
})

tmp9178 := Call2(__e, ns2_1set, symshen_4compile_1synonyms, tmp9157)


_ = tmp9178

tmp9179 := MakeNative1(func(__e *ControlFlow, V7232 Obj) Obj {
tmp9180 := MakeNative0(func(__e *ControlFlow) Obj {
return V7232
panic("unreachable")
})

let__8580 := tmp9180
_ = let__8580

tmp9204 := GConsP(__e, V7232)

if True == tmp9204 {
tmp9181 := GTl(__e, V7232)

let__8581 := tmp9181
_ = let__8581

tmp9201 := GHd(__e, V7232)

tmp9202 := GEq(__e, sym_c_4, tmp9201)

if True == tmp9202 {
tmp9199 := GConsP(__e, let__8581)

if True == tmp9199 {
tmp9182 := GHd(__e, let__8581)

let__8582 := tmp9182
_ = let__8582

tmp9183 := GTl(__e, let__8581)

let__8583 := tmp9183
_ = let__8583

tmp9197 := GConsP(__e, let__8583)

if True == tmp9197 {
tmp9184 := GTl(__e, let__8583)

let__8584 := tmp9184
_ = let__8584

tmp9195 := GConsP(__e, let__8584)

if True == tmp9195 {
tmp9185 := GCons(__e, sym_c_4, let__8583)

tmp9186 := Call1(__e, SymFn(symshen_4process_1lambda), tmp9185)


tmp9187 := GCons(__e, tmp9186, Nil)

tmp9188 := GCons(__e, let__8582, tmp9187)

return GCons(__e, symlambda, tmp9188)


} else {
tmp9193 := GEq(__e, Nil, let__8584)

if True == tmp9193 {
tmp9191 := PrimGuard1(__e, symvariable_2, PrimIsVariable, let__8582)

if True == tmp9191 {
return GCons(__e, symlambda, let__8581)
} else {
tmp9189 := Call3(__e, SymFn(symshen_4app), let__8582, MakeString(" is not a variable\n"), symshen_4s)


return PrimGuard1(__e, symsimple_1error, PrimSimpleError, tmp9189)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8580)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8580)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8580)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8580)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8580)
}


panic("unreachable")
})

tmp9205 := Call2(__e, ns2_1set, symshen_4process_1lambda, tmp9179)


_ = tmp9205

tmp9206 := MakeNative1(func(__e *ControlFlow, V7243 Obj) Obj {
tmp9207 := MakeNative0(func(__e *ControlFlow) Obj {
return V7243
panic("unreachable")
})

let__8585 := tmp9207
_ = let__8585

tmp9241 := GConsP(__e, V7243)

if True == tmp9241 {
tmp9208 := GTl(__e, V7243)

let__8586 := tmp9208
_ = let__8586

tmp9238 := GHd(__e, V7243)

tmp9239 := GEq(__e, symcases, tmp9238)

if True == tmp9239 {
tmp9236 := GConsP(__e, let__8586)

if True == tmp9236 {
tmp9209 := GHd(__e, let__8586)

let__8587 := tmp9209
_ = let__8587

tmp9210 := GTl(__e, let__8586)

let__8588 := tmp9210
_ = let__8588

tmp9234 := GEq(__e, True, let__8587)

var ifres9231 Obj

if True == tmp9234 {
tmp9233 := GConsP(__e, let__8588)

var ifres9232 Obj

if True == tmp9233 {
ifres9232 = True


} else {
ifres9232 = False


}

ifres9231 = ifres9232


} else {
ifres9231 = False


}

if True == ifres9231 {
return GHd(__e, let__8588)
} else {
tmp9211 := MakeNative0(func(__e *ControlFlow) Obj {
tmp9213 := GEq(__e, Nil, let__8588)

if True == tmp9213 {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("error: odd number of case elements\n"))
} else {
return TailCall1(__e, SymFn(symthaw), let__8585)
}


panic("unreachable")
})

let__8589 := tmp9211
_ = let__8589

tmp9229 := GConsP(__e, let__8588)

if True == tmp9229 {
tmp9214 := GHd(__e, let__8588)

let__8590 := tmp9214
_ = let__8590

tmp9215 := GTl(__e, let__8588)

let__8591 := tmp9215
_ = let__8591

tmp9227 := GEq(__e, Nil, let__8591)

if True == tmp9227 {
tmp9216 := GCons(__e, MakeString("error: cases exhausted"), Nil)

tmp9217 := GCons(__e, symsimple_1error, tmp9216)

tmp9218 := GCons(__e, tmp9217, Nil)

tmp9219 := GCons(__e, let__8590, tmp9218)

tmp9220 := GCons(__e, let__8587, tmp9219)

return GCons(__e, symif, tmp9220)


} else {
tmp9221 := GCons(__e, symcases, let__8591)

tmp9222 := Call1(__e, SymFn(symshen_4process_1cases), tmp9221)


tmp9223 := GCons(__e, tmp9222, Nil)

tmp9224 := GCons(__e, let__8590, tmp9223)

tmp9225 := GCons(__e, let__8587, tmp9224)

return GCons(__e, symif, tmp9225)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8589)
}


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8585)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8585)
}


} else {
return TailCall1(__e, SymFn(symthaw), let__8585)
}


panic("unreachable")
})

tmp9242 := Call2(__e, ns2_1set, symshen_4process_1cases, tmp9206)


_ = tmp9242

tmp9243 := MakeNative1(func(__e *ControlFlow, V7253 Obj) Obj {
tmp9244 := GCons(__e, symrun, Nil)

tmp9245 := GCons(__e, symget_1time, tmp9244)

tmp9246 := GCons(__e, symrun, Nil)

tmp9247 := GCons(__e, symget_1time, tmp9246)

tmp9248 := GCons(__e, symStart, Nil)

tmp9249 := GCons(__e, symFinish, tmp9248)

tmp9250 := GCons(__e, sym_1, tmp9249)

tmp9251 := GCons(__e, symTime, Nil)

tmp9252 := GCons(__e, symstr, tmp9251)

tmp9253 := GCons(__e, MakeString(" secs\n"), Nil)

tmp9254 := GCons(__e, tmp9252, tmp9253)

tmp9255 := GCons(__e, symcn, tmp9254)

tmp9256 := GCons(__e, tmp9255, Nil)

tmp9257 := GCons(__e, MakeString("\nrun time: "), tmp9256)

tmp9258 := GCons(__e, symcn, tmp9257)

tmp9259 := GCons(__e, symstoutput, Nil)

tmp9260 := GCons(__e, tmp9259, Nil)

tmp9261 := GCons(__e, tmp9258, tmp9260)

tmp9262 := GCons(__e, sympr, tmp9261)

tmp9263 := GCons(__e, symResult, Nil)

tmp9264 := GCons(__e, tmp9262, tmp9263)

tmp9265 := GCons(__e, symMessage, tmp9264)

tmp9266 := GCons(__e, tmp9250, tmp9265)

tmp9267 := GCons(__e, symTime, tmp9266)

tmp9268 := GCons(__e, tmp9247, tmp9267)

tmp9269 := GCons(__e, symFinish, tmp9268)

tmp9270 := GCons(__e, V7253, tmp9269)

tmp9271 := GCons(__e, symResult, tmp9270)

tmp9272 := GCons(__e, tmp9245, tmp9271)

tmp9273 := GCons(__e, symStart, tmp9272)

return GCons(__e, symlet, tmp9273)


panic("unreachable")
})

tmp9274 := Call2(__e, ns2_1set, symshen_4process_1time, tmp9243)


_ = tmp9274

tmp9275 := MakeNative1(func(__e *ControlFlow, V7254 Obj) Obj {
tmp9296 := GConsP(__e, V7254)

var ifres9284 Obj

if True == tmp9296 {
tmp9286 := GTl(__e, V7254)

let__8592 := tmp9286
_ = let__8592

tmp9295 := GConsP(__e, let__8592)

var ifres9287 Obj

if True == tmp9295 {
tmp9289 := GTl(__e, let__8592)

let__8593 := tmp9289
_ = let__8593

tmp9294 := GConsP(__e, let__8593)

var ifres9290 Obj

if True == tmp9294 {
tmp9292 := GTl(__e, let__8593)

tmp9293 := GConsP(__e, tmp9292)

var ifres9291 Obj

if True == tmp9293 {
ifres9291 = True


} else {
ifres9291 = False


}

ifres9290 = ifres9291


} else {
ifres9290 = False


}

var ifres9288 Obj

if True == ifres9290 {
ifres9288 = True


} else {
ifres9288 = False


}

ifres9287 = ifres9288


} else {
ifres9287 = False


}

var ifres9285 Obj

if True == ifres9287 {
ifres9285 = True


} else {
ifres9285 = False


}

ifres9284 = ifres9285


} else {
ifres9284 = False


}

if True == ifres9284 {
tmp9276 := GHd(__e, V7254)

let__8594 := tmp9276
_ = let__8594

tmp9277 := GTl(__e, V7254)

let__8595 := tmp9277
_ = let__8595

tmp9278 := GHd(__e, let__8595)

tmp9279 := GTl(__e, let__8595)

tmp9280 := GCons(__e, let__8594, tmp9279)

tmp9281 := GCons(__e, tmp9280, Nil)

tmp9282 := GCons(__e, tmp9278, tmp9281)

return GCons(__e, let__8594, tmp9282)


} else {
return V7254
}


panic("unreachable")
})

tmp9297 := Call2(__e, ns2_1set, symshen_4process_1assoc, tmp9275)


_ = tmp9297

tmp9298 := MakeNative1(func(__e *ControlFlow, V7259 Obj) Obj {
tmp9299 := PrimGuard1(__e, symstr, PrimStr, V7259)

tmp9300 := Call1(__e, SymFn(symshen_4mu_1h), tmp9299)


return PrimGuard1(__e, symintern, PrimIntern, tmp9300)


panic("unreachable")
})

tmp9301 := Call2(__e, ns2_1set, symshen_4make_1uppercase, tmp9298)


_ = tmp9301

tmp9302 := MakeNative1(func(__e *ControlFlow, V7260 Obj) Obj {
tmp9318 := GEq(__e, MakeString(""), V7260)

if True == tmp9318 {
return MakeString("")
} else {
tmp9316 := Call1(__e, SymFn(symshen_4_7string_2), V7260)


if True == tmp9316 {
tmp9303 := Call1(__e, SymFn(symhdstr), V7260)


tmp9304 := PrimGuard1(__e, symstring_1_6n, PrimStringToNumber, tmp9303)

let__8596 := tmp9304
_ = let__8596

let__8597 := GSub(__e, let__8596, MakeInteger(32))
_ = let__8597

tmp9312 := GGe(__e, let__8596, MakeInteger(97))

var ifres9309 Obj

if True == tmp9312 {
tmp9311 := GLe(__e, let__8596, MakeInteger(122))

var ifres9310 Obj

if True == tmp9311 {
ifres9310 = True


} else {
ifres9310 = False


}

ifres9309 = ifres9310


} else {
ifres9309 = False


}

var ifres9306 Obj

if True == ifres9309 {
tmp9307 := PrimGuard1(__e, symn_1_6string, PrimNumberToString, let__8597)

ifres9306 = tmp9307


} else {
tmp9308 := Call1(__e, SymFn(symhdstr), V7260)


ifres9306 = tmp9308


}

let__8598 := ifres9306
_ = let__8598

tmp9314 := Call1(__e, SymFn(symshen_4mu_1h), PrimGuard1(__e, symtlstr, PrimTailString, V7260))


return TailCall2(__e, SymFn(sym_8s), let__8598, tmp9314)


} else {
return TailCall1(__e, SymFn(symshen_4f_1error), symshen_4mu_1h)
}


}


panic("unreachable")
})

tmp9319 := Call2(__e, ns2_1set, symshen_4mu_1h, tmp9302)


_ = tmp9319

tmp9320 := MakeNative2(func(__e *ControlFlow, V7264, V7265 Obj) Obj {
tmp9321 := GValue(__e, sym_dmacros_d)

tmp9322 := Call3(__e, SymFn(symshen_4update_1assoc), V7264, V7265, tmp9321)


return PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp9322)


panic("unreachable")
})

tmp9323 := Call2(__e, ns2_1set, symshen_4record_1macro, tmp9320)


_ = tmp9323

tmp9324 := MakeNative3(func(__e *ControlFlow, V7275, V7276, V7277 Obj) Obj {
tmp9341 := GEq(__e, Nil, V7277)

if True == tmp9341 {
tmp9325 := GCons(__e, V7275, V7276)

return GCons(__e, tmp9325, Nil)


} else {
tmp9326 := MakeNative0(func(__e *ControlFlow) Obj {
return PrimGuard1(__e, symsimple_1error, PrimSimpleError, MakeString("implementation error in shen.update-assoc"))
panic("unreachable")
})

let__8599 := tmp9326
_ = let__8599

tmp9339 := GConsP(__e, V7277)

if True == tmp9339 {
tmp9327 := GHd(__e, V7277)

let__8600 := tmp9327
_ = let__8600

tmp9328 := GTl(__e, V7277)

let__8601 := tmp9328
_ = let__8601

tmp9337 := GConsP(__e, let__8600)

var ifres9333 Obj

if True == tmp9337 {
tmp9335 := GHd(__e, let__8600)

tmp9336 := GEq(__e, V7275, tmp9335)

var ifres9334 Obj

if True == tmp9336 {
ifres9334 = True


} else {
ifres9334 = False


}

ifres9333 = ifres9334


} else {
ifres9333 = False


}

if True == ifres9333 {
tmp9329 := GHd(__e, let__8600)

tmp9330 := GCons(__e, tmp9329, V7276)

return GCons(__e, tmp9330, let__8601)


} else {
tmp9331 := Call3(__e, SymFn(symshen_4update_1assoc), V7275, V7276, let__8601)


return GCons(__e, let__8600, tmp9331)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8599)
}


}


panic("unreachable")
})

tmp9342 := Call2(__e, ns2_1set, symshen_4update_1assoc, tmp9324)


_ = tmp9342

tmp9343 := MakeNative0(func(__e *ControlFlow) Obj {
tmp9351 := Call0(__e, SymFn(symstinput))


tmp9352 := Call1(__e, SymFn(symchar_1stream_2), tmp9351)


if True == tmp9352 {
tmp9344 := GCons(__e, symstinput, Nil)

tmp9345 := GCons(__e, tmp9344, Nil)

tmp9346 := GCons(__e, symread_1unit_1string, tmp9345)

tmp9347 := GCons(__e, tmp9346, Nil)

return GCons(__e, symstring_1_6n, tmp9347)


} else {
tmp9348 := GCons(__e, symstinput, Nil)

tmp9349 := GCons(__e, tmp9348, Nil)

return GCons(__e, symread_1byte, tmp9349)


}


panic("unreachable")
})

tmp9353 := Call2(__e, ns2_1set, symshen_4process_1read_1byte, tmp9343)


_ = tmp9353

tmp9354 := MakeNative1(func(__e *ControlFlow, V7281 Obj) Obj {
tmp9355 := GCons(__e, symshen_4prolog_1vector, Nil)

let__8602 := tmp9355
_ = let__8602

tmp9356 := GCons(__e, MakeInteger(0), Nil)

tmp9357 := GCons(__e, symvector, tmp9356)

tmp9358 := GCons(__e, tmp9357, Nil)

tmp9359 := GCons(__e, MakeInteger(0), tmp9358)

tmp9360 := GCons(__e, True, tmp9359)

tmp9361 := GCons(__e, sym_8v, tmp9360)

let__8603 := tmp9361
_ = let__8603

let__8604 := MakeInteger(0)
_ = let__8604

tmp9362 := GCons(__e, True, Nil)

tmp9363 := GCons(__e, symfreeze, tmp9362)

let__8605 := tmp9363
_ = let__8605

tmp9364 := MakeNative1(func(__e *ControlFlow, Z7287 Obj) Obj {
return TailCall1(__e, SymFn(symshen_4_5body_6), Z7287)
panic("unreachable")
})

tmp9365 := Call2(__e, SymFn(symcompile), tmp9364, V7281)


let__8606 := tmp9365
_ = let__8606

tmp9366 := Call1(__e, SymFn(symshen_4received), V7281)


let__8607 := tmp9366
_ = let__8607

tmp9367 := Call1(__e, SymFn(symgensym), symV)


let__8608 := tmp9367
_ = let__8608

tmp9368 := Call1(__e, SymFn(symgensym), symL)


let__8609 := tmp9368
_ = let__8609

tmp9369 := Call1(__e, SymFn(symgensym), symK)


let__8610 := tmp9369
_ = let__8610

tmp9370 := Call1(__e, SymFn(symgensym), symC)


let__8611 := tmp9370
_ = let__8611

tmp9371 := Call6(__e, SymFn(symshen_4continue), let__8607, let__8606, let__8608, let__8609, let__8610, let__8611)


tmp9372 := GCons(__e, tmp9371, Nil)

tmp9373 := GCons(__e, let__8611, tmp9372)

tmp9374 := GCons(__e, symlambda, tmp9373)

tmp9375 := GCons(__e, tmp9374, Nil)

tmp9376 := GCons(__e, let__8610, tmp9375)

tmp9377 := GCons(__e, symlambda, tmp9376)

tmp9378 := GCons(__e, tmp9377, Nil)

tmp9379 := GCons(__e, let__8609, tmp9378)

tmp9380 := GCons(__e, symlambda, tmp9379)

tmp9381 := GCons(__e, tmp9380, Nil)

tmp9382 := GCons(__e, let__8608, tmp9381)

tmp9383 := GCons(__e, symlambda, tmp9382)

let__8612 := tmp9383
_ = let__8612

tmp9384 := GCons(__e, let__8605, Nil)

tmp9385 := GCons(__e, let__8604, tmp9384)

tmp9386 := GCons(__e, let__8603, tmp9385)

tmp9387 := GCons(__e, let__8602, tmp9386)

return GCons(__e, let__8612, tmp9387)


panic("unreachable")
})

tmp9388 := Call2(__e, ns2_1set, symshen_4call_1prolog, tmp9354)


_ = tmp9388

tmp9389 := MakeNative1(func(__e *ControlFlow, V7296 Obj) Obj {
tmp9390 := MakeNative0(func(__e *ControlFlow) Obj {
return Nil
panic("unreachable")
})

let__8613 := tmp9390
_ = let__8613

tmp9405 := GConsP(__e, V7296)

if True == tmp9405 {
tmp9391 := GHd(__e, V7296)

let__8614 := tmp9391
_ = let__8614

tmp9392 := GTl(__e, V7296)

let__8615 := tmp9392
_ = let__8615

tmp9403 := GEq(__e, symreceive, let__8614)

var ifres9396 Obj

if True == tmp9403 {
tmp9402 := GConsP(__e, let__8615)

var ifres9398 Obj

if True == tmp9402 {
tmp9400 := GTl(__e, let__8615)

tmp9401 := GEq(__e, Nil, tmp9400)

var ifres9399 Obj

if True == tmp9401 {
ifres9399 = True


} else {
ifres9399 = False


}

ifres9398 = ifres9399


} else {
ifres9398 = False


}

var ifres9397 Obj

if True == ifres9398 {
ifres9397 = True


} else {
ifres9397 = False


}

ifres9396 = ifres9397


} else {
ifres9396 = False


}

if True == ifres9396 {
return let__8615
} else {
tmp9393 := Call1(__e, SymFn(symshen_4received), let__8614)


tmp9394 := Call1(__e, SymFn(symshen_4received), let__8615)


return TailCall2(__e, SymFn(symunion), tmp9393, tmp9394)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8613)
}


panic("unreachable")
})

tmp9406 := Call2(__e, ns2_1set, symshen_4received, tmp9389)


_ = tmp9406

tmp9407 := MakeNative0(func(__e *ControlFlow) Obj {
tmp9408 := GValue(__e, symshen_4_dprolog_1memory_d)

tmp9409 := PrimGuard1(__e, symabsvector, PrimAbsvector, tmp9408)

let__8616 := tmp9409
_ = let__8616

tmp9410 := GVecSet(__e, let__8616, MakeInteger(0), symshen_4print_1prolog_1vector)

let__8617 := tmp9410
_ = let__8617

tmp9411 := GVecSet(__e, let__8616, MakeInteger(1), MakeInteger(2))

let__8618 := tmp9411
_ = let__8618

return let__8618


panic("unreachable")
})

tmp9412 := Call2(__e, ns2_1set, symshen_4prolog_1vector, tmp9407)


_ = tmp9412

tmp9413 := MakeNative1(func(__e *ControlFlow, V7303 Obj) Obj {
return V7303
panic("unreachable")
})

tmp9414 := Call2(__e, ns2_1set, symreceive, tmp9413)


_ = tmp9414

tmp9415 := MakeNative1(func(__e *ControlFlow, V7304 Obj) Obj {
tmp9416 := MakeNative0(func(__e *ControlFlow) Obj {
return V7304
panic("unreachable")
})

let__8619 := tmp9416
_ = let__8619

tmp9433 := GConsP(__e, V7304)

if True == tmp9433 {
tmp9417 := GHd(__e, V7304)

let__8620 := tmp9417
_ = let__8620

tmp9418 := GTl(__e, V7304)

let__8621 := tmp9418
_ = let__8621

tmp9431 := GEq(__e, symbar_b, let__8620)

var ifres9424 Obj

if True == tmp9431 {
tmp9430 := GConsP(__e, let__8621)

var ifres9426 Obj

if True == tmp9430 {
tmp9428 := GTl(__e, let__8621)

tmp9429 := GEq(__e, Nil, tmp9428)

var ifres9427 Obj

if True == tmp9429 {
ifres9427 = True


} else {
ifres9427 = False


}

ifres9426 = ifres9427


} else {
ifres9426 = False


}

var ifres9425 Obj

if True == ifres9426 {
ifres9425 = True


} else {
ifres9425 = False


}

ifres9424 = ifres9425


} else {
ifres9424 = False


}

if True == ifres9424 {
return GHd(__e, let__8621)
} else {
tmp9419 := Call1(__e, SymFn(symshen_4rcons__form), let__8620)


tmp9420 := Call1(__e, SymFn(symshen_4rcons__form), let__8621)


tmp9421 := GCons(__e, tmp9420, Nil)

tmp9422 := GCons(__e, tmp9419, tmp9421)

return GCons(__e, symcons, tmp9422)


}


} else {
return TailCall1(__e, SymFn(symthaw), let__8619)
}


panic("unreachable")
})

tmp9434 := Call2(__e, ns2_1set, symshen_4rcons__form, tmp9415)


_ = tmp9434

tmp9435 := MakeNative1(func(__e *ControlFlow, V7308 Obj) Obj {
tmp9442 := GConsP(__e, V7308)

if True == tmp9442 {
tmp9436 := GHd(__e, V7308)

tmp9437 := GTl(__e, V7308)

tmp9438 := Call1(__e, SymFn(symshen_4tuple_1up), tmp9437)


tmp9439 := GCons(__e, tmp9438, Nil)

tmp9440 := GCons(__e, tmp9436, tmp9439)

return GCons(__e, sym_8p, tmp9440)


} else {
return V7308
}


panic("unreachable")
})

tmp9443 := Call2(__e, ns2_1set, symshen_4tuple_1up, tmp9435)


_ = tmp9443

tmp9444 := MakeNative1(func(__e *ControlFlow, V7309 Obj) Obj {
tmp9445 := GValue(__e, sym_dmacros_d)

tmp9446 := Call2(__e, SymFn(symassoc), V7309, tmp9445)


tmp9447 := GValue(__e, sym_dmacros_d)

tmp9448 := Call2(__e, SymFn(symremove), tmp9446, tmp9447)


tmp9449 := PrimGuard2(__e, symset, PrimSet, sym_dmacros_d, tmp9448)

_ = tmp9449

return V7309


panic("unreachable")
})

return TailCall2(__e, ns2_1set, symundefmacro, tmp9444)




panic("unreachable")
})

