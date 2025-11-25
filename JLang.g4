grammar JLang;

/*=====================================================
 =                   PARSER RULES                    =
 =====================================================*/

program
    : pkgDecl impList? topDecl* EOF
    ;

pkgDecl
    : KW_PKG IDENT
    ;

impList
    : impDecl+
    ;

impDecl
    : KW_IMP importSpec
    | KW_IMP LPAREN importSpec (COMMA importSpec)* RPAREN
    ;

importSpec
    : STRING_LIT
    | IDENT STRING_LIT
    ;

topDecl
    : varDecl
    | consDecl
    | typeDecl
    | funcDecl
    | stmt
    ;

/*---------------- VAR / CONST / TYPE ----------------*/

varDecl
    : KW_VAR varSpec (';' varSpec)* (';')?
    ;

varSpec
    : identList (type_ ('=' exprList)? | '=' exprList)?
    ;

consDecl
    : KW_CONS constSpec (';' constSpec)* (';')?
    ;

constSpec
    : identList (type_ '=' exprList | '=' exprList)
    ;

typeDecl
    : KW_TYPE typeSpec (';' typeSpec)* (';')?
    ;

typeSpec
    : IDENT ('=' type_ | type_)
    ;

/*-------------------- FUNCTIONS ---------------------*/

funcDecl
    : KW_DEF IDENT signature block
    ;

signature
    : LPAREN paramList? RPAREN resultType?
    ;

paramList
    : param (COMMA param)*
    ;

param
    : IDENT COLON type_
    ;

resultType
    : COLON type_
    ;

/*-------------------- BLOCKS ------------------------*/

block
    : LBRACE stmtList? RBRACE
    ;

stmtList
    : stmt (';' stmt)* (';')?
    ;

/*-------------------- STATEMENTS --------------------*/

stmt
    : simpleStmt
    | declStmt
    | ifStmt
    | switchStmt
    | forStmt
    | rangeStmt
    | selectStmt
    | deferStmt
    | returnStmt
    | breakStmt
    | continueStmt
    | jotoStmt
    | panicStmt
    | labeledStmt
    | block
    ;

declStmt
    : varDecl
    | consDecl
    | typeDecl
    ;

labeledStmt
    : IDENT COLON stmt
    ;

simpleStmt
    : exprStmt
    | assignStmt
    | shortVarDecl
    | sendStmt
    | spawnStmt
    ;

exprStmt
    : expr
    ;

assignStmt
    : lhs assignOp exprList
    ;

lhs
    : primaryExpr
    ;

assignOp
    : ASSIGN | ADDEQ | SUBEQ | MULEQ | DIVEQ | MODEQ
    | ANDEQ | OREQ | XOREQ | SHLEQ | SHREQ
    ;

shortVarDecl
    : identList DECL exprList
    ;

sendStmt
    : expr CH_SEND expr
    ;

spawnStmt
    : KW_J expr
    ;

returnStmt
    : KW_RET exprList?
    ;

deferStmt
    : KW_LATER expr
    ;

breakStmt
    : KW_BREAK IDENT?
    ;

continueStmt
    : KW_CONTINUE IDENT?
    ;

jotoStmt
    : KW_JOTO IDENT
    ;

/* NEW: panic statement */
panicStmt
    : KW_PANIC expr
    ;

/*-------------------- IF / SWITCH --------------------*/

ifStmt
    : KW_IF expr block elseOpt
    ;

elseOpt
    : KW_ELSE ifStmt
    | KW_ELSE block
    |
    ;

switchStmt
    : KW_SWITCH switchHead? LBRACE caseClauses? RBRACE
    ;

switchHead
    : expr
    ;

caseClauses
    : caseClause+
    ;

caseClause
    : KW_CASE exprList COLON stmtList fallOpt
    | KW_DFT COLON stmtList
    ;

fallOpt
    : KW_FALL
    |
    ;

/*-------------------- LOOPS --------------------------*/

forStmt
    : KW_FR forClause? block
    ;

forClause
    : simpleStmt? SEMI expr? SEMI simpleStmt?
    ;

rangeStmt
    : KW_FR KW_RANGE rangeHeader block
    ;

rangeHeader
    : (identList DECL)? expr
    ;

/*-------------------- SELECT -------------------------*/

selectStmt
    : KW_SELECT LBRACE commClauses? RBRACE
    ;

commClauses
    : commClause+
    ;

commClause
    : KW_CASE commClauseBody COLON stmtList
    | KW_DFT COLON stmtList
    ;

commClauseBody
    : sendStmt
    | recvStmt
    ;

recvStmt
    : exprList DECL expr
    | expr
    ;

/*-------------------- TYPES --------------------------*/

type_
    : TYPE_NAME
    | arrayType
    | sliceType
    | mappingType
    | channelType
    | structType
    | interfaceType
    | IDENT
    ;

arrayType
    : LBRACK INT_LIT RBRACK type_
    ;

sliceType
    : LBRACK RBRACK type_
    ;

mappingType
    : KW_MAPPING LBRACK type_ RBRACK type_
    ;

channelType
    : KW_CHANNEL type_
    ;

structType
    : KW_STRUCT LBRACE typeFieldList? RBRACE
    ;

typeFieldList
    : typeField (';' typeField)* (';')?
    ;

typeField
    : IDENT COLON type_
    ;

interfaceType
    : KW_INTERFACE LBRACE methodList? RBRACE
    ;

methodList
    : methodDecl (';' methodDecl)* (';')?
    ;

methodDecl
    : IDENT signature
    ;

/*-------------------- EXPRESSIONS --------------------*/

expr
    : logicalOrExpr
    ;

logicalOrExpr
    : logicalAndExpr (OROR logicalAndExpr)*
    ;

logicalAndExpr
    : bitOrExpr (ANDAND bitOrExpr)*
    ;

bitOrExpr
    : bitXorExpr (BOR bitXorExpr)*
    ;

bitXorExpr
    : bitAndExpr (BXOR bitAndExpr)*
    ;

bitAndExpr
    : equalityExpr (BAND equalityExpr)*
    ;

/* == != */
equalityExpr
    : relationalExpr ((EQ | NE) relationalExpr)*
    ;

/* < <= > >= */
relationalExpr
    : shiftExpr ((LT | LE | GT | GE) shiftExpr)*
    ;

/* << >> */
shiftExpr
    : additiveExpr ((SHL | SHR) additiveExpr)*
    ;

/* + - */
additiveExpr
    : multiplicativeExpr ((PLUS | MINUS) multiplicativeExpr)*
    ;

/* * / % */
multiplicativeExpr
    : unaryExpr ((STAR | SLASH | PERCENT) unaryExpr)*
    ;

unaryExpr
    : primaryExpr
    | PLUS unaryExpr
    | MINUS unaryExpr
    | BANG unaryExpr
    | TILDE unaryExpr
    | STAR unaryExpr
    | BAND unaryExpr
    | CH_SEND unaryExpr
    ;

primaryExpr
    : operand
    | makeExpr         
    | primaryExpr selector
    | primaryExpr index
    | primaryExpr call
    ;

operand
    : IDENT
    | INT_LIT
    | FLOAT_LIT
    | STRING_LIT
    | CHAR_LIT
    | RAW_STR
    | LPAREN expr RPAREN
    | KW_RECOVER LPAREN RPAREN
    ;

makeExpr
    : KW_MAKE LPAREN type_ (COMMA exprList)? RPAREN
    ;

selector
    : DOT IDENT
    ;

index
    : LBRACK expr RBRACK
    ;

call
    : LPAREN argList? RPAREN
    ;

argList
    : expr (COMMA expr)*
    ;

identList
    : IDENT (COMMA IDENT)*
    ;

exprList
    : expr (COMMA expr)*
    ;

/*=====================================================
 =                     LEXER RULES                   =
 =====================================================*/

KW_PKG      : 'pkg';
KW_IMP      : 'imp';
KW_DEF      : 'def';
KW_VAR      : 'var';
KW_CONS     : 'cons';
KW_TYPE     : 'type';
KW_STRUCT   : 'struct';
KW_INTERFACE: 'interface';
KW_MAPPING  : 'mapping';
KW_CHANNEL  : 'channel';
KW_J        : 'j';
KW_SELECT   : 'select';
KW_LATER    : 'later';
KW_RET      : 'ret';
KW_IF       : 'if';
KW_ELSE     : 'else';
KW_SWITCH   : 'switch';
KW_CASE     : 'case';
KW_FALL     : 'fall';
KW_FR       : 'fr';
KW_RANGE    : 'range';
KW_BREAK    : 'break';
KW_CONTINUE : 'continue';
KW_JOTO     : 'joto';
KW_DFT      : 'dft';
KW_PANIC    : 'panic';
KW_RECOVER  : 'recover';
KW_MAKE     : 'make'; 

TYPE_NAME
    : 'i8'|'i16'|'i32'|'i64'
    | 'u8'|'u16'|'u32'|'u64'
    | 'f32'|'f64'
    | 'bool'
    | 'string'
    ;

IDENT
    : [_\p{L}][_\p{L}\p{N}]*
    ;

INT_LIT
    : '0'
    | [1-9][0-9_]*
    | '0'[xX][0-9a-fA-F_]+
    | '0'[bB][01_]+
    | '0'[oO][0-7_]+
    ;

FLOAT_LIT
    : [0-9][0-9_]*'.'[0-9][0-9_]*([eE][+\-]?[0-9][0-9_]*)?
    | [0-9][0-9_]*[eE][+\-]?[0-9][0-9_]*
    | '.'[0-9][0-9_]*([eE][+\-]?[0-9][0-9_]*)?
    ;

STRING_LIT
    : '"' ( '\\' . | ~["\\\r\n] )* '"'
    ;

/* 支持跨行 raw string，例如 Go 风格的 `...` */
RAW_STR
    : '`' ( . | '\r' | '\n' )*? '`'
    ;

CHAR_LIT
    : '\'' ( '\\' . | ~['\\\r\n] ) '\''
    ;

CH_SEND : '<-';
LE      : '<=';
GE      : '>=';
EQ      : '==';
NE      : '!=';
SHLEQ   : '<<=';
SHREQ   : '>>=';
SHL     : '<<';
SHR     : '>>';
ADDEQ   : '+=';
SUBEQ   : '-=';
MULEQ   : '*=';
DIVEQ   : '/=';
MODEQ   : '%=';
ANDEQ   : '&=';
OREQ    : '|=';
XOREQ   : '^=';

LPAREN  : '(';
RPAREN  : ')';
LBRACE  : '{';
RBRACE  : '}';
LBRACK  : '[';
RBRACK  : ']';
COMMA   : ',';
SEMI    : ';';
COLON   : ':';
DOT     : '.';
ASSIGN  : '=';
DECL    : ':=';
PLUS    : '+';
MINUS   : '-';
STAR    : '*';
SLASH   : '/';
PERCENT : '%';
LT      : '<';
GT      : '>';
BAND    : '&';
BOR     : '|';
BXOR    : '^';
ANDAND  : '&&';
OROR    : '||';
BANG    : '!';
TILDE   : '~';

WS : [ \t\r\n]+ -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

BLOCK_COMMENT
    : '/*' ( . | '\r' | '\n' )*? '*/' -> skip
    ;
