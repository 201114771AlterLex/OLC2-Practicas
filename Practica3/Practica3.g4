
//
grammar Practica3;

options{
    language='Go';
}

@parser::members{
    type Data struct{
        DataType string
        DataValue string
    }

}

PRight: ')';
PLeft: '(';
Plus: '+';
Times: '*';
Point: '.';
Numbers:[0-9]+;
Decimal: Numbers Point Numbers;
Whitespace: [ \r\n\t]+ -> skip;

start
    : e EOF 
            {
                //fmt.Println($e.D1.DataValue)
            } 
    ;

e   returns [*Data D1]
    : t1=t  ep1=ep[$t1.D1] 
            {
                $D1=$ep1.D1
            } 
    ;

ep [*Data D] returns [*Data D1] locals [*Data DTemp]
    : Plus t1=t
        {
            $DTemp=&Data{}
            if ($D.DataType=="D" || $t1.D1.DataType=="D"){   
                $DTemp.DataType="D"
                NT1,_:=strconv.ParseFloat($D.DataValue,64)
                NT2,_:=strconv.ParseFloat($t1.D1.DataValue,64)
                Result:=NT1+NT2
                $DTemp.DataValue=fmt.Sprintf("%g",Result)                        
            }else{
                $DTemp.DataType="E"
                NT1,_:=strconv.Atoi($D.DataValue)
                NT2,_:=strconv.Atoi($t1.D1.DataValue)
                Result:=NT1+NT2
                $DTemp.DataValue=strconv.Itoa(Result)
            }
        }
      ep1=ep [$DTemp] 
            {
                $D1=$ep1.D1
            }
    |{$D1=$D} 
    ;

t returns [*Data D1]
    : f1=f tp1=tp [$f1.D1]
                {
                    $D1=$tp1.D1
                } 
    ;

tp [*Data D] returns [*Data D1] locals [*Data DTemp]
    : Times f1=f
                {
                    $DTemp=&Data{}
                    if ($D.DataType=="D" || $f1.D1.DataType=="D"){   
                        $DTemp.DataType="D"
                        NT1,_:=strconv.ParseFloat($D.DataValue,64)
                        NT2,_:=strconv.ParseFloat($f1.D1.DataValue,64)
                        Result:=NT1*NT2
                        $DTemp.DataValue=fmt.Sprintf("%g",Result)                        
                    }else{
                        $DTemp.DataType="E"
                        NT1,_:=strconv.Atoi($D.DataValue)
                        NT2,_:=strconv.Atoi($f1.D1.DataValue)
                        Result:=NT1*NT2
                        $DTemp.DataValue=strconv.Itoa(Result)
                    }
                } 
     tp1=tp[$DTemp] 
                {
                    $D1=$tp1.D1
                } 
    | {$D1=$D} 
    ;

f   returns [*Data D1]
: PLeft e PRight    {
                        $D1 = $e.D1
                    } 
| ne=Numbers    {
                    $D1 = &Data{}
                    $D1.DataType="E"
                    $D1.DataValue=$ne.text
                } 
| nd=Decimal {
                $D1 = &Data{}
                $D1.DataType="D"
                $D1.DataValue=$nd.text
            } 
;