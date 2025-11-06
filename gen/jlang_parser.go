// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // JLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JLangParser struct {
	*antlr.BaseParser
}

var JLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jlangParserInit() {
	staticData := &JLangParserStaticData
	staticData.LiteralNames = []string{
		"", "'pkg'", "'imp'", "'def'", "'var'", "'cons'", "'type'", "'struct'",
		"'interface'", "'mapping'", "'channel'", "'j'", "'select'", "'later'",
		"'ret'", "'if'", "'else'", "'switch'", "'case'", "'fall'", "'fr'", "'range'",
		"'break'", "'continue'", "'joto'", "'dft'", "'panic'", "'recover'",
		"", "", "", "", "", "", "", "'<-'", "'<='", "'>='", "'=='", "'!='",
		"'<<='", "'>>='", "'<<'", "'>>'", "'+='", "'-='", "'*='", "'/='", "'%='",
		"'&='", "'|='", "'^='", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
		"';'", "':'", "'.'", "'='", "':='", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'<'", "'>'", "'&'", "'|'", "'^'", "'&&'", "'||'", "'!'", "'~'",
	}
	staticData.SymbolicNames = []string{
		"", "KW_PKG", "KW_IMP", "KW_DEF", "KW_VAR", "KW_CONS", "KW_TYPE", "KW_STRUCT",
		"KW_INTERFACE", "KW_MAPPING", "KW_CHANNEL", "KW_J", "KW_SELECT", "KW_LATER",
		"KW_RET", "KW_IF", "KW_ELSE", "KW_SWITCH", "KW_CASE", "KW_FALL", "KW_FR",
		"KW_RANGE", "KW_BREAK", "KW_CONTINUE", "KW_JOTO", "KW_DFT", "KW_PANIC",
		"KW_RECOVER", "TYPE_NAME", "IDENT", "INT_LIT", "FLOAT_LIT", "STRING_LIT",
		"RAW_STR", "CHAR_LIT", "CH_SEND", "LE", "GE", "EQ", "NE", "SHLEQ", "SHREQ",
		"SHL", "SHR", "ADDEQ", "SUBEQ", "MULEQ", "DIVEQ", "MODEQ", "ANDEQ",
		"OREQ", "XOREQ", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"COMMA", "SEMI", "COLON", "DOT", "ASSIGN", "DECL", "PLUS", "MINUS",
		"STAR", "SLASH", "PERCENT", "LT", "GT", "BAND", "BOR", "BXOR", "ANDAND",
		"OROR", "BANG", "TILDE", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "pkgDecl", "impList", "impDecl", "importSpec", "topDecl",
		"varDecl", "varSpec", "consDecl", "constSpec", "typeDecl", "typeSpec",
		"funcDecl", "signature", "paramList", "param", "resultType", "block",
		"stmtList", "stmt", "declStmt", "labeledStmt", "simpleStmt", "exprStmt",
		"assignStmt", "lhs", "assignOp", "shortVarDecl", "sendStmt", "spawnStmt",
		"returnStmt", "deferStmt", "breakStmt", "continueStmt", "jotoStmt",
		"ifStmt", "elseOpt", "switchStmt", "switchHead", "caseClauses", "caseClause",
		"fallOpt", "forStmt", "forClause", "rangeStmt", "rangeHeader", "selectStmt",
		"commClauses", "commClause", "commClauseBody", "recvStmt", "type_",
		"arrayType", "sliceType", "mappingType", "channelType", "structType",
		"typeFieldList", "typeField", "interfaceType", "methodList", "methodDecl",
		"expr", "unaryExpr", "primaryExpr", "operand", "selector", "index",
		"call", "argList", "identList", "exprList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 668, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 1, 0, 1, 0, 3, 0, 147,
		8, 0, 1, 0, 5, 0, 150, 8, 0, 10, 0, 12, 0, 153, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 2, 4, 2, 161, 8, 2, 11, 2, 12, 2, 162, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 172, 8, 3, 10, 3, 12, 3, 175, 9, 3, 1, 3,
		1, 3, 3, 3, 179, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 184, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 191, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 197, 8, 6,
		10, 6, 12, 6, 200, 9, 6, 1, 6, 3, 6, 203, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		3, 7, 209, 8, 7, 1, 7, 1, 7, 3, 7, 213, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 219, 8, 8, 10, 8, 12, 8, 222, 9, 8, 1, 8, 3, 8, 225, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 234, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 240, 8, 10, 10, 10, 12, 10, 243, 9, 10, 1, 10, 3, 10, 246,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 252, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 261, 8, 13, 1, 13, 1, 13, 3, 13,
		265, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 270, 8, 14, 10, 14, 12, 14, 273,
		9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 288, 8, 17, 1, 18, 1, 18, 1, 18, 5, 18,
		293, 8, 18, 10, 18, 12, 18, 296, 9, 18, 1, 18, 3, 18, 299, 8, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 3, 19, 315, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 320, 8,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		331, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 356, 8, 30, 1, 31, 1, 31, 1, 31, 1,
		32, 1, 32, 3, 32, 363, 8, 32, 1, 33, 1, 33, 3, 33, 367, 8, 33, 1, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 382, 8, 36, 1, 37, 1, 37, 3, 37, 386, 8, 37, 1, 37, 1, 37,
		3, 37, 390, 8, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 4, 39, 397, 8, 39,
		11, 39, 12, 39, 398, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 3, 40, 410, 8, 40, 1, 41, 1, 41, 3, 41, 414, 8, 41, 1, 42, 1,
		42, 3, 42, 418, 8, 42, 1, 42, 1, 42, 1, 43, 3, 43, 423, 8, 43, 1, 43, 1,
		43, 3, 43, 427, 8, 43, 1, 43, 1, 43, 3, 43, 431, 8, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 3, 45, 441, 8, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 3, 46, 448, 8, 46, 1, 46, 1, 46, 1, 47, 4, 47, 453,
		8, 47, 11, 47, 12, 47, 454, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 3, 48, 465, 8, 48, 1, 49, 1, 49, 3, 49, 469, 8, 49, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 3, 50, 476, 8, 50, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 486, 8, 51, 1, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 3, 56, 509, 8, 56, 1,
		56, 1, 56, 1, 57, 1, 57, 1, 57, 5, 57, 516, 8, 57, 10, 57, 12, 57, 519,
		9, 57, 1, 57, 3, 57, 522, 8, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1,
		59, 1, 59, 3, 59, 531, 8, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 5, 60,
		538, 8, 60, 10, 60, 12, 60, 541, 9, 60, 1, 60, 3, 60, 544, 8, 60, 1, 61,
		1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1,
		62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62,
		1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1,
		62, 1, 62, 1, 62, 1, 62, 5, 62, 582, 8, 62, 10, 62, 12, 62, 585, 9, 62,
		1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1,
		63, 1, 63, 1, 63, 3, 63, 600, 8, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64,
		1, 64, 1, 64, 1, 64, 1, 64, 5, 64, 611, 8, 64, 10, 64, 12, 64, 614, 9,
		64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65,
		1, 65, 1, 65, 1, 65, 3, 65, 629, 8, 65, 1, 66, 1, 66, 1, 66, 1, 67, 1,
		67, 1, 67, 1, 67, 1, 68, 1, 68, 3, 68, 640, 8, 68, 1, 68, 1, 68, 1, 69,
		1, 69, 1, 69, 5, 69, 647, 8, 69, 10, 69, 12, 69, 650, 9, 69, 1, 70, 1,
		70, 1, 70, 5, 70, 655, 8, 70, 10, 70, 12, 70, 658, 9, 70, 1, 71, 1, 71,
		1, 71, 5, 71, 663, 8, 71, 10, 71, 12, 71, 666, 9, 71, 1, 71, 0, 2, 124,
		128, 72, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134,
		136, 138, 140, 142, 0, 6, 3, 0, 40, 41, 44, 51, 62, 62, 1, 0, 38, 39, 2,
		0, 36, 37, 69, 70, 1, 0, 42, 43, 1, 0, 64, 65, 1, 0, 66, 68, 704, 0, 144,
		1, 0, 0, 0, 2, 156, 1, 0, 0, 0, 4, 160, 1, 0, 0, 0, 6, 178, 1, 0, 0, 0,
		8, 183, 1, 0, 0, 0, 10, 190, 1, 0, 0, 0, 12, 192, 1, 0, 0, 0, 14, 204,
		1, 0, 0, 0, 16, 214, 1, 0, 0, 0, 18, 226, 1, 0, 0, 0, 20, 235, 1, 0, 0,
		0, 22, 247, 1, 0, 0, 0, 24, 253, 1, 0, 0, 0, 26, 258, 1, 0, 0, 0, 28, 266,
		1, 0, 0, 0, 30, 274, 1, 0, 0, 0, 32, 278, 1, 0, 0, 0, 34, 287, 1, 0, 0,
		0, 36, 289, 1, 0, 0, 0, 38, 314, 1, 0, 0, 0, 40, 319, 1, 0, 0, 0, 42, 321,
		1, 0, 0, 0, 44, 330, 1, 0, 0, 0, 46, 332, 1, 0, 0, 0, 48, 334, 1, 0, 0,
		0, 50, 338, 1, 0, 0, 0, 52, 340, 1, 0, 0, 0, 54, 342, 1, 0, 0, 0, 56, 346,
		1, 0, 0, 0, 58, 350, 1, 0, 0, 0, 60, 353, 1, 0, 0, 0, 62, 357, 1, 0, 0,
		0, 64, 360, 1, 0, 0, 0, 66, 364, 1, 0, 0, 0, 68, 368, 1, 0, 0, 0, 70, 371,
		1, 0, 0, 0, 72, 381, 1, 0, 0, 0, 74, 383, 1, 0, 0, 0, 76, 393, 1, 0, 0,
		0, 78, 396, 1, 0, 0, 0, 80, 409, 1, 0, 0, 0, 82, 413, 1, 0, 0, 0, 84, 415,
		1, 0, 0, 0, 86, 422, 1, 0, 0, 0, 88, 432, 1, 0, 0, 0, 90, 440, 1, 0, 0,
		0, 92, 444, 1, 0, 0, 0, 94, 452, 1, 0, 0, 0, 96, 464, 1, 0, 0, 0, 98, 468,
		1, 0, 0, 0, 100, 475, 1, 0, 0, 0, 102, 485, 1, 0, 0, 0, 104, 487, 1, 0,
		0, 0, 106, 492, 1, 0, 0, 0, 108, 496, 1, 0, 0, 0, 110, 502, 1, 0, 0, 0,
		112, 505, 1, 0, 0, 0, 114, 512, 1, 0, 0, 0, 116, 523, 1, 0, 0, 0, 118,
		527, 1, 0, 0, 0, 120, 534, 1, 0, 0, 0, 122, 545, 1, 0, 0, 0, 124, 548,
		1, 0, 0, 0, 126, 599, 1, 0, 0, 0, 128, 601, 1, 0, 0, 0, 130, 628, 1, 0,
		0, 0, 132, 630, 1, 0, 0, 0, 134, 633, 1, 0, 0, 0, 136, 637, 1, 0, 0, 0,
		138, 643, 1, 0, 0, 0, 140, 651, 1, 0, 0, 0, 142, 659, 1, 0, 0, 0, 144,
		146, 3, 2, 1, 0, 145, 147, 3, 4, 2, 0, 146, 145, 1, 0, 0, 0, 146, 147,
		1, 0, 0, 0, 147, 151, 1, 0, 0, 0, 148, 150, 3, 10, 5, 0, 149, 148, 1, 0,
		0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 155, 5, 0, 0, 1, 155,
		1, 1, 0, 0, 0, 156, 157, 5, 1, 0, 0, 157, 158, 5, 29, 0, 0, 158, 3, 1,
		0, 0, 0, 159, 161, 3, 6, 3, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 5, 1, 0, 0, 0, 164,
		165, 5, 2, 0, 0, 165, 179, 3, 8, 4, 0, 166, 167, 5, 2, 0, 0, 167, 168,
		5, 52, 0, 0, 168, 173, 3, 8, 4, 0, 169, 170, 5, 58, 0, 0, 170, 172, 3,
		8, 4, 0, 171, 169, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0,
		0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176,
		177, 5, 53, 0, 0, 177, 179, 1, 0, 0, 0, 178, 164, 1, 0, 0, 0, 178, 166,
		1, 0, 0, 0, 179, 7, 1, 0, 0, 0, 180, 184, 5, 32, 0, 0, 181, 182, 5, 29,
		0, 0, 182, 184, 5, 32, 0, 0, 183, 180, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		184, 9, 1, 0, 0, 0, 185, 191, 3, 12, 6, 0, 186, 191, 3, 16, 8, 0, 187,
		191, 3, 20, 10, 0, 188, 191, 3, 24, 12, 0, 189, 191, 3, 38, 19, 0, 190,
		185, 1, 0, 0, 0, 190, 186, 1, 0, 0, 0, 190, 187, 1, 0, 0, 0, 190, 188,
		1, 0, 0, 0, 190, 189, 1, 0, 0, 0, 191, 11, 1, 0, 0, 0, 192, 193, 5, 4,
		0, 0, 193, 198, 3, 14, 7, 0, 194, 195, 5, 59, 0, 0, 195, 197, 3, 14, 7,
		0, 196, 194, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 203,
		5, 59, 0, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 13, 1, 0,
		0, 0, 204, 212, 3, 140, 70, 0, 205, 208, 3, 102, 51, 0, 206, 207, 5, 62,
		0, 0, 207, 209, 3, 142, 71, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0,
		0, 209, 213, 1, 0, 0, 0, 210, 211, 5, 62, 0, 0, 211, 213, 3, 142, 71, 0,
		212, 205, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		15, 1, 0, 0, 0, 214, 215, 5, 5, 0, 0, 215, 220, 3, 18, 9, 0, 216, 217,
		5, 59, 0, 0, 217, 219, 3, 18, 9, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1,
		0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 224, 1, 0, 0,
		0, 222, 220, 1, 0, 0, 0, 223, 225, 5, 59, 0, 0, 224, 223, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 17, 1, 0, 0, 0, 226, 233, 3, 140, 70, 0, 227, 228,
		3, 102, 51, 0, 228, 229, 5, 62, 0, 0, 229, 230, 3, 142, 71, 0, 230, 234,
		1, 0, 0, 0, 231, 232, 5, 62, 0, 0, 232, 234, 3, 142, 71, 0, 233, 227, 1,
		0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 19, 1, 0, 0, 0, 235, 236, 5, 6, 0,
		0, 236, 241, 3, 22, 11, 0, 237, 238, 5, 59, 0, 0, 238, 240, 3, 22, 11,
		0, 239, 237, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241,
		242, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 246,
		5, 59, 0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 21, 1, 0,
		0, 0, 247, 251, 5, 29, 0, 0, 248, 249, 5, 62, 0, 0, 249, 252, 3, 102, 51,
		0, 250, 252, 3, 102, 51, 0, 251, 248, 1, 0, 0, 0, 251, 250, 1, 0, 0, 0,
		252, 23, 1, 0, 0, 0, 253, 254, 5, 3, 0, 0, 254, 255, 5, 29, 0, 0, 255,
		256, 3, 26, 13, 0, 256, 257, 3, 34, 17, 0, 257, 25, 1, 0, 0, 0, 258, 260,
		5, 52, 0, 0, 259, 261, 3, 28, 14, 0, 260, 259, 1, 0, 0, 0, 260, 261, 1,
		0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 264, 5, 53, 0, 0, 263, 265, 3, 32,
		16, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 27, 1, 0, 0, 0,
		266, 271, 3, 30, 15, 0, 267, 268, 5, 58, 0, 0, 268, 270, 3, 30, 15, 0,
		269, 267, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271,
		272, 1, 0, 0, 0, 272, 29, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 5,
		29, 0, 0, 275, 276, 5, 60, 0, 0, 276, 277, 3, 102, 51, 0, 277, 31, 1, 0,
		0, 0, 278, 279, 5, 60, 0, 0, 279, 280, 3, 102, 51, 0, 280, 33, 1, 0, 0,
		0, 281, 282, 5, 54, 0, 0, 282, 288, 5, 55, 0, 0, 283, 284, 5, 54, 0, 0,
		284, 285, 3, 36, 18, 0, 285, 286, 5, 55, 0, 0, 286, 288, 1, 0, 0, 0, 287,
		281, 1, 0, 0, 0, 287, 283, 1, 0, 0, 0, 288, 35, 1, 0, 0, 0, 289, 294, 3,
		38, 19, 0, 290, 291, 5, 59, 0, 0, 291, 293, 3, 38, 19, 0, 292, 290, 1,
		0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0,
		0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 299, 5, 59, 0, 0, 298,
		297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 37, 1, 0, 0, 0, 300, 315, 3,
		44, 22, 0, 301, 315, 3, 40, 20, 0, 302, 315, 3, 70, 35, 0, 303, 315, 3,
		74, 37, 0, 304, 315, 3, 84, 42, 0, 305, 315, 3, 88, 44, 0, 306, 315, 3,
		92, 46, 0, 307, 315, 3, 62, 31, 0, 308, 315, 3, 60, 30, 0, 309, 315, 3,
		64, 32, 0, 310, 315, 3, 66, 33, 0, 311, 315, 3, 68, 34, 0, 312, 315, 3,
		42, 21, 0, 313, 315, 3, 34, 17, 0, 314, 300, 1, 0, 0, 0, 314, 301, 1, 0,
		0, 0, 314, 302, 1, 0, 0, 0, 314, 303, 1, 0, 0, 0, 314, 304, 1, 0, 0, 0,
		314, 305, 1, 0, 0, 0, 314, 306, 1, 0, 0, 0, 314, 307, 1, 0, 0, 0, 314,
		308, 1, 0, 0, 0, 314, 309, 1, 0, 0, 0, 314, 310, 1, 0, 0, 0, 314, 311,
		1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 313, 1, 0, 0, 0, 315, 39, 1, 0,
		0, 0, 316, 320, 3, 12, 6, 0, 317, 320, 3, 16, 8, 0, 318, 320, 3, 20, 10,
		0, 319, 316, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 318, 1, 0, 0, 0, 320,
		41, 1, 0, 0, 0, 321, 322, 5, 29, 0, 0, 322, 323, 5, 60, 0, 0, 323, 324,
		3, 38, 19, 0, 324, 43, 1, 0, 0, 0, 325, 331, 3, 46, 23, 0, 326, 331, 3,
		48, 24, 0, 327, 331, 3, 54, 27, 0, 328, 331, 3, 56, 28, 0, 329, 331, 3,
		58, 29, 0, 330, 325, 1, 0, 0, 0, 330, 326, 1, 0, 0, 0, 330, 327, 1, 0,
		0, 0, 330, 328, 1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331, 45, 1, 0, 0, 0,
		332, 333, 3, 124, 62, 0, 333, 47, 1, 0, 0, 0, 334, 335, 3, 50, 25, 0, 335,
		336, 3, 52, 26, 0, 336, 337, 3, 142, 71, 0, 337, 49, 1, 0, 0, 0, 338, 339,
		3, 128, 64, 0, 339, 51, 1, 0, 0, 0, 340, 341, 7, 0, 0, 0, 341, 53, 1, 0,
		0, 0, 342, 343, 3, 140, 70, 0, 343, 344, 5, 63, 0, 0, 344, 345, 3, 142,
		71, 0, 345, 55, 1, 0, 0, 0, 346, 347, 3, 124, 62, 0, 347, 348, 5, 35, 0,
		0, 348, 349, 3, 124, 62, 0, 349, 57, 1, 0, 0, 0, 350, 351, 5, 11, 0, 0,
		351, 352, 3, 124, 62, 0, 352, 59, 1, 0, 0, 0, 353, 355, 5, 14, 0, 0, 354,
		356, 3, 142, 71, 0, 355, 354, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 61,
		1, 0, 0, 0, 357, 358, 5, 13, 0, 0, 358, 359, 3, 124, 62, 0, 359, 63, 1,
		0, 0, 0, 360, 362, 5, 22, 0, 0, 361, 363, 5, 29, 0, 0, 362, 361, 1, 0,
		0, 0, 362, 363, 1, 0, 0, 0, 363, 65, 1, 0, 0, 0, 364, 366, 5, 23, 0, 0,
		365, 367, 5, 29, 0, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367,
		67, 1, 0, 0, 0, 368, 369, 5, 24, 0, 0, 369, 370, 5, 29, 0, 0, 370, 69,
		1, 0, 0, 0, 371, 372, 5, 15, 0, 0, 372, 373, 3, 124, 62, 0, 373, 374, 3,
		34, 17, 0, 374, 375, 3, 72, 36, 0, 375, 71, 1, 0, 0, 0, 376, 377, 5, 16,
		0, 0, 377, 382, 3, 70, 35, 0, 378, 379, 5, 16, 0, 0, 379, 382, 3, 34, 17,
		0, 380, 382, 1, 0, 0, 0, 381, 376, 1, 0, 0, 0, 381, 378, 1, 0, 0, 0, 381,
		380, 1, 0, 0, 0, 382, 73, 1, 0, 0, 0, 383, 385, 5, 17, 0, 0, 384, 386,
		3, 76, 38, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 1,
		0, 0, 0, 387, 389, 5, 54, 0, 0, 388, 390, 3, 78, 39, 0, 389, 388, 1, 0,
		0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 5, 55, 0, 0,
		392, 75, 1, 0, 0, 0, 393, 394, 3, 124, 62, 0, 394, 77, 1, 0, 0, 0, 395,
		397, 3, 80, 40, 0, 396, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 396,
		1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 79, 1, 0, 0, 0, 400, 401, 5, 18,
		0, 0, 401, 402, 3, 142, 71, 0, 402, 403, 5, 60, 0, 0, 403, 404, 3, 36,
		18, 0, 404, 405, 3, 82, 41, 0, 405, 410, 1, 0, 0, 0, 406, 407, 5, 25, 0,
		0, 407, 408, 5, 60, 0, 0, 408, 410, 3, 36, 18, 0, 409, 400, 1, 0, 0, 0,
		409, 406, 1, 0, 0, 0, 410, 81, 1, 0, 0, 0, 411, 414, 5, 19, 0, 0, 412,
		414, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 413, 412, 1, 0, 0, 0, 414, 83, 1,
		0, 0, 0, 415, 417, 5, 20, 0, 0, 416, 418, 3, 86, 43, 0, 417, 416, 1, 0,
		0, 0, 417, 418, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 3, 34, 17,
		0, 420, 85, 1, 0, 0, 0, 421, 423, 3, 44, 22, 0, 422, 421, 1, 0, 0, 0, 422,
		423, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 426, 5, 59, 0, 0, 425, 427,
		3, 124, 62, 0, 426, 425, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428, 1,
		0, 0, 0, 428, 430, 5, 59, 0, 0, 429, 431, 3, 44, 22, 0, 430, 429, 1, 0,
		0, 0, 430, 431, 1, 0, 0, 0, 431, 87, 1, 0, 0, 0, 432, 433, 5, 20, 0, 0,
		433, 434, 5, 21, 0, 0, 434, 435, 3, 90, 45, 0, 435, 436, 3, 34, 17, 0,
		436, 89, 1, 0, 0, 0, 437, 438, 3, 140, 70, 0, 438, 439, 5, 63, 0, 0, 439,
		441, 1, 0, 0, 0, 440, 437, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 442,
		1, 0, 0, 0, 442, 443, 3, 124, 62, 0, 443, 91, 1, 0, 0, 0, 444, 445, 5,
		12, 0, 0, 445, 447, 5, 54, 0, 0, 446, 448, 3, 94, 47, 0, 447, 446, 1, 0,
		0, 0, 447, 448, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 5, 55, 0, 0,
		450, 93, 1, 0, 0, 0, 451, 453, 3, 96, 48, 0, 452, 451, 1, 0, 0, 0, 453,
		454, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 95, 1,
		0, 0, 0, 456, 457, 5, 18, 0, 0, 457, 458, 3, 98, 49, 0, 458, 459, 5, 60,
		0, 0, 459, 460, 3, 36, 18, 0, 460, 465, 1, 0, 0, 0, 461, 462, 5, 25, 0,
		0, 462, 463, 5, 60, 0, 0, 463, 465, 3, 36, 18, 0, 464, 456, 1, 0, 0, 0,
		464, 461, 1, 0, 0, 0, 465, 97, 1, 0, 0, 0, 466, 469, 3, 56, 28, 0, 467,
		469, 3, 100, 50, 0, 468, 466, 1, 0, 0, 0, 468, 467, 1, 0, 0, 0, 469, 99,
		1, 0, 0, 0, 470, 471, 3, 142, 71, 0, 471, 472, 5, 63, 0, 0, 472, 473, 3,
		124, 62, 0, 473, 476, 1, 0, 0, 0, 474, 476, 3, 124, 62, 0, 475, 470, 1,
		0, 0, 0, 475, 474, 1, 0, 0, 0, 476, 101, 1, 0, 0, 0, 477, 486, 5, 28, 0,
		0, 478, 486, 3, 104, 52, 0, 479, 486, 3, 106, 53, 0, 480, 486, 3, 108,
		54, 0, 481, 486, 3, 110, 55, 0, 482, 486, 3, 112, 56, 0, 483, 486, 3, 118,
		59, 0, 484, 486, 5, 29, 0, 0, 485, 477, 1, 0, 0, 0, 485, 478, 1, 0, 0,
		0, 485, 479, 1, 0, 0, 0, 485, 480, 1, 0, 0, 0, 485, 481, 1, 0, 0, 0, 485,
		482, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 485, 484, 1, 0, 0, 0, 486, 103,
		1, 0, 0, 0, 487, 488, 5, 56, 0, 0, 488, 489, 5, 30, 0, 0, 489, 490, 5,
		57, 0, 0, 490, 491, 3, 102, 51, 0, 491, 105, 1, 0, 0, 0, 492, 493, 5, 56,
		0, 0, 493, 494, 5, 57, 0, 0, 494, 495, 3, 102, 51, 0, 495, 107, 1, 0, 0,
		0, 496, 497, 5, 9, 0, 0, 497, 498, 5, 56, 0, 0, 498, 499, 3, 102, 51, 0,
		499, 500, 5, 57, 0, 0, 500, 501, 3, 102, 51, 0, 501, 109, 1, 0, 0, 0, 502,
		503, 5, 10, 0, 0, 503, 504, 3, 102, 51, 0, 504, 111, 1, 0, 0, 0, 505, 506,
		5, 7, 0, 0, 506, 508, 5, 54, 0, 0, 507, 509, 3, 114, 57, 0, 508, 507, 1,
		0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 511, 5, 55, 0,
		0, 511, 113, 1, 0, 0, 0, 512, 517, 3, 116, 58, 0, 513, 514, 5, 59, 0, 0,
		514, 516, 3, 116, 58, 0, 515, 513, 1, 0, 0, 0, 516, 519, 1, 0, 0, 0, 517,
		515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517,
		1, 0, 0, 0, 520, 522, 5, 59, 0, 0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0,
		0, 0, 522, 115, 1, 0, 0, 0, 523, 524, 5, 29, 0, 0, 524, 525, 5, 60, 0,
		0, 525, 526, 3, 102, 51, 0, 526, 117, 1, 0, 0, 0, 527, 528, 5, 8, 0, 0,
		528, 530, 5, 54, 0, 0, 529, 531, 3, 120, 60, 0, 530, 529, 1, 0, 0, 0, 530,
		531, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 533, 5, 55, 0, 0, 533, 119,
		1, 0, 0, 0, 534, 539, 3, 122, 61, 0, 535, 536, 5, 59, 0, 0, 536, 538, 3,
		122, 61, 0, 537, 535, 1, 0, 0, 0, 538, 541, 1, 0, 0, 0, 539, 537, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 543, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0,
		542, 544, 5, 59, 0, 0, 543, 542, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544,
		121, 1, 0, 0, 0, 545, 546, 5, 29, 0, 0, 546, 547, 3, 26, 13, 0, 547, 123,
		1, 0, 0, 0, 548, 549, 6, 62, -1, 0, 549, 550, 3, 126, 63, 0, 550, 583,
		1, 0, 0, 0, 551, 552, 10, 11, 0, 0, 552, 553, 5, 75, 0, 0, 553, 582, 3,
		124, 62, 12, 554, 555, 10, 10, 0, 0, 555, 556, 5, 74, 0, 0, 556, 582, 3,
		124, 62, 11, 557, 558, 10, 9, 0, 0, 558, 559, 5, 72, 0, 0, 559, 582, 3,
		124, 62, 10, 560, 561, 10, 8, 0, 0, 561, 562, 5, 73, 0, 0, 562, 582, 3,
		124, 62, 9, 563, 564, 10, 7, 0, 0, 564, 565, 5, 71, 0, 0, 565, 582, 3,
		124, 62, 8, 566, 567, 10, 6, 0, 0, 567, 568, 7, 1, 0, 0, 568, 582, 3, 124,
		62, 7, 569, 570, 10, 5, 0, 0, 570, 571, 7, 2, 0, 0, 571, 582, 3, 124, 62,
		6, 572, 573, 10, 4, 0, 0, 573, 574, 7, 3, 0, 0, 574, 582, 3, 124, 62, 5,
		575, 576, 10, 3, 0, 0, 576, 577, 7, 4, 0, 0, 577, 582, 3, 124, 62, 4, 578,
		579, 10, 2, 0, 0, 579, 580, 7, 5, 0, 0, 580, 582, 3, 124, 62, 3, 581, 551,
		1, 0, 0, 0, 581, 554, 1, 0, 0, 0, 581, 557, 1, 0, 0, 0, 581, 560, 1, 0,
		0, 0, 581, 563, 1, 0, 0, 0, 581, 566, 1, 0, 0, 0, 581, 569, 1, 0, 0, 0,
		581, 572, 1, 0, 0, 0, 581, 575, 1, 0, 0, 0, 581, 578, 1, 0, 0, 0, 582,
		585, 1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584, 125,
		1, 0, 0, 0, 585, 583, 1, 0, 0, 0, 586, 600, 3, 128, 64, 0, 587, 588, 5,
		64, 0, 0, 588, 600, 3, 126, 63, 0, 589, 590, 5, 65, 0, 0, 590, 600, 3,
		126, 63, 0, 591, 592, 5, 76, 0, 0, 592, 600, 3, 126, 63, 0, 593, 594, 5,
		77, 0, 0, 594, 600, 3, 126, 63, 0, 595, 596, 5, 66, 0, 0, 596, 600, 3,
		126, 63, 0, 597, 598, 5, 71, 0, 0, 598, 600, 3, 126, 63, 0, 599, 586, 1,
		0, 0, 0, 599, 587, 1, 0, 0, 0, 599, 589, 1, 0, 0, 0, 599, 591, 1, 0, 0,
		0, 599, 593, 1, 0, 0, 0, 599, 595, 1, 0, 0, 0, 599, 597, 1, 0, 0, 0, 600,
		127, 1, 0, 0, 0, 601, 602, 6, 64, -1, 0, 602, 603, 3, 130, 65, 0, 603,
		612, 1, 0, 0, 0, 604, 605, 10, 3, 0, 0, 605, 611, 3, 132, 66, 0, 606, 607,
		10, 2, 0, 0, 607, 611, 3, 134, 67, 0, 608, 609, 10, 1, 0, 0, 609, 611,
		3, 136, 68, 0, 610, 604, 1, 0, 0, 0, 610, 606, 1, 0, 0, 0, 610, 608, 1,
		0, 0, 0, 611, 614, 1, 0, 0, 0, 612, 610, 1, 0, 0, 0, 612, 613, 1, 0, 0,
		0, 613, 129, 1, 0, 0, 0, 614, 612, 1, 0, 0, 0, 615, 629, 5, 29, 0, 0, 616,
		629, 5, 30, 0, 0, 617, 629, 5, 31, 0, 0, 618, 629, 5, 32, 0, 0, 619, 629,
		5, 34, 0, 0, 620, 629, 5, 33, 0, 0, 621, 622, 5, 52, 0, 0, 622, 623, 3,
		124, 62, 0, 623, 624, 5, 53, 0, 0, 624, 629, 1, 0, 0, 0, 625, 626, 5, 27,
		0, 0, 626, 627, 5, 52, 0, 0, 627, 629, 5, 53, 0, 0, 628, 615, 1, 0, 0,
		0, 628, 616, 1, 0, 0, 0, 628, 617, 1, 0, 0, 0, 628, 618, 1, 0, 0, 0, 628,
		619, 1, 0, 0, 0, 628, 620, 1, 0, 0, 0, 628, 621, 1, 0, 0, 0, 628, 625,
		1, 0, 0, 0, 629, 131, 1, 0, 0, 0, 630, 631, 5, 61, 0, 0, 631, 632, 5, 29,
		0, 0, 632, 133, 1, 0, 0, 0, 633, 634, 5, 56, 0, 0, 634, 635, 3, 124, 62,
		0, 635, 636, 5, 57, 0, 0, 636, 135, 1, 0, 0, 0, 637, 639, 5, 52, 0, 0,
		638, 640, 3, 138, 69, 0, 639, 638, 1, 0, 0, 0, 639, 640, 1, 0, 0, 0, 640,
		641, 1, 0, 0, 0, 641, 642, 5, 53, 0, 0, 642, 137, 1, 0, 0, 0, 643, 648,
		3, 124, 62, 0, 644, 645, 5, 58, 0, 0, 645, 647, 3, 124, 62, 0, 646, 644,
		1, 0, 0, 0, 647, 650, 1, 0, 0, 0, 648, 646, 1, 0, 0, 0, 648, 649, 1, 0,
		0, 0, 649, 139, 1, 0, 0, 0, 650, 648, 1, 0, 0, 0, 651, 656, 5, 29, 0, 0,
		652, 653, 5, 58, 0, 0, 653, 655, 5, 29, 0, 0, 654, 652, 1, 0, 0, 0, 655,
		658, 1, 0, 0, 0, 656, 654, 1, 0, 0, 0, 656, 657, 1, 0, 0, 0, 657, 141,
		1, 0, 0, 0, 658, 656, 1, 0, 0, 0, 659, 664, 3, 124, 62, 0, 660, 661, 5,
		58, 0, 0, 661, 663, 3, 124, 62, 0, 662, 660, 1, 0, 0, 0, 663, 666, 1, 0,
		0, 0, 664, 662, 1, 0, 0, 0, 664, 665, 1, 0, 0, 0, 665, 143, 1, 0, 0, 0,
		666, 664, 1, 0, 0, 0, 62, 146, 151, 162, 173, 178, 183, 190, 198, 202,
		208, 212, 220, 224, 233, 241, 245, 251, 260, 264, 271, 287, 294, 298, 314,
		319, 330, 355, 362, 366, 381, 385, 389, 398, 409, 413, 417, 422, 426, 430,
		440, 447, 454, 464, 468, 475, 485, 508, 517, 521, 530, 539, 543, 581, 583,
		599, 610, 612, 628, 639, 648, 656, 664,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JLangParserInit initializes any static state used to implement JLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JLangParserInit() {
	staticData := &JLangParserStaticData
	staticData.once.Do(jlangParserInit)
}

// NewJLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJLangParser(input antlr.TokenStream) *JLangParser {
	JLangParserInit()
	this := new(JLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JLang.g4"

	return this
}

// JLangParser tokens.
const (
	JLangParserEOF           = antlr.TokenEOF
	JLangParserKW_PKG        = 1
	JLangParserKW_IMP        = 2
	JLangParserKW_DEF        = 3
	JLangParserKW_VAR        = 4
	JLangParserKW_CONS       = 5
	JLangParserKW_TYPE       = 6
	JLangParserKW_STRUCT     = 7
	JLangParserKW_INTERFACE  = 8
	JLangParserKW_MAPPING    = 9
	JLangParserKW_CHANNEL    = 10
	JLangParserKW_J          = 11
	JLangParserKW_SELECT     = 12
	JLangParserKW_LATER      = 13
	JLangParserKW_RET        = 14
	JLangParserKW_IF         = 15
	JLangParserKW_ELSE       = 16
	JLangParserKW_SWITCH     = 17
	JLangParserKW_CASE       = 18
	JLangParserKW_FALL       = 19
	JLangParserKW_FR         = 20
	JLangParserKW_RANGE      = 21
	JLangParserKW_BREAK      = 22
	JLangParserKW_CONTINUE   = 23
	JLangParserKW_JOTO       = 24
	JLangParserKW_DFT        = 25
	JLangParserKW_PANIC      = 26
	JLangParserKW_RECOVER    = 27
	JLangParserTYPE_NAME     = 28
	JLangParserIDENT         = 29
	JLangParserINT_LIT       = 30
	JLangParserFLOAT_LIT     = 31
	JLangParserSTRING_LIT    = 32
	JLangParserRAW_STR       = 33
	JLangParserCHAR_LIT      = 34
	JLangParserCH_SEND       = 35
	JLangParserLE            = 36
	JLangParserGE            = 37
	JLangParserEQ            = 38
	JLangParserNE            = 39
	JLangParserSHLEQ         = 40
	JLangParserSHREQ         = 41
	JLangParserSHL           = 42
	JLangParserSHR           = 43
	JLangParserADDEQ         = 44
	JLangParserSUBEQ         = 45
	JLangParserMULEQ         = 46
	JLangParserDIVEQ         = 47
	JLangParserMODEQ         = 48
	JLangParserANDEQ         = 49
	JLangParserOREQ          = 50
	JLangParserXOREQ         = 51
	JLangParserLPAREN        = 52
	JLangParserRPAREN        = 53
	JLangParserLBRACE        = 54
	JLangParserRBRACE        = 55
	JLangParserLBRACK        = 56
	JLangParserRBRACK        = 57
	JLangParserCOMMA         = 58
	JLangParserSEMI          = 59
	JLangParserCOLON         = 60
	JLangParserDOT           = 61
	JLangParserASSIGN        = 62
	JLangParserDECL          = 63
	JLangParserPLUS          = 64
	JLangParserMINUS         = 65
	JLangParserSTAR          = 66
	JLangParserSLASH         = 67
	JLangParserPERCENT       = 68
	JLangParserLT            = 69
	JLangParserGT            = 70
	JLangParserBAND          = 71
	JLangParserBOR           = 72
	JLangParserBXOR          = 73
	JLangParserANDAND        = 74
	JLangParserOROR          = 75
	JLangParserBANG          = 76
	JLangParserTILDE         = 77
	JLangParserWS            = 78
	JLangParserLINE_COMMENT  = 79
	JLangParserBLOCK_COMMENT = 80
)

// JLangParser rules.
const (
	JLangParserRULE_program        = 0
	JLangParserRULE_pkgDecl        = 1
	JLangParserRULE_impList        = 2
	JLangParserRULE_impDecl        = 3
	JLangParserRULE_importSpec     = 4
	JLangParserRULE_topDecl        = 5
	JLangParserRULE_varDecl        = 6
	JLangParserRULE_varSpec        = 7
	JLangParserRULE_consDecl       = 8
	JLangParserRULE_constSpec      = 9
	JLangParserRULE_typeDecl       = 10
	JLangParserRULE_typeSpec       = 11
	JLangParserRULE_funcDecl       = 12
	JLangParserRULE_signature      = 13
	JLangParserRULE_paramList      = 14
	JLangParserRULE_param          = 15
	JLangParserRULE_resultType     = 16
	JLangParserRULE_block          = 17
	JLangParserRULE_stmtList       = 18
	JLangParserRULE_stmt           = 19
	JLangParserRULE_declStmt       = 20
	JLangParserRULE_labeledStmt    = 21
	JLangParserRULE_simpleStmt     = 22
	JLangParserRULE_exprStmt       = 23
	JLangParserRULE_assignStmt     = 24
	JLangParserRULE_lhs            = 25
	JLangParserRULE_assignOp       = 26
	JLangParserRULE_shortVarDecl   = 27
	JLangParserRULE_sendStmt       = 28
	JLangParserRULE_spawnStmt      = 29
	JLangParserRULE_returnStmt     = 30
	JLangParserRULE_deferStmt      = 31
	JLangParserRULE_breakStmt      = 32
	JLangParserRULE_continueStmt   = 33
	JLangParserRULE_jotoStmt       = 34
	JLangParserRULE_ifStmt         = 35
	JLangParserRULE_elseOpt        = 36
	JLangParserRULE_switchStmt     = 37
	JLangParserRULE_switchHead     = 38
	JLangParserRULE_caseClauses    = 39
	JLangParserRULE_caseClause     = 40
	JLangParserRULE_fallOpt        = 41
	JLangParserRULE_forStmt        = 42
	JLangParserRULE_forClause      = 43
	JLangParserRULE_rangeStmt      = 44
	JLangParserRULE_rangeHeader    = 45
	JLangParserRULE_selectStmt     = 46
	JLangParserRULE_commClauses    = 47
	JLangParserRULE_commClause     = 48
	JLangParserRULE_commClauseBody = 49
	JLangParserRULE_recvStmt       = 50
	JLangParserRULE_type_          = 51
	JLangParserRULE_arrayType      = 52
	JLangParserRULE_sliceType      = 53
	JLangParserRULE_mappingType    = 54
	JLangParserRULE_channelType    = 55
	JLangParserRULE_structType     = 56
	JLangParserRULE_typeFieldList  = 57
	JLangParserRULE_typeField      = 58
	JLangParserRULE_interfaceType  = 59
	JLangParserRULE_methodList     = 60
	JLangParserRULE_methodDecl     = 61
	JLangParserRULE_expr           = 62
	JLangParserRULE_unaryExpr      = 63
	JLangParserRULE_primaryExpr    = 64
	JLangParserRULE_operand        = 65
	JLangParserRULE_selector       = 66
	JLangParserRULE_index          = 67
	JLangParserRULE_call           = 68
	JLangParserRULE_argList        = 69
	JLangParserRULE_identList      = 70
	JLangParserRULE_exprList       = 71
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PkgDecl() IPkgDeclContext
	EOF() antlr.TerminalNode
	ImpList() IImpListContext
	AllTopDecl() []ITopDeclContext
	TopDecl(i int) ITopDeclContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PkgDecl() IPkgDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkgDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkgDeclContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(JLangParserEOF, 0)
}

func (s *ProgramContext) ImpList() IImpListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImpListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImpListContext)
}

func (s *ProgramContext) AllTopDecl() []ITopDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopDeclContext); ok {
			len++
		}
	}

	tst := make([]ITopDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopDeclContext); ok {
			tst[i] = t.(ITopDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TopDecl(i int) ITopDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *JLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.PkgDecl()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_IMP {
		{
			p.SetState(145)
			p.ImpList()
		}

	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&22518032124541048) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&12423) != 0) {
		{
			p.SetState(148)
			p.TopDecl()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Match(JLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPkgDeclContext is an interface to support dynamic dispatch.
type IPkgDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_PKG() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsPkgDeclContext differentiates from other interfaces.
	IsPkgDeclContext()
}

type PkgDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgDeclContext() *PkgDeclContext {
	var p = new(PkgDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_pkgDecl
	return p
}

func InitEmptyPkgDeclContext(p *PkgDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_pkgDecl
}

func (*PkgDeclContext) IsPkgDeclContext() {}

func NewPkgDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgDeclContext {
	var p = new(PkgDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_pkgDecl

	return p
}

func (s *PkgDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgDeclContext) KW_PKG() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_PKG, 0)
}

func (s *PkgDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *PkgDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterPkgDecl(s)
	}
}

func (s *PkgDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitPkgDecl(s)
	}
}

func (p *JLangParser) PkgDecl() (localctx IPkgDeclContext) {
	localctx = NewPkgDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JLangParserRULE_pkgDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(JLangParserKW_PKG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImpListContext is an interface to support dynamic dispatch.
type IImpListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllImpDecl() []IImpDeclContext
	ImpDecl(i int) IImpDeclContext

	// IsImpListContext differentiates from other interfaces.
	IsImpListContext()
}

type ImpListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpListContext() *ImpListContext {
	var p = new(ImpListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_impList
	return p
}

func InitEmptyImpListContext(p *ImpListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_impList
}

func (*ImpListContext) IsImpListContext() {}

func NewImpListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpListContext {
	var p = new(ImpListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_impList

	return p
}

func (s *ImpListContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpListContext) AllImpDecl() []IImpDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImpDeclContext); ok {
			len++
		}
	}

	tst := make([]IImpDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImpDeclContext); ok {
			tst[i] = t.(IImpDeclContext)
			i++
		}
	}

	return tst
}

func (s *ImpListContext) ImpDecl(i int) IImpDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImpDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImpDeclContext)
}

func (s *ImpListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterImpList(s)
	}
}

func (s *ImpListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitImpList(s)
	}
}

func (p *JLangParser) ImpList() (localctx IImpListContext) {
	localctx = NewImpListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JLangParserRULE_impList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_IMP {
		{
			p.SetState(159)
			p.ImpDecl()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImpDeclContext is an interface to support dynamic dispatch.
type IImpDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_IMP() antlr.TerminalNode
	AllImportSpec() []IImportSpecContext
	ImportSpec(i int) IImportSpecContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsImpDeclContext differentiates from other interfaces.
	IsImpDeclContext()
}

type ImpDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpDeclContext() *ImpDeclContext {
	var p = new(ImpDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_impDecl
	return p
}

func InitEmptyImpDeclContext(p *ImpDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_impDecl
}

func (*ImpDeclContext) IsImpDeclContext() {}

func NewImpDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpDeclContext {
	var p = new(ImpDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_impDecl

	return p
}

func (s *ImpDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpDeclContext) KW_IMP() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_IMP, 0)
}

func (s *ImpDeclContext) AllImportSpec() []IImportSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportSpecContext); ok {
			len++
		}
	}

	tst := make([]IImportSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportSpecContext); ok {
			tst[i] = t.(IImportSpecContext)
			i++
		}
	}

	return tst
}

func (s *ImpDeclContext) ImportSpec(i int) IImportSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImpDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserLPAREN, 0)
}

func (s *ImpDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserRPAREN, 0)
}

func (s *ImpDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JLangParserCOMMA)
}

func (s *ImpDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, i)
}

func (s *ImpDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterImpDecl(s)
	}
}

func (s *ImpDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitImpDecl(s)
	}
}

func (p *JLangParser) ImpDecl() (localctx IImpDeclContext) {
	localctx = NewImpDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JLangParserRULE_impDecl)
	var _la int

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(JLangParserKW_IMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.ImportSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(JLangParserKW_IMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.ImportSpec()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == JLangParserCOMMA {
			{
				p.SetState(169)
				p.Match(JLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(170)
				p.ImportSpec()
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(176)
			p.Match(JLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LIT() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_importSpec
	return p
}

func InitEmptyImportSpecContext(p *ImportSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_importSpec
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserSTRING_LIT, 0)
}

func (s *ImportSpecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (p *JLangParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JLangParserRULE_importSpec)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(JLangParserSTRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(JLangParserSTRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopDeclContext is an interface to support dynamic dispatch.
type ITopDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	ConsDecl() IConsDeclContext
	TypeDecl() ITypeDeclContext
	FuncDecl() IFuncDeclContext
	Stmt() IStmtContext

	// IsTopDeclContext differentiates from other interfaces.
	IsTopDeclContext()
}

type TopDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclContext() *TopDeclContext {
	var p = new(TopDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_topDecl
	return p
}

func InitEmptyTopDeclContext(p *TopDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_topDecl
}

func (*TopDeclContext) IsTopDeclContext() {}

func NewTopDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclContext {
	var p = new(TopDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_topDecl

	return p
}

func (s *TopDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *TopDeclContext) ConsDecl() IConsDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsDeclContext)
}

func (s *TopDeclContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TopDeclContext) FuncDecl() IFuncDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopDeclContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *TopDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterTopDecl(s)
	}
}

func (s *TopDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitTopDecl(s)
	}
}

func (p *JLangParser) TopDecl() (localctx ITopDeclContext) {
	localctx = NewTopDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JLangParserRULE_topDecl)
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.VarDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.ConsDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.TypeDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(188)
			p.FuncDecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(189)
			p.Stmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_VAR() antlr.TerminalNode
	AllVarSpec() []IVarSpecContext
	VarSpec(i int) IVarSpecContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) KW_VAR() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_VAR, 0)
}

func (s *VarDeclContext) AllVarSpec() []IVarSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarSpecContext); ok {
			len++
		}
	}

	tst := make([]IVarSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarSpecContext); ok {
			tst[i] = t.(IVarSpecContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarSpec(i int) IVarSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarSpecContext)
}

func (s *VarDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *VarDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *JLangParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JLangParserRULE_varDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(JLangParserKW_VAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.VarSpec()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(194)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(195)
				p.VarSpec()
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(201)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarSpecContext is an interface to support dynamic dispatch.
type IVarSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentList() IIdentListContext
	Type_() IType_Context
	ASSIGN() antlr.TerminalNode
	ExprList() IExprListContext

	// IsVarSpecContext differentiates from other interfaces.
	IsVarSpecContext()
}

type VarSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSpecContext() *VarSpecContext {
	var p = new(VarSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_varSpec
	return p
}

func InitEmptyVarSpecContext(p *VarSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_varSpec
}

func (*VarSpecContext) IsVarSpecContext() {}

func NewVarSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSpecContext {
	var p = new(VarSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_varSpec

	return p
}

func (s *VarSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSpecContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *VarSpecContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JLangParserASSIGN, 0)
}

func (s *VarSpecContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *VarSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterVarSpec(s)
	}
}

func (s *VarSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitVarSpec(s)
	}
}

func (p *JLangParser) VarSpec() (localctx IVarSpecContext) {
	localctx = NewVarSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JLangParserRULE_varSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.IdentList()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(205)
			p.Type_()
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JLangParserASSIGN {
			{
				p.SetState(206)
				p.Match(JLangParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(207)
				p.ExprList()
			}

		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(210)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.ExprList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConsDeclContext is an interface to support dynamic dispatch.
type IConsDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_CONS() antlr.TerminalNode
	AllConstSpec() []IConstSpecContext
	ConstSpec(i int) IConstSpecContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsConsDeclContext differentiates from other interfaces.
	IsConsDeclContext()
}

type ConsDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsDeclContext() *ConsDeclContext {
	var p = new(ConsDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_consDecl
	return p
}

func InitEmptyConsDeclContext(p *ConsDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_consDecl
}

func (*ConsDeclContext) IsConsDeclContext() {}

func NewConsDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsDeclContext {
	var p = new(ConsDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_consDecl

	return p
}

func (s *ConsDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsDeclContext) KW_CONS() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_CONS, 0)
}

func (s *ConsDeclContext) AllConstSpec() []IConstSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstSpecContext); ok {
			len++
		}
	}

	tst := make([]IConstSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstSpecContext); ok {
			tst[i] = t.(IConstSpecContext)
			i++
		}
	}

	return tst
}

func (s *ConsDeclContext) ConstSpec(i int) IConstSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstSpecContext)
}

func (s *ConsDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *ConsDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *ConsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterConsDecl(s)
	}
}

func (s *ConsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitConsDecl(s)
	}
}

func (p *JLangParser) ConsDecl() (localctx IConsDeclContext) {
	localctx = NewConsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JLangParserRULE_consDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(JLangParserKW_CONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.ConstSpec()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(216)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(217)
				p.ConstSpec()
			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(223)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstSpecContext is an interface to support dynamic dispatch.
type IConstSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentList() IIdentListContext
	Type_() IType_Context
	ASSIGN() antlr.TerminalNode
	ExprList() IExprListContext

	// IsConstSpecContext differentiates from other interfaces.
	IsConstSpecContext()
}

type ConstSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstSpecContext() *ConstSpecContext {
	var p = new(ConstSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_constSpec
	return p
}

func InitEmptyConstSpecContext(p *ConstSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_constSpec
}

func (*ConstSpecContext) IsConstSpecContext() {}

func NewConstSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstSpecContext {
	var p = new(ConstSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_constSpec

	return p
}

func (s *ConstSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstSpecContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *ConstSpecContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ConstSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JLangParserASSIGN, 0)
}

func (s *ConstSpecContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ConstSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterConstSpec(s)
	}
}

func (s *ConstSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitConstSpec(s)
	}
}

func (p *JLangParser) ConstSpec() (localctx IConstSpecContext) {
	localctx = NewConstSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JLangParserRULE_constSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.IdentList()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_STRUCT, JLangParserKW_INTERFACE, JLangParserKW_MAPPING, JLangParserKW_CHANNEL, JLangParserTYPE_NAME, JLangParserIDENT, JLangParserLBRACK:
		{
			p.SetState(227)
			p.Type_()
		}
		{
			p.SetState(228)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.ExprList()
		}

	case JLangParserASSIGN:
		{
			p.SetState(231)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.ExprList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_TYPE() antlr.TerminalNode
	AllTypeSpec() []ITypeSpecContext
	TypeSpec(i int) ITypeSpecContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_TYPE, 0)
}

func (s *TypeDeclContext) AllTypeSpec() []ITypeSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeSpecContext); ok {
			len++
		}
	}

	tst := make([]ITypeSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeSpecContext); ok {
			tst[i] = t.(ITypeSpecContext)
			i++
		}
	}

	return tst
}

func (s *TypeDeclContext) TypeSpec(i int) ITypeSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeDeclContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *TypeDeclContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *JLangParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JLangParserRULE_typeDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(JLangParserKW_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.TypeSpec()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(237)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)
				p.TypeSpec()
			}

		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Type_() IType_Context

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeSpec
	return p
}

func InitEmptyTypeSpecContext(p *TypeSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeSpec
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *TypeSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JLangParserASSIGN, 0)
}

func (s *TypeSpecContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (p *JLangParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JLangParserRULE_typeSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserASSIGN:
		{
			p.SetState(248)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Type_()
		}

	case JLangParserKW_STRUCT, JLangParserKW_INTERFACE, JLangParserKW_MAPPING, JLangParserKW_CHANNEL, JLangParserTYPE_NAME, JLangParserIDENT, JLangParserLBRACK:
		{
			p.SetState(250)
			p.Type_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_DEF() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	Signature() ISignatureContext
	Block() IBlockContext

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) KW_DEF() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_DEF, 0)
}

func (s *FuncDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *FuncDeclContext) Signature() ISignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignatureContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (p *JLangParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JLangParserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(JLangParserKW_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Signature()
	}
	{
		p.SetState(256)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignatureContext is an interface to support dynamic dispatch.
type ISignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ParamList() IParamListContext
	ResultType() IResultTypeContext

	// IsSignatureContext differentiates from other interfaces.
	IsSignatureContext()
}

type SignatureContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignatureContext() *SignatureContext {
	var p = new(SignatureContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_signature
	return p
}

func InitEmptySignatureContext(p *SignatureContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_signature
}

func (*SignatureContext) IsSignatureContext() {}

func NewSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignatureContext {
	var p = new(SignatureContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_signature

	return p
}

func (s *SignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *SignatureContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserLPAREN, 0)
}

func (s *SignatureContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserRPAREN, 0)
}

func (s *SignatureContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *SignatureContext) ResultType() IResultTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultTypeContext)
}

func (s *SignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSignature(s)
	}
}

func (s *SignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSignature(s)
	}
}

func (p *JLangParser) Signature() (localctx ISignatureContext) {
	localctx = NewSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JLangParserRULE_signature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(JLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(259)
			p.ParamList()
		}

	}
	{
		p.SetState(262)
		p.Match(JLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserCOLON {
		{
			p.SetState(263)
			p.ResultType()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JLangParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *JLangParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JLangParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Param()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(267)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Param()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() IType_Context

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *ParamContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *JLangParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JLangParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultTypeContext is an interface to support dynamic dispatch.
type IResultTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	Type_() IType_Context

	// IsResultTypeContext differentiates from other interfaces.
	IsResultTypeContext()
}

type ResultTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultTypeContext() *ResultTypeContext {
	var p = new(ResultTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_resultType
	return p
}

func InitEmptyResultTypeContext(p *ResultTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_resultType
}

func (*ResultTypeContext) IsResultTypeContext() {}

func NewResultTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultTypeContext {
	var p = new(ResultTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_resultType

	return p
}

func (s *ResultTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *ResultTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ResultTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterResultType(s)
	}
}

func (s *ResultTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitResultType(s)
	}
}

func (p *JLangParser) ResultType() (localctx IResultTypeContext) {
	localctx = NewResultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JLangParserRULE_resultType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	StmtList() IStmtListContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACE, 0)
}

func (s *BlockContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *JLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JLangParserRULE_block)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(JLangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(JLangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Match(JLangParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.StmtList()
		}
		{
			p.SetState(285)
			p.Match(JLangParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtListContext is an interface to support dynamic dispatch.
type IStmtListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsStmtListContext differentiates from other interfaces.
	IsStmtListContext()
}

type StmtListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtListContext() *StmtListContext {
	var p = new(StmtListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_stmtList
	return p
}

func InitEmptyStmtListContext(p *StmtListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_stmtList
}

func (*StmtListContext) IsStmtListContext() {}

func NewStmtListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtListContext {
	var p = new(StmtListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_stmtList

	return p
}

func (s *StmtListContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtListContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtListContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *StmtListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *StmtListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterStmtList(s)
	}
}

func (s *StmtListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitStmtList(s)
	}
}

func (p *JLangParser) StmtList() (localctx IStmtListContext) {
	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JLangParserRULE_stmtList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Stmt()
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(291)
				p.Stmt()
			}

		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(297)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleStmt() ISimpleStmtContext
	DeclStmt() IDeclStmtContext
	IfStmt() IIfStmtContext
	SwitchStmt() ISwitchStmtContext
	ForStmt() IForStmtContext
	RangeStmt() IRangeStmtContext
	SelectStmt() ISelectStmtContext
	DeferStmt() IDeferStmtContext
	ReturnStmt() IReturnStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	JotoStmt() IJotoStmtContext
	LabeledStmt() ILabeledStmtContext
	Block() IBlockContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StmtContext) DeclStmt() IDeclStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StmtContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtContext) RangeStmt() IRangeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeStmtContext)
}

func (s *StmtContext) SelectStmt() ISelectStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStmtContext)
}

func (s *StmtContext) DeferStmt() IDeferStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeferStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeferStmtContext)
}

func (s *StmtContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StmtContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StmtContext) JotoStmt() IJotoStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJotoStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJotoStmtContext)
}

func (s *StmtContext) LabeledStmt() ILabeledStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabeledStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabeledStmtContext)
}

func (s *StmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *JLangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JLangParserRULE_stmt)
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.SimpleStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.DeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(302)
			p.IfStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(303)
			p.SwitchStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(304)
			p.ForStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(305)
			p.RangeStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(306)
			p.SelectStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(307)
			p.DeferStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(308)
			p.ReturnStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(309)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(310)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(311)
			p.JotoStmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(312)
			p.LabeledStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(313)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclStmtContext is an interface to support dynamic dispatch.
type IDeclStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDecl() IVarDeclContext
	ConsDecl() IConsDeclContext
	TypeDecl() ITypeDeclContext

	// IsDeclStmtContext differentiates from other interfaces.
	IsDeclStmtContext()
}

type DeclStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclStmtContext() *DeclStmtContext {
	var p = new(DeclStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_declStmt
	return p
}

func InitEmptyDeclStmtContext(p *DeclStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_declStmt
}

func (*DeclStmtContext) IsDeclStmtContext() {}

func NewDeclStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclStmtContext {
	var p = new(DeclStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_declStmt

	return p
}

func (s *DeclStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclStmtContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *DeclStmtContext) ConsDecl() IConsDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsDeclContext)
}

func (s *DeclStmtContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterDeclStmt(s)
	}
}

func (s *DeclStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitDeclStmt(s)
	}
}

func (p *JLangParser) DeclStmt() (localctx IDeclStmtContext) {
	localctx = NewDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JLangParserRULE_declStmt)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_VAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.VarDecl()
		}

	case JLangParserKW_CONS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.ConsDecl()
		}

	case JLangParserKW_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
			p.TypeDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabeledStmtContext is an interface to support dynamic dispatch.
type ILabeledStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Stmt() IStmtContext

	// IsLabeledStmtContext differentiates from other interfaces.
	IsLabeledStmtContext()
}

type LabeledStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeledStmtContext() *LabeledStmtContext {
	var p = new(LabeledStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_labeledStmt
	return p
}

func InitEmptyLabeledStmtContext(p *LabeledStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_labeledStmt
}

func (*LabeledStmtContext) IsLabeledStmtContext() {}

func NewLabeledStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabeledStmtContext {
	var p = new(LabeledStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_labeledStmt

	return p
}

func (s *LabeledStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LabeledStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *LabeledStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *LabeledStmtContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *LabeledStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabeledStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabeledStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterLabeledStmt(s)
	}
}

func (s *LabeledStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitLabeledStmt(s)
	}
}

func (p *JLangParser) LabeledStmt() (localctx ILabeledStmtContext) {
	localctx = NewLabeledStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JLangParserRULE_labeledStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Stmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprStmt() IExprStmtContext
	AssignStmt() IAssignStmtContext
	ShortVarDecl() IShortVarDeclContext
	SendStmt() ISendStmtContext
	SpawnStmt() ISpawnStmtContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *SimpleStmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *SimpleStmtContext) ShortVarDecl() IShortVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShortVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShortVarDeclContext)
}

func (s *SimpleStmtContext) SendStmt() ISendStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendStmtContext)
}

func (s *SimpleStmtContext) SpawnStmt() ISpawnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpawnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpawnStmtContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (p *JLangParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JLangParserRULE_simpleStmt)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.ExprStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.AssignStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.ShortVarDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(328)
			p.SendStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)
			p.SpawnStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (p *JLangParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JLangParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lhs() ILhsContext
	AssignOp() IAssignOpContext
	ExprList() IExprListContext

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_assignStmt
	return p
}

func InitEmptyAssignStmtContext(p *AssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_assignStmt
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) Lhs() ILhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *AssignStmtContext) AssignOp() IAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignOpContext)
}

func (s *AssignStmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (p *JLangParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JLangParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Lhs()
	}
	{
		p.SetState(335)
		p.AssignOp()
	}
	{
		p.SetState(336)
		p.ExprList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext

	// IsLhsContext differentiates from other interfaces.
	IsLhsContext()
}

type LhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsContext() *LhsContext {
	var p = new(LhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_lhs
	return p
}

func InitEmptyLhsContext(p *LhsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_lhs
}

func (*LhsContext) IsLhsContext() {}

func NewLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsContext {
	var p = new(LhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_lhs

	return p
}

func (s *LhsContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *LhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterLhs(s)
	}
}

func (s *LhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitLhs(s)
	}
}

func (p *JLangParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, JLangParserRULE_lhs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.primaryExpr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignOpContext is an interface to support dynamic dispatch.
type IAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	ADDEQ() antlr.TerminalNode
	SUBEQ() antlr.TerminalNode
	MULEQ() antlr.TerminalNode
	DIVEQ() antlr.TerminalNode
	MODEQ() antlr.TerminalNode
	ANDEQ() antlr.TerminalNode
	OREQ() antlr.TerminalNode
	XOREQ() antlr.TerminalNode
	SHLEQ() antlr.TerminalNode
	SHREQ() antlr.TerminalNode

	// IsAssignOpContext differentiates from other interfaces.
	IsAssignOpContext()
}

type AssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOpContext() *AssignOpContext {
	var p = new(AssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_assignOp
	return p
}

func InitEmptyAssignOpContext(p *AssignOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_assignOp
}

func (*AssignOpContext) IsAssignOpContext() {}

func NewAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOpContext {
	var p = new(AssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_assignOp

	return p
}

func (s *AssignOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOpContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JLangParserASSIGN, 0)
}

func (s *AssignOpContext) ADDEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserADDEQ, 0)
}

func (s *AssignOpContext) SUBEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserSUBEQ, 0)
}

func (s *AssignOpContext) MULEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserMULEQ, 0)
}

func (s *AssignOpContext) DIVEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserDIVEQ, 0)
}

func (s *AssignOpContext) MODEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserMODEQ, 0)
}

func (s *AssignOpContext) ANDEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserANDEQ, 0)
}

func (s *AssignOpContext) OREQ() antlr.TerminalNode {
	return s.GetToken(JLangParserOREQ, 0)
}

func (s *AssignOpContext) XOREQ() antlr.TerminalNode {
	return s.GetToken(JLangParserXOREQ, 0)
}

func (s *AssignOpContext) SHLEQ() antlr.TerminalNode {
	return s.GetToken(JLangParserSHLEQ, 0)
}

func (s *AssignOpContext) SHREQ() antlr.TerminalNode {
	return s.GetToken(JLangParserSHREQ, 0)
}

func (s *AssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterAssignOp(s)
	}
}

func (s *AssignOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitAssignOp(s)
	}
}

func (p *JLangParser) AssignOp() (localctx IAssignOpContext) {
	localctx = NewAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, JLangParserRULE_assignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4616175324403597312) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShortVarDeclContext is an interface to support dynamic dispatch.
type IShortVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentList() IIdentListContext
	DECL() antlr.TerminalNode
	ExprList() IExprListContext

	// IsShortVarDeclContext differentiates from other interfaces.
	IsShortVarDeclContext()
}

type ShortVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortVarDeclContext() *ShortVarDeclContext {
	var p = new(ShortVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_shortVarDecl
	return p
}

func InitEmptyShortVarDeclContext(p *ShortVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_shortVarDecl
}

func (*ShortVarDeclContext) IsShortVarDeclContext() {}

func NewShortVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShortVarDeclContext {
	var p = new(ShortVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_shortVarDecl

	return p
}

func (s *ShortVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ShortVarDeclContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *ShortVarDeclContext) DECL() antlr.TerminalNode {
	return s.GetToken(JLangParserDECL, 0)
}

func (s *ShortVarDeclContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ShortVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShortVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShortVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterShortVarDecl(s)
	}
}

func (s *ShortVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitShortVarDecl(s)
	}
}

func (p *JLangParser) ShortVarDecl() (localctx IShortVarDeclContext) {
	localctx = NewShortVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, JLangParserRULE_shortVarDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.IdentList()
	}
	{
		p.SetState(343)
		p.Match(JLangParserDECL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.ExprList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISendStmtContext is an interface to support dynamic dispatch.
type ISendStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CH_SEND() antlr.TerminalNode

	// IsSendStmtContext differentiates from other interfaces.
	IsSendStmtContext()
}

type SendStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendStmtContext() *SendStmtContext {
	var p = new(SendStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_sendStmt
	return p
}

func InitEmptySendStmtContext(p *SendStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_sendStmt
}

func (*SendStmtContext) IsSendStmtContext() {}

func NewSendStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendStmtContext {
	var p = new(SendStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_sendStmt

	return p
}

func (s *SendStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SendStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SendStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SendStmtContext) CH_SEND() antlr.TerminalNode {
	return s.GetToken(JLangParserCH_SEND, 0)
}

func (s *SendStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSendStmt(s)
	}
}

func (s *SendStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSendStmt(s)
	}
}

func (p *JLangParser) SendStmt() (localctx ISendStmtContext) {
	localctx = NewSendStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, JLangParserRULE_sendStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.expr(0)
	}
	{
		p.SetState(347)
		p.Match(JLangParserCH_SEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpawnStmtContext is an interface to support dynamic dispatch.
type ISpawnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_J() antlr.TerminalNode
	Expr() IExprContext

	// IsSpawnStmtContext differentiates from other interfaces.
	IsSpawnStmtContext()
}

type SpawnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpawnStmtContext() *SpawnStmtContext {
	var p = new(SpawnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_spawnStmt
	return p
}

func InitEmptySpawnStmtContext(p *SpawnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_spawnStmt
}

func (*SpawnStmtContext) IsSpawnStmtContext() {}

func NewSpawnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpawnStmtContext {
	var p = new(SpawnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_spawnStmt

	return p
}

func (s *SpawnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SpawnStmtContext) KW_J() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_J, 0)
}

func (s *SpawnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SpawnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpawnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpawnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSpawnStmt(s)
	}
}

func (s *SpawnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSpawnStmt(s)
	}
}

func (p *JLangParser) SpawnStmt() (localctx ISpawnStmtContext) {
	localctx = NewSpawnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, JLangParserRULE_spawnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(JLangParserKW_J)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_RET() antlr.TerminalNode
	ExprList() IExprListContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) KW_RET() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_RET, 0)
}

func (s *ReturnStmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (p *JLangParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, JLangParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(JLangParserKW_RET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(354)
			p.ExprList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeferStmtContext is an interface to support dynamic dispatch.
type IDeferStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_LATER() antlr.TerminalNode
	Expr() IExprContext

	// IsDeferStmtContext differentiates from other interfaces.
	IsDeferStmtContext()
}

type DeferStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeferStmtContext() *DeferStmtContext {
	var p = new(DeferStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_deferStmt
	return p
}

func InitEmptyDeferStmtContext(p *DeferStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_deferStmt
}

func (*DeferStmtContext) IsDeferStmtContext() {}

func NewDeferStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeferStmtContext {
	var p = new(DeferStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_deferStmt

	return p
}

func (s *DeferStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeferStmtContext) KW_LATER() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_LATER, 0)
}

func (s *DeferStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeferStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeferStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeferStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterDeferStmt(s)
	}
}

func (s *DeferStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitDeferStmt(s)
	}
}

func (p *JLangParser) DeferStmt() (localctx IDeferStmtContext) {
	localctx = NewDeferStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, JLangParserRULE_deferStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(JLangParserKW_LATER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_BREAK() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) KW_BREAK() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_BREAK, 0)
}

func (s *BreakStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (p *JLangParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, JLangParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(JLangParserKW_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_CONTINUE() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) KW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_CONTINUE, 0)
}

func (s *ContinueStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (p *JLangParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, JLangParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(JLangParserKW_CONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJotoStmtContext is an interface to support dynamic dispatch.
type IJotoStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_JOTO() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsJotoStmtContext differentiates from other interfaces.
	IsJotoStmtContext()
}

type JotoStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJotoStmtContext() *JotoStmtContext {
	var p = new(JotoStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_jotoStmt
	return p
}

func InitEmptyJotoStmtContext(p *JotoStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_jotoStmt
}

func (*JotoStmtContext) IsJotoStmtContext() {}

func NewJotoStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JotoStmtContext {
	var p = new(JotoStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_jotoStmt

	return p
}

func (s *JotoStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *JotoStmtContext) KW_JOTO() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_JOTO, 0)
}

func (s *JotoStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *JotoStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JotoStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JotoStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterJotoStmt(s)
	}
}

func (s *JotoStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitJotoStmt(s)
	}
}

func (p *JLangParser) JotoStmt() (localctx IJotoStmtContext) {
	localctx = NewJotoStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, JLangParserRULE_jotoStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(JLangParserKW_JOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_IF() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext
	ElseOpt() IElseOptContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_IF, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ElseOpt() IElseOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseOptContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *JLangParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, JLangParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(JLangParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.expr(0)
	}
	{
		p.SetState(373)
		p.Block()
	}
	{
		p.SetState(374)
		p.ElseOpt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseOptContext is an interface to support dynamic dispatch.
type IElseOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_ELSE() antlr.TerminalNode
	IfStmt() IIfStmtContext
	Block() IBlockContext

	// IsElseOptContext differentiates from other interfaces.
	IsElseOptContext()
}

type ElseOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseOptContext() *ElseOptContext {
	var p = new(ElseOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_elseOpt
	return p
}

func InitEmptyElseOptContext(p *ElseOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_elseOpt
}

func (*ElseOptContext) IsElseOptContext() {}

func NewElseOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseOptContext {
	var p = new(ElseOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_elseOpt

	return p
}

func (s *ElseOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseOptContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_ELSE, 0)
}

func (s *ElseOptContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *ElseOptContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterElseOpt(s)
	}
}

func (s *ElseOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitElseOpt(s)
	}
}

func (p *JLangParser) ElseOpt() (localctx IElseOptContext) {
	localctx = NewElseOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, JLangParserRULE_elseOpt)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)
			p.Match(JLangParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.IfStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)
			p.Match(JLangParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SWITCH() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	SwitchHead() ISwitchHeadContext
	CaseClauses() ICaseClausesContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) KW_SWITCH() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_SWITCH, 0)
}

func (s *SwitchStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACE, 0)
}

func (s *SwitchStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACE, 0)
}

func (s *SwitchStmtContext) SwitchHead() ISwitchHeadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchHeadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchHeadContext)
}

func (s *SwitchStmtContext) CaseClauses() ICaseClausesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseClausesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseClausesContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (p *JLangParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, JLangParserRULE_switchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(JLangParserKW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&1707404152537341) != 0 {
		{
			p.SetState(384)
			p.SwitchHead()
		}

	}
	{
		p.SetState(387)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(388)
			p.CaseClauses()
		}

	}
	{
		p.SetState(391)
		p.Match(JLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchHeadContext is an interface to support dynamic dispatch.
type ISwitchHeadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsSwitchHeadContext differentiates from other interfaces.
	IsSwitchHeadContext()
}

type SwitchHeadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchHeadContext() *SwitchHeadContext {
	var p = new(SwitchHeadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_switchHead
	return p
}

func InitEmptySwitchHeadContext(p *SwitchHeadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_switchHead
}

func (*SwitchHeadContext) IsSwitchHeadContext() {}

func NewSwitchHeadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchHeadContext {
	var p = new(SwitchHeadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_switchHead

	return p
}

func (s *SwitchHeadContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchHeadContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchHeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchHeadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchHeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSwitchHead(s)
	}
}

func (s *SwitchHeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSwitchHead(s)
	}
}

func (p *JLangParser) SwitchHead() (localctx ISwitchHeadContext) {
	localctx = NewSwitchHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, JLangParserRULE_switchHead)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseClausesContext is an interface to support dynamic dispatch.
type ICaseClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCaseClause() []ICaseClauseContext
	CaseClause(i int) ICaseClauseContext

	// IsCaseClausesContext differentiates from other interfaces.
	IsCaseClausesContext()
}

type CaseClausesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClausesContext() *CaseClausesContext {
	var p = new(CaseClausesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_caseClauses
	return p
}

func InitEmptyCaseClausesContext(p *CaseClausesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_caseClauses
}

func (*CaseClausesContext) IsCaseClausesContext() {}

func NewCaseClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClausesContext {
	var p = new(CaseClausesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_caseClauses

	return p
}

func (s *CaseClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClausesContext) AllCaseClause() []ICaseClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseClauseContext); ok {
			len++
		}
	}

	tst := make([]ICaseClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseClauseContext); ok {
			tst[i] = t.(ICaseClauseContext)
			i++
		}
	}

	return tst
}

func (s *CaseClausesContext) CaseClause(i int) ICaseClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseClauseContext)
}

func (s *CaseClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCaseClauses(s)
	}
}

func (s *CaseClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCaseClauses(s)
	}
}

func (p *JLangParser) CaseClauses() (localctx ICaseClausesContext) {
	localctx = NewCaseClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, JLangParserRULE_caseClauses)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(395)
			p.CaseClause()
		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseClauseContext is an interface to support dynamic dispatch.
type ICaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_CASE() antlr.TerminalNode
	ExprList() IExprListContext
	COLON() antlr.TerminalNode
	StmtList() IStmtListContext
	FallOpt() IFallOptContext
	KW_DFT() antlr.TerminalNode

	// IsCaseClauseContext differentiates from other interfaces.
	IsCaseClauseContext()
}

type CaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClauseContext() *CaseClauseContext {
	var p = new(CaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_caseClause
	return p
}

func InitEmptyCaseClauseContext(p *CaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_caseClause
}

func (*CaseClauseContext) IsCaseClauseContext() {}

func NewCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClauseContext {
	var p = new(CaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_caseClause

	return p
}

func (s *CaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClauseContext) KW_CASE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_CASE, 0)
}

func (s *CaseClauseContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *CaseClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *CaseClauseContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *CaseClauseContext) FallOpt() IFallOptContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFallOptContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFallOptContext)
}

func (s *CaseClauseContext) KW_DFT() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_DFT, 0)
}

func (s *CaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCaseClause(s)
	}
}

func (s *CaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCaseClause(s)
	}
}

func (p *JLangParser) CaseClause() (localctx ICaseClauseContext) {
	localctx = NewCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, JLangParserRULE_caseClause)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_CASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.Match(JLangParserKW_CASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.ExprList()
		}
		{
			p.SetState(402)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.StmtList()
		}
		{
			p.SetState(404)
			p.FallOpt()
		}

	case JLangParserKW_DFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Match(JLangParserKW_DFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.StmtList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFallOptContext is an interface to support dynamic dispatch.
type IFallOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FALL() antlr.TerminalNode

	// IsFallOptContext differentiates from other interfaces.
	IsFallOptContext()
}

type FallOptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFallOptContext() *FallOptContext {
	var p = new(FallOptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_fallOpt
	return p
}

func InitEmptyFallOptContext(p *FallOptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_fallOpt
}

func (*FallOptContext) IsFallOptContext() {}

func NewFallOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FallOptContext {
	var p = new(FallOptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_fallOpt

	return p
}

func (s *FallOptContext) GetParser() antlr.Parser { return s.parser }

func (s *FallOptContext) KW_FALL() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_FALL, 0)
}

func (s *FallOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FallOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FallOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterFallOpt(s)
	}
}

func (s *FallOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitFallOpt(s)
	}
}

func (p *JLangParser) FallOpt() (localctx IFallOptContext) {
	localctx = NewFallOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, JLangParserRULE_fallOpt)
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_FALL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(JLangParserKW_FALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserKW_CASE, JLangParserKW_DFT, JLangParserRBRACE:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FR() antlr.TerminalNode
	Block() IBlockContext
	ForClause() IForClauseContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) KW_FR() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_FR, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) ForClause() IForClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForClauseContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (p *JLangParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, JLangParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Match(JLangParserKW_FR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&580964385887881216) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&12423) != 0) {
		{
			p.SetState(416)
			p.ForClause()
		}

	}
	{
		p.SetState(419)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForClauseContext is an interface to support dynamic dispatch.
type IForClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	AllSimpleStmt() []ISimpleStmtContext
	SimpleStmt(i int) ISimpleStmtContext
	Expr() IExprContext

	// IsForClauseContext differentiates from other interfaces.
	IsForClauseContext()
}

type ForClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForClauseContext() *ForClauseContext {
	var p = new(ForClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_forClause
	return p
}

func InitEmptyForClauseContext(p *ForClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_forClause
}

func (*ForClauseContext) IsForClauseContext() {}

func NewForClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForClauseContext {
	var p = new(ForClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_forClause

	return p
}

func (s *ForClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ForClauseContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *ForClauseContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *ForClauseContext) AllSimpleStmt() []ISimpleStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStmtContext); ok {
			tst[i] = t.(ISimpleStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClauseContext) SimpleStmt(i int) ISimpleStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *ForClauseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterForClause(s)
	}
}

func (s *ForClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitForClause(s)
	}
}

func (p *JLangParser) ForClause() (localctx IForClauseContext) {
	localctx = NewForClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, JLangParserRULE_forClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503633584457728) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&12423) != 0) {
		{
			p.SetState(421)
			p.SimpleStmt()
		}

	}
	{
		p.SetState(424)
		p.Match(JLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&1707404152537341) != 0 {
		{
			p.SetState(425)
			p.expr(0)
		}

	}
	{
		p.SetState(428)
		p.Match(JLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503633584457728) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&12423) != 0) {
		{
			p.SetState(429)
			p.SimpleStmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeStmtContext is an interface to support dynamic dispatch.
type IRangeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FR() antlr.TerminalNode
	KW_RANGE() antlr.TerminalNode
	RangeHeader() IRangeHeaderContext
	Block() IBlockContext

	// IsRangeStmtContext differentiates from other interfaces.
	IsRangeStmtContext()
}

type RangeStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeStmtContext() *RangeStmtContext {
	var p = new(RangeStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_rangeStmt
	return p
}

func InitEmptyRangeStmtContext(p *RangeStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_rangeStmt
}

func (*RangeStmtContext) IsRangeStmtContext() {}

func NewRangeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeStmtContext {
	var p = new(RangeStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_rangeStmt

	return p
}

func (s *RangeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeStmtContext) KW_FR() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_FR, 0)
}

func (s *RangeStmtContext) KW_RANGE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_RANGE, 0)
}

func (s *RangeStmtContext) RangeHeader() IRangeHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeHeaderContext)
}

func (s *RangeStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RangeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterRangeStmt(s)
	}
}

func (s *RangeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitRangeStmt(s)
	}
}

func (p *JLangParser) RangeStmt() (localctx IRangeStmtContext) {
	localctx = NewRangeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, JLangParserRULE_rangeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(JLangParserKW_FR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(JLangParserKW_RANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.RangeHeader()
	}
	{
		p.SetState(435)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeHeaderContext is an interface to support dynamic dispatch.
type IRangeHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	IdentList() IIdentListContext
	DECL() antlr.TerminalNode

	// IsRangeHeaderContext differentiates from other interfaces.
	IsRangeHeaderContext()
}

type RangeHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeHeaderContext() *RangeHeaderContext {
	var p = new(RangeHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_rangeHeader
	return p
}

func InitEmptyRangeHeaderContext(p *RangeHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_rangeHeader
}

func (*RangeHeaderContext) IsRangeHeaderContext() {}

func NewRangeHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeHeaderContext {
	var p = new(RangeHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_rangeHeader

	return p
}

func (s *RangeHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeHeaderContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangeHeaderContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *RangeHeaderContext) DECL() antlr.TerminalNode {
	return s.GetToken(JLangParserDECL, 0)
}

func (s *RangeHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterRangeHeader(s)
	}
}

func (s *RangeHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitRangeHeader(s)
	}
}

func (p *JLangParser) RangeHeader() (localctx IRangeHeaderContext) {
	localctx = NewRangeHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, JLangParserRULE_rangeHeader)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(437)
			p.IdentList()
		}
		{
			p.SetState(438)
			p.Match(JLangParserDECL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(442)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectStmtContext is an interface to support dynamic dispatch.
type ISelectStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SELECT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	CommClauses() ICommClausesContext

	// IsSelectStmtContext differentiates from other interfaces.
	IsSelectStmtContext()
}

type SelectStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStmtContext() *SelectStmtContext {
	var p = new(SelectStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_selectStmt
	return p
}

func InitEmptySelectStmtContext(p *SelectStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_selectStmt
}

func (*SelectStmtContext) IsSelectStmtContext() {}

func NewSelectStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStmtContext {
	var p = new(SelectStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_selectStmt

	return p
}

func (s *SelectStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStmtContext) KW_SELECT() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_SELECT, 0)
}

func (s *SelectStmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACE, 0)
}

func (s *SelectStmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACE, 0)
}

func (s *SelectStmtContext) CommClauses() ICommClausesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommClausesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommClausesContext)
}

func (s *SelectStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSelectStmt(s)
	}
}

func (s *SelectStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSelectStmt(s)
	}
}

func (p *JLangParser) SelectStmt() (localctx ISelectStmtContext) {
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, JLangParserRULE_selectStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(JLangParserKW_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(445)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(446)
			p.CommClauses()
		}

	}
	{
		p.SetState(449)
		p.Match(JLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommClausesContext is an interface to support dynamic dispatch.
type ICommClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommClause() []ICommClauseContext
	CommClause(i int) ICommClauseContext

	// IsCommClausesContext differentiates from other interfaces.
	IsCommClausesContext()
}

type CommClausesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommClausesContext() *CommClausesContext {
	var p = new(CommClausesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClauses
	return p
}

func InitEmptyCommClausesContext(p *CommClausesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClauses
}

func (*CommClausesContext) IsCommClausesContext() {}

func NewCommClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommClausesContext {
	var p = new(CommClausesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_commClauses

	return p
}

func (s *CommClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *CommClausesContext) AllCommClause() []ICommClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommClauseContext); ok {
			len++
		}
	}

	tst := make([]ICommClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommClauseContext); ok {
			tst[i] = t.(ICommClauseContext)
			i++
		}
	}

	return tst
}

func (s *CommClausesContext) CommClause(i int) ICommClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommClauseContext)
}

func (s *CommClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCommClauses(s)
	}
}

func (s *CommClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCommClauses(s)
	}
}

func (p *JLangParser) CommClauses() (localctx ICommClausesContext) {
	localctx = NewCommClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, JLangParserRULE_commClauses)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(451)
			p.CommClause()
		}

		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommClauseContext is an interface to support dynamic dispatch.
type ICommClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_CASE() antlr.TerminalNode
	CommClauseBody() ICommClauseBodyContext
	COLON() antlr.TerminalNode
	StmtList() IStmtListContext
	KW_DFT() antlr.TerminalNode

	// IsCommClauseContext differentiates from other interfaces.
	IsCommClauseContext()
}

type CommClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommClauseContext() *CommClauseContext {
	var p = new(CommClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClause
	return p
}

func InitEmptyCommClauseContext(p *CommClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClause
}

func (*CommClauseContext) IsCommClauseContext() {}

func NewCommClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommClauseContext {
	var p = new(CommClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_commClause

	return p
}

func (s *CommClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CommClauseContext) KW_CASE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_CASE, 0)
}

func (s *CommClauseContext) CommClauseBody() ICommClauseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommClauseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommClauseBodyContext)
}

func (s *CommClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *CommClauseContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *CommClauseContext) KW_DFT() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_DFT, 0)
}

func (s *CommClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCommClause(s)
	}
}

func (s *CommClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCommClause(s)
	}
}

func (p *JLangParser) CommClause() (localctx ICommClauseContext) {
	localctx = NewCommClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, JLangParserRULE_commClause)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_CASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.Match(JLangParserKW_CASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.CommClauseBody()
		}
		{
			p.SetState(458)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.StmtList()
		}

	case JLangParserKW_DFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.Match(JLangParserKW_DFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.StmtList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommClauseBodyContext is an interface to support dynamic dispatch.
type ICommClauseBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SendStmt() ISendStmtContext
	RecvStmt() IRecvStmtContext

	// IsCommClauseBodyContext differentiates from other interfaces.
	IsCommClauseBodyContext()
}

type CommClauseBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommClauseBodyContext() *CommClauseBodyContext {
	var p = new(CommClauseBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClauseBody
	return p
}

func InitEmptyCommClauseBodyContext(p *CommClauseBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_commClauseBody
}

func (*CommClauseBodyContext) IsCommClauseBodyContext() {}

func NewCommClauseBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommClauseBodyContext {
	var p = new(CommClauseBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_commClauseBody

	return p
}

func (s *CommClauseBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *CommClauseBodyContext) SendStmt() ISendStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendStmtContext)
}

func (s *CommClauseBodyContext) RecvStmt() IRecvStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecvStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecvStmtContext)
}

func (s *CommClauseBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommClauseBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommClauseBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCommClauseBody(s)
	}
}

func (s *CommClauseBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCommClauseBody(s)
	}
}

func (p *JLangParser) CommClauseBody() (localctx ICommClauseBodyContext) {
	localctx = NewCommClauseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, JLangParserRULE_commClauseBody)
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.SendStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)
			p.RecvStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecvStmtContext is an interface to support dynamic dispatch.
type IRecvStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprList() IExprListContext
	DECL() antlr.TerminalNode
	Expr() IExprContext

	// IsRecvStmtContext differentiates from other interfaces.
	IsRecvStmtContext()
}

type RecvStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecvStmtContext() *RecvStmtContext {
	var p = new(RecvStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_recvStmt
	return p
}

func InitEmptyRecvStmtContext(p *RecvStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_recvStmt
}

func (*RecvStmtContext) IsRecvStmtContext() {}

func NewRecvStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecvStmtContext {
	var p = new(RecvStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_recvStmt

	return p
}

func (s *RecvStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RecvStmtContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *RecvStmtContext) DECL() antlr.TerminalNode {
	return s.GetToken(JLangParserDECL, 0)
}

func (s *RecvStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RecvStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecvStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecvStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterRecvStmt(s)
	}
}

func (s *RecvStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitRecvStmt(s)
	}
}

func (p *JLangParser) RecvStmt() (localctx IRecvStmtContext) {
	localctx = NewRecvStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, JLangParserRULE_recvStmt)
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(470)
			p.ExprList()
		}
		{
			p.SetState(471)
			p.Match(JLangParserDECL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(474)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE_NAME() antlr.TerminalNode
	ArrayType() IArrayTypeContext
	SliceType() ISliceTypeContext
	MappingType() IMappingTypeContext
	ChannelType() IChannelTypeContext
	StructType() IStructTypeContext
	InterfaceType() IInterfaceTypeContext
	IDENT() antlr.TerminalNode

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(JLangParserTYPE_NAME, 0)
}

func (s *Type_Context) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *Type_Context) SliceType() ISliceTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceTypeContext)
}

func (s *Type_Context) MappingType() IMappingTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMappingTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMappingTypeContext)
}

func (s *Type_Context) ChannelType() IChannelTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChannelTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChannelTypeContext)
}

func (s *Type_Context) StructType() IStructTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *Type_Context) InterfaceType() IInterfaceTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeContext)
}

func (s *Type_Context) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *JLangParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, JLangParserRULE_type_)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.Match(JLangParserTYPE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.ArrayType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.SliceType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(480)
			p.MappingType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(481)
			p.ChannelType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(482)
			p.StructType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(483)
			p.InterfaceType()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(484)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() IType_Context

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACK, 0)
}

func (s *ArrayTypeContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserINT_LIT, 0)
}

func (s *ArrayTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACK, 0)
}

func (s *ArrayTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *JLangParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, JLangParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Match(JLangParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceTypeContext is an interface to support dynamic dispatch.
type ISliceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() IType_Context

	// IsSliceTypeContext differentiates from other interfaces.
	IsSliceTypeContext()
}

type SliceTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceTypeContext() *SliceTypeContext {
	var p = new(SliceTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_sliceType
	return p
}

func InitEmptySliceTypeContext(p *SliceTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_sliceType
}

func (*SliceTypeContext) IsSliceTypeContext() {}

func NewSliceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceTypeContext {
	var p = new(SliceTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_sliceType

	return p
}

func (s *SliceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACK, 0)
}

func (s *SliceTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACK, 0)
}

func (s *SliceTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *SliceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSliceType(s)
	}
}

func (s *SliceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSliceType(s)
	}
}

func (p *JLangParser) SliceType() (localctx ISliceTypeContext) {
	localctx = NewSliceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, JLangParserRULE_sliceType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(494)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMappingTypeContext is an interface to support dynamic dispatch.
type IMappingTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_MAPPING() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllType_() []IType_Context
	Type_(i int) IType_Context
	RBRACK() antlr.TerminalNode

	// IsMappingTypeContext differentiates from other interfaces.
	IsMappingTypeContext()
}

type MappingTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMappingTypeContext() *MappingTypeContext {
	var p = new(MappingTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_mappingType
	return p
}

func InitEmptyMappingTypeContext(p *MappingTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_mappingType
}

func (*MappingTypeContext) IsMappingTypeContext() {}

func NewMappingTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingTypeContext {
	var p = new(MappingTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_mappingType

	return p
}

func (s *MappingTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingTypeContext) KW_MAPPING() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_MAPPING, 0)
}

func (s *MappingTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACK, 0)
}

func (s *MappingTypeContext) AllType_() []IType_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_Context); ok {
			len++
		}
	}

	tst := make([]IType_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_Context); ok {
			tst[i] = t.(IType_Context)
			i++
		}
	}

	return tst
}

func (s *MappingTypeContext) Type_(i int) IType_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MappingTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACK, 0)
}

func (s *MappingTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterMappingType(s)
	}
}

func (s *MappingTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitMappingType(s)
	}
}

func (p *JLangParser) MappingType() (localctx IMappingTypeContext) {
	localctx = NewMappingTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, JLangParserRULE_mappingType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Match(JLangParserKW_MAPPING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(498)
		p.Type_()
	}
	{
		p.SetState(499)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChannelTypeContext is an interface to support dynamic dispatch.
type IChannelTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_CHANNEL() antlr.TerminalNode
	Type_() IType_Context

	// IsChannelTypeContext differentiates from other interfaces.
	IsChannelTypeContext()
}

type ChannelTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannelTypeContext() *ChannelTypeContext {
	var p = new(ChannelTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_channelType
	return p
}

func InitEmptyChannelTypeContext(p *ChannelTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_channelType
}

func (*ChannelTypeContext) IsChannelTypeContext() {}

func NewChannelTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelTypeContext {
	var p = new(ChannelTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_channelType

	return p
}

func (s *ChannelTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ChannelTypeContext) KW_CHANNEL() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_CHANNEL, 0)
}

func (s *ChannelTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ChannelTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChannelTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterChannelType(s)
	}
}

func (s *ChannelTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitChannelType(s)
	}
}

func (p *JLangParser) ChannelType() (localctx IChannelTypeContext) {
	localctx = NewChannelTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, JLangParserRULE_channelType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.Match(JLangParserKW_CHANNEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(503)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructTypeContext is an interface to support dynamic dispatch.
type IStructTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	TypeFieldList() ITypeFieldListContext

	// IsStructTypeContext differentiates from other interfaces.
	IsStructTypeContext()
}

type StructTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeContext() *StructTypeContext {
	var p = new(StructTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_structType
	return p
}

func InitEmptyStructTypeContext(p *StructTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_structType
}

func (*StructTypeContext) IsStructTypeContext() {}

func NewStructTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeContext {
	var p = new(StructTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_structType

	return p
}

func (s *StructTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeContext) KW_STRUCT() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_STRUCT, 0)
}

func (s *StructTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACE, 0)
}

func (s *StructTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACE, 0)
}

func (s *StructTypeContext) TypeFieldList() ITypeFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeFieldListContext)
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterStructType(s)
	}
}

func (s *StructTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitStructType(s)
	}
}

func (p *JLangParser) StructType() (localctx IStructTypeContext) {
	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, JLangParserRULE_structType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Match(JLangParserKW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(507)
			p.TypeFieldList()
		}

	}
	{
		p.SetState(510)
		p.Match(JLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeFieldListContext is an interface to support dynamic dispatch.
type ITypeFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeField() []ITypeFieldContext
	TypeField(i int) ITypeFieldContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsTypeFieldListContext differentiates from other interfaces.
	IsTypeFieldListContext()
}

type TypeFieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFieldListContext() *TypeFieldListContext {
	var p = new(TypeFieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeFieldList
	return p
}

func InitEmptyTypeFieldListContext(p *TypeFieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeFieldList
}

func (*TypeFieldListContext) IsTypeFieldListContext() {}

func NewTypeFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFieldListContext {
	var p = new(TypeFieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_typeFieldList

	return p
}

func (s *TypeFieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFieldListContext) AllTypeField() []ITypeFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeFieldContext); ok {
			len++
		}
	}

	tst := make([]ITypeFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeFieldContext); ok {
			tst[i] = t.(ITypeFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeFieldListContext) TypeField(i int) ITypeFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeFieldContext)
}

func (s *TypeFieldListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *TypeFieldListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *TypeFieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterTypeFieldList(s)
	}
}

func (s *TypeFieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitTypeFieldList(s)
	}
}

func (p *JLangParser) TypeFieldList() (localctx ITypeFieldListContext) {
	localctx = NewTypeFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, JLangParserRULE_typeFieldList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.TypeField()
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(513)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(514)
				p.TypeField()
			}

		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(520)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeFieldContext is an interface to support dynamic dispatch.
type ITypeFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() IType_Context

	// IsTypeFieldContext differentiates from other interfaces.
	IsTypeFieldContext()
}

type TypeFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFieldContext() *TypeFieldContext {
	var p = new(TypeFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeField
	return p
}

func InitEmptyTypeFieldContext(p *TypeFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_typeField
}

func (*TypeFieldContext) IsTypeFieldContext() {}

func NewTypeFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFieldContext {
	var p = new(TypeFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_typeField

	return p
}

func (s *TypeFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *TypeFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(JLangParserCOLON, 0)
}

func (s *TypeFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterTypeField(s)
	}
}

func (s *TypeFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitTypeField(s)
	}
}

func (p *JLangParser) TypeField() (localctx ITypeFieldContext) {
	localctx = NewTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, JLangParserRULE_typeField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceTypeContext is an interface to support dynamic dispatch.
type IInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_INTERFACE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	MethodList() IMethodListContext

	// IsInterfaceTypeContext differentiates from other interfaces.
	IsInterfaceTypeContext()
}

type InterfaceTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeContext() *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_interfaceType
	return p
}

func InitEmptyInterfaceTypeContext(p *InterfaceTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_interfaceType
}

func (*InterfaceTypeContext) IsInterfaceTypeContext() {}

func NewInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_interfaceType

	return p
}

func (s *InterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeContext) KW_INTERFACE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_INTERFACE, 0)
}

func (s *InterfaceTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACE, 0)
}

func (s *InterfaceTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACE, 0)
}

func (s *InterfaceTypeContext) MethodList() IMethodListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodListContext)
}

func (s *InterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterInterfaceType(s)
	}
}

func (s *InterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitInterfaceType(s)
	}
}

func (p *JLangParser) InterfaceType() (localctx IInterfaceTypeContext) {
	localctx = NewInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, JLangParserRULE_interfaceType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(527)
		p.Match(JLangParserKW_INTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(529)
			p.MethodList()
		}

	}
	{
		p.SetState(532)
		p.Match(JLangParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodListContext is an interface to support dynamic dispatch.
type IMethodListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMethodDecl() []IMethodDeclContext
	MethodDecl(i int) IMethodDeclContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsMethodListContext differentiates from other interfaces.
	IsMethodListContext()
}

type MethodListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodListContext() *MethodListContext {
	var p = new(MethodListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_methodList
	return p
}

func InitEmptyMethodListContext(p *MethodListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_methodList
}

func (*MethodListContext) IsMethodListContext() {}

func NewMethodListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodListContext {
	var p = new(MethodListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_methodList

	return p
}

func (s *MethodListContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodListContext) AllMethodDecl() []IMethodDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodDeclContext); ok {
			len++
		}
	}

	tst := make([]IMethodDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodDeclContext); ok {
			tst[i] = t.(IMethodDeclContext)
			i++
		}
	}

	return tst
}

func (s *MethodListContext) MethodDecl(i int) IMethodDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodDeclContext)
}

func (s *MethodListContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSEMI)
}

func (s *MethodListContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSEMI, i)
}

func (s *MethodListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterMethodList(s)
	}
}

func (s *MethodListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitMethodList(s)
	}
}

func (p *JLangParser) MethodList() (localctx IMethodListContext) {
	localctx = NewMethodListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, JLangParserRULE_methodList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.MethodDecl()
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(535)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(536)
				p.MethodDecl()
			}

		}
		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(542)
			p.Match(JLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethodDeclContext is an interface to support dynamic dispatch.
type IMethodDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	Signature() ISignatureContext

	// IsMethodDeclContext differentiates from other interfaces.
	IsMethodDeclContext()
}

type MethodDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclContext() *MethodDeclContext {
	var p = new(MethodDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_methodDecl
	return p
}

func InitEmptyMethodDeclContext(p *MethodDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_methodDecl
}

func (*MethodDeclContext) IsMethodDeclContext() {}

func NewMethodDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclContext {
	var p = new(MethodDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_methodDecl

	return p
}

func (s *MethodDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *MethodDeclContext) Signature() ISignatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignatureContext)
}

func (s *MethodDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterMethodDecl(s)
	}
}

func (s *MethodDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitMethodDecl(s)
	}
}

func (p *JLangParser) MethodDecl() (localctx IMethodDeclContext) {
	localctx = NewMethodDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, JLangParserRULE_methodDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(545)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(546)
		p.Signature()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	OROR() antlr.TerminalNode
	ANDAND() antlr.TerminalNode
	BOR() antlr.TerminalNode
	BXOR() antlr.TerminalNode
	BAND() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	SHL() antlr.TerminalNode
	SHR() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PERCENT() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) OROR() antlr.TerminalNode {
	return s.GetToken(JLangParserOROR, 0)
}

func (s *ExprContext) ANDAND() antlr.TerminalNode {
	return s.GetToken(JLangParserANDAND, 0)
}

func (s *ExprContext) BOR() antlr.TerminalNode {
	return s.GetToken(JLangParserBOR, 0)
}

func (s *ExprContext) BXOR() antlr.TerminalNode {
	return s.GetToken(JLangParserBXOR, 0)
}

func (s *ExprContext) BAND() antlr.TerminalNode {
	return s.GetToken(JLangParserBAND, 0)
}

func (s *ExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(JLangParserEQ, 0)
}

func (s *ExprContext) NE() antlr.TerminalNode {
	return s.GetToken(JLangParserNE, 0)
}

func (s *ExprContext) LT() antlr.TerminalNode {
	return s.GetToken(JLangParserLT, 0)
}

func (s *ExprContext) LE() antlr.TerminalNode {
	return s.GetToken(JLangParserLE, 0)
}

func (s *ExprContext) GT() antlr.TerminalNode {
	return s.GetToken(JLangParserGT, 0)
}

func (s *ExprContext) GE() antlr.TerminalNode {
	return s.GetToken(JLangParserGE, 0)
}

func (s *ExprContext) SHL() antlr.TerminalNode {
	return s.GetToken(JLangParserSHL, 0)
}

func (s *ExprContext) SHR() antlr.TerminalNode {
	return s.GetToken(JLangParserSHR, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JLangParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JLangParserMINUS, 0)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(JLangParserSTAR, 0)
}

func (s *ExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(JLangParserSLASH, 0)
}

func (s *ExprContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(JLangParserPERCENT, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *JLangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *JLangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 124
	p.EnterRecursionRule(localctx, 124, JLangParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.UnaryExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(581)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(551)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(552)
					p.Match(JLangParserOROR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(553)
					p.expr(12)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(554)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(555)
					p.Match(JLangParserANDAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(556)
					p.expr(11)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(557)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(558)
					p.Match(JLangParserBOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(559)
					p.expr(10)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(560)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(561)
					p.Match(JLangParserBXOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(562)
					p.expr(9)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(563)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(564)
					p.Match(JLangParserBAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(565)
					p.expr(8)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(566)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(567)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JLangParserEQ || _la == JLangParserNE) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(568)
					p.expr(7)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(569)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(570)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&25769803779) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(571)
					p.expr(6)
				}

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(572)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(573)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JLangParserSHL || _la == JLangParserSHR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(574)
					p.expr(5)
				}

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(575)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(576)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JLangParserPLUS || _la == JLangParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(577)
					p.expr(4)
				}

			case 10:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_expr)
				p.SetState(578)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(579)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&7) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(580)
					p.expr(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(585)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	PLUS() antlr.TerminalNode
	UnaryExpr() IUnaryExprContext
	MINUS() antlr.TerminalNode
	BANG() antlr.TerminalNode
	TILDE() antlr.TerminalNode
	STAR() antlr.TerminalNode
	BAND() antlr.TerminalNode

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JLangParserPLUS, 0)
}

func (s *UnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JLangParserMINUS, 0)
}

func (s *UnaryExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(JLangParserBANG, 0)
}

func (s *UnaryExprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(JLangParserTILDE, 0)
}

func (s *UnaryExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(JLangParserSTAR, 0)
}

func (s *UnaryExprContext) BAND() antlr.TerminalNode {
	return s.GetToken(JLangParserBAND, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (p *JLangParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, JLangParserRULE_unaryExpr)
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_RECOVER, JLangParserIDENT, JLangParserINT_LIT, JLangParserFLOAT_LIT, JLangParserSTRING_LIT, JLangParserRAW_STR, JLangParserCHAR_LIT, JLangParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(586)
			p.primaryExpr(0)
		}

	case JLangParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(587)
			p.Match(JLangParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.UnaryExpr()
		}

	case JLangParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(589)
			p.Match(JLangParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.UnaryExpr()
		}

	case JLangParserBANG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(591)
			p.Match(JLangParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(592)
			p.UnaryExpr()
		}

	case JLangParserTILDE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(593)
			p.Match(JLangParserTILDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)
			p.UnaryExpr()
		}

	case JLangParserSTAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(595)
			p.Match(JLangParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(596)
			p.UnaryExpr()
		}

	case JLangParserBAND:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(597)
			p.Match(JLangParserBAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)
			p.UnaryExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	PrimaryExpr() IPrimaryExprContext
	Selector() ISelectorContext
	Index() IIndexContext
	Call() ICallContext

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *JLangParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *JLangParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 128
	p.EnterRecursionRule(localctx, 128, JLangParserRULE_primaryExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(602)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(610)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(604)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(605)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(606)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(607)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(608)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(609)
					p.Call()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode
	FLOAT_LIT() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	CHAR_LIT() antlr.TerminalNode
	RAW_STR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	KW_RECOVER() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *OperandContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserINT_LIT, 0)
}

func (s *OperandContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserFLOAT_LIT, 0)
}

func (s *OperandContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserSTRING_LIT, 0)
}

func (s *OperandContext) CHAR_LIT() antlr.TerminalNode {
	return s.GetToken(JLangParserCHAR_LIT, 0)
}

func (s *OperandContext) RAW_STR() antlr.TerminalNode {
	return s.GetToken(JLangParserRAW_STR, 0)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserLPAREN, 0)
}

func (s *OperandContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserRPAREN, 0)
}

func (s *OperandContext) KW_RECOVER() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_RECOVER, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *JLangParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, JLangParserRULE_operand)
	p.SetState(628)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(615)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserINT_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(616)
			p.Match(JLangParserINT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(617)
			p.Match(JLangParserFLOAT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(618)
			p.Match(JLangParserSTRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserCHAR_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(619)
			p.Match(JLangParserCHAR_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserRAW_STR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(620)
			p.Match(JLangParserRAW_STR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserLPAREN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(621)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.expr(0)
		}
		{
			p.SetState(623)
			p.Match(JLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserKW_RECOVER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(625)
			p.Match(JLangParserKW_RECOVER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(627)
			p.Match(JLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(JLangParserDOT, 0)
}

func (s *SelectorContext) IDENT() antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *JLangParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, JLangParserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(630)
		p.Match(JLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(631)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Expr() IExprContext
	RBRACK() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserLBRACK, 0)
}

func (s *IndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(JLangParserRBRACK, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *JLangParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, JLangParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(633)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(634)
		p.expr(0)
	}
	{
		p.SetState(635)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgList() IArgListContext

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserLPAREN, 0)
}

func (s *CallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserRPAREN, 0)
}

func (s *CallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *JLangParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, JLangParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(637)
		p.Match(JLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&1707404152537341) != 0 {
		{
			p.SetState(638)
			p.ArgList()
		}

	}
	{
		p.SetState(641)
		p.Match(JLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JLangParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *JLangParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, JLangParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(643)
		p.expr(0)
	}
	p.SetState(648)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(644)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(645)
			p.expr(0)
		}

		p.SetState(650)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentListContext is an interface to support dynamic dispatch.
type IIdentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentListContext differentiates from other interfaces.
	IsIdentListContext()
}

type IdentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentListContext() *IdentListContext {
	var p = new(IdentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_identList
	return p
}

func InitEmptyIdentListContext(p *IdentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_identList
}

func (*IdentListContext) IsIdentListContext() {}

func NewIdentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentListContext {
	var p = new(IdentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_identList

	return p
}

func (s *IdentListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentListContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(JLangParserIDENT)
}

func (s *IdentListContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserIDENT, i)
}

func (s *IdentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JLangParserCOMMA)
}

func (s *IdentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, i)
}

func (s *IdentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterIdentList(s)
	}
}

func (s *IdentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitIdentList(s)
	}
}

func (p *JLangParser) IdentList() (localctx IIdentListContext) {
	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, JLangParserRULE_identList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(652)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(653)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(658)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JLangParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *JLangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, JLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(659)
		p.expr(0)
	}
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(660)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(661)
			p.expr(0)
		}

		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *JLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 62:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 64:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JLangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JLangParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
