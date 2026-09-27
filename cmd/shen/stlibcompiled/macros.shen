(package maths [expt =r gcd lcd isqrt sqrt nthrt floor ceiling round mod lcm random min max
                reseed ~ positive? negative? natural? converge series odd? even?
                cos sin tan radians pi e tan30 cos30 cos45 sin45 sqrt2 tan60 sin120
                tan120 sin135 cos135 cos150 tan150 cos210 tan210 sin225 cos225 sin240
                tan240 sin300 tan300 sin315 cos315 cos330 tan330 sinh cosh tanh sech
                csch power factorial prime? unix div modf product summation set-tolerance tolerance
                coth for sq cube newv abs approx log log2 loge log10 g to stop done]

(defmacro maths-macro
  [log10 N] -> [log10 N [tolerance]]
  [log2 N] -> [log2 N [tolerance]]
  [loge N] -> [loge N [tolerance]]
  [log M N] -> [log M N [tolerance]]
  [sin N] -> [sin N [tolerance]]
  [tan N] -> [tan N [tolerance]]
  [cos N] -> [cos N [tolerance]]
  [tanh N] -> [tanh N [tolerance]]
  [cosh N] -> [cosh N [tolerance]]
  [sinh N] -> [sinh N [tolerance]]
  [sech N] -> [sech N [tolerance]]
  [csch N] -> [csch N [tolerance]]
  [coth N] -> [coth N [tolerance]]
  [nthrt N Root] -> [nthrt N Root [tolerance]]
  [sqrt N] -> [sqrt N [tolerance]]
  [expt M N] -> [expt M N [tolerance]]
  [max W X Y | Z] -> [max W [max X Y | Z]]
  [min W X Y | Z] -> [min W [min X Y | Z]]
  [tolerance N] -> [tolerance=n N]
  [for X = Val stop Stop | Options+Procedure] -> [upto Val Stop | (process-options X Options+Procedure)]
  [for X = Val to N | Options+Procedure] -> [upto Val [< N]  | (process-options X Options+Procedure)])
)

(package numerals [hex octal duodecimal binary
                   n-op2 n-op1 n+ n- n* n/
                   | (append (external numerals) (external maths))]

(defmacro numeral-macro
  [n-op2 Op M N] -> [n-op2 Op M N [radix M]]
  [n-op1 Op M] -> [n-op1 Op M [radix M]]
  [n+ M N] -> [n+ M N [radix M]]
  [n- M N] -> [n- M N [radix M]]
  [n* M N] -> [n* M N [radix M]]
  [n/ M N] -> [n/ M N [radix M]])
)

(defmacro string-macros
  [s-op1 F S] -> [s-op1 F S [/. (protect X) (protect X)]]
  [s-op2 F S1 S2] -> [s-op2 F S1 S2 [/. (protect X) (protect X)]])

(package vector [newv array for to]

(defmacro vector-macros

  \\ read access
  [:= V [cons I []]] -> [<-vector V I]

  [:= V [cons I Is]] -> [:= [<-vector V I] Is]

  \\ write access
  [V [cons I []] := X] -> [vector-> V I X]

  [V [cons I Is] := X] -> (let V2 (newv)
                               [let V2 [<-vector V I]
                                       [V2 Is := X]])
  \\ array construction
  [array [cons Dim []]] -> [vector Dim]
  [array [cons Dim Dims]]
    -> (let V (newv)
            N (newv)
         [let V [vector Dim]
           [do [for N = 1 to Dim
                 [vector-> V N [array Dims]]]
               V]]))
)

(package print (append (external maths)
                       [pps pprint pretty-string linelength indentation set-linelength set-indentation])

(datatype print

   _______________________________
   (value *indentation*) : number;

   _______________________________
   (value *linelength*) : number;)

(defmacro pprint-macro
  [pprint X] -> [pprint X [stoutput]]
  [pps F] -> [pps F [stoutput]])
)

(package file [append-files append-files-with-open-stream mapc
                copy-file file-size reopen errout copy-file-with-open-stream
                file-exists? newv ascii]

(defmacro file-macro
  [errout X Default ErrFile] -> (let  Err (newv)
                                      Open (newv)
                                      Record (newv)
                                      Close (newv)
                                      E (newv)
                                      [trap-error X [/. E [let  Err [error-to-string E]
                                                                Open [trap-error [reopen ErrFile] [/. E [open ErrFile out]]]
                                                                Record [pr Err Open]
                                                                Close [close Open]
                                                                Default]]]))
)
