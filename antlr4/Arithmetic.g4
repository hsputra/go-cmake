grammar Arithmetic;

expression: '(' expression ')'              # parensExpression
          | expression op=('*' | '/') expression   # mulDivExpression
          | expression op=('+' | '-') expression   # addSubExpression
          | NUMBER                            # numberExpression
          ;

NUMBER: [0-9]+ ('.' [0-9]+)?;
