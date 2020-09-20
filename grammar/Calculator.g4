grammar Calculator;


// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
POW: '**';

NUMBER: [0-9]+;



WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;


expression: '(' expression ')' #parenthesis
   | expression POW expression # Pow
   | expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | NUMBER                             # Number
   ;


