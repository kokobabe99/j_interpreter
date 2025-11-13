// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser_src // JLang
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
		"'make'", "", "", "", "", "", "", "", "'<-'", "'<='", "'>='", "'=='",
		"'!='", "'<<='", "'>>='", "'<<'", "'>>'", "'+='", "'-='", "'*='", "'/='",
		"'%='", "'&='", "'|='", "'^='", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"','", "';'", "':'", "'.'", "'='", "':='", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'<'", "'>'", "'&'", "'|'", "'^'", "'&&'", "'||'", "'!'", "'~'",
	}
	staticData.SymbolicNames = []string{
		"", "KW_PKG", "KW_IMP", "KW_DEF", "KW_VAR", "KW_CONS", "KW_TYPE", "KW_STRUCT",
		"KW_INTERFACE", "KW_MAPPING", "KW_CHANNEL", "KW_J", "KW_SELECT", "KW_LATER",
		"KW_RET", "KW_IF", "KW_ELSE", "KW_SWITCH", "KW_CASE", "KW_FALL", "KW_FR",
		"KW_RANGE", "KW_BREAK", "KW_CONTINUE", "KW_JOTO", "KW_DFT", "KW_PANIC",
		"KW_RECOVER", "KW_MAKE", "TYPE_NAME", "IDENT", "INT_LIT", "FLOAT_LIT",
		"STRING_LIT", "RAW_STR", "CHAR_LIT", "CH_SEND", "LE", "GE", "EQ", "NE",
		"SHLEQ", "SHREQ", "SHL", "SHR", "ADDEQ", "SUBEQ", "MULEQ", "DIVEQ",
		"MODEQ", "ANDEQ", "OREQ", "XOREQ", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"LBRACK", "RBRACK", "COMMA", "SEMI", "COLON", "DOT", "ASSIGN", "DECL",
		"PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "LT", "GT", "BAND", "BOR",
		"BXOR", "ANDAND", "OROR", "BANG", "TILDE", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "pkgDecl", "impList", "impDecl", "importSpec", "topDecl",
		"varDecl", "varSpec", "consDecl", "constSpec", "typeDecl", "typeSpec",
		"funcDecl", "signature", "paramList", "param", "resultType", "block",
		"stmtList", "stmt", "declStmt", "labeledStmt", "simpleStmt", "exprStmt",
		"assignStmt", "lhs", "assignOp", "shortVarDecl", "sendStmt", "spawnStmt",
		"returnStmt", "deferStmt", "breakStmt", "continueStmt", "jotoStmt",
		"panicStmt", "ifStmt", "elseOpt", "switchStmt", "switchHead", "caseClauses",
		"caseClause", "fallOpt", "forStmt", "forClause", "rangeStmt", "rangeHeader",
		"selectStmt", "commClauses", "commClause", "commClauseBody", "recvStmt",
		"type_", "arrayType", "sliceType", "mappingType", "channelType", "structType",
		"typeFieldList", "typeField", "interfaceType", "methodList", "methodDecl",
		"expr", "logicalOrExpr", "logicalAndExpr", "bitOrExpr", "bitXorExpr",
		"bitAndExpr", "equalityExpr", "relationalExpr", "shiftExpr", "additiveExpr",
		"multiplicativeExpr", "unaryExpr", "primaryExpr", "operand", "makeExpr",
		"selector", "index", "call", "argList", "identList", "exprList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 81, 749, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78, 7, 78,
		2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7, 83, 1,
		0, 1, 0, 3, 0, 171, 8, 0, 1, 0, 5, 0, 174, 8, 0, 10, 0, 12, 0, 177, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 185, 8, 2, 11, 2, 12, 2, 186,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 196, 8, 3, 10, 3, 12, 3,
		199, 9, 3, 1, 3, 1, 3, 3, 3, 203, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 208, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 215, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 221, 8, 6, 10, 6, 12, 6, 224, 9, 6, 1, 6, 3, 6, 227, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 233, 8, 7, 1, 7, 1, 7, 3, 7, 237, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 243, 8, 8, 10, 8, 12, 8, 246, 9, 8, 1, 8, 3, 8,
		249, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 258, 8, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 5, 10, 264, 8, 10, 10, 10, 12, 10, 267, 9, 10,
		1, 10, 3, 10, 270, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 276, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 285, 8, 13, 1,
		13, 1, 13, 3, 13, 289, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 294, 8, 14, 10,
		14, 12, 14, 297, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 3, 17, 308, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 5,
		18, 315, 8, 18, 10, 18, 12, 18, 318, 9, 18, 1, 18, 3, 18, 321, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 338, 8, 19, 1, 20, 1, 20, 1, 20, 3,
		20, 343, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 3, 22, 354, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 379, 8, 30, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 3, 32, 386, 8, 32, 1, 33, 1, 33, 3, 33, 390, 8,
		33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 408, 8, 37, 1, 38, 1,
		38, 3, 38, 412, 8, 38, 1, 38, 1, 38, 3, 38, 416, 8, 38, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 40, 4, 40, 423, 8, 40, 11, 40, 12, 40, 424, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 436, 8, 41, 1,
		42, 1, 42, 3, 42, 440, 8, 42, 1, 43, 1, 43, 3, 43, 444, 8, 43, 1, 43, 1,
		43, 1, 44, 3, 44, 449, 8, 44, 1, 44, 1, 44, 3, 44, 453, 8, 44, 1, 44, 1,
		44, 3, 44, 457, 8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46,
		1, 46, 3, 46, 467, 8, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 3, 47, 474,
		8, 47, 1, 47, 1, 47, 1, 48, 4, 48, 479, 8, 48, 11, 48, 12, 48, 480, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 491, 8, 49,
		1, 50, 1, 50, 3, 50, 495, 8, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3,
		51, 502, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52,
		3, 52, 512, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56,
		1, 57, 1, 57, 1, 57, 3, 57, 535, 8, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1,
		58, 5, 58, 542, 8, 58, 10, 58, 12, 58, 545, 9, 58, 1, 58, 3, 58, 548, 8,
		58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 3, 60, 557, 8, 60,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 5, 61, 564, 8, 61, 10, 61, 12, 61, 567,
		9, 61, 1, 61, 3, 61, 570, 8, 61, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1,
		64, 1, 64, 1, 64, 5, 64, 580, 8, 64, 10, 64, 12, 64, 583, 9, 64, 1, 65,
		1, 65, 1, 65, 5, 65, 588, 8, 65, 10, 65, 12, 65, 591, 9, 65, 1, 66, 1,
		66, 1, 66, 5, 66, 596, 8, 66, 10, 66, 12, 66, 599, 9, 66, 1, 67, 1, 67,
		1, 67, 5, 67, 604, 8, 67, 10, 67, 12, 67, 607, 9, 67, 1, 68, 1, 68, 1,
		68, 5, 68, 612, 8, 68, 10, 68, 12, 68, 615, 9, 68, 1, 69, 1, 69, 1, 69,
		5, 69, 620, 8, 69, 10, 69, 12, 69, 623, 9, 69, 1, 70, 1, 70, 1, 70, 5,
		70, 628, 8, 70, 10, 70, 12, 70, 631, 9, 70, 1, 71, 1, 71, 1, 71, 5, 71,
		636, 8, 71, 10, 71, 12, 71, 639, 9, 71, 1, 72, 1, 72, 1, 72, 5, 72, 644,
		8, 72, 10, 72, 12, 72, 647, 9, 72, 1, 73, 1, 73, 1, 73, 5, 73, 652, 8,
		73, 10, 73, 12, 73, 655, 9, 73, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74,
		1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 3, 74, 670, 8, 74, 1,
		75, 1, 75, 1, 75, 3, 75, 675, 8, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75,
		1, 75, 5, 75, 683, 8, 75, 10, 75, 12, 75, 686, 9, 75, 1, 76, 1, 76, 1,
		76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76,
		3, 76, 701, 8, 76, 1, 77, 1, 77, 1, 77, 1, 77, 1, 77, 3, 77, 708, 8, 77,
		1, 77, 1, 77, 1, 78, 1, 78, 1, 78, 1, 79, 1, 79, 1, 79, 1, 79, 1, 80, 1,
		80, 3, 80, 721, 8, 80, 1, 80, 1, 80, 1, 81, 1, 81, 1, 81, 5, 81, 728, 8,
		81, 10, 81, 12, 81, 731, 9, 81, 1, 82, 1, 82, 1, 82, 5, 82, 736, 8, 82,
		10, 82, 12, 82, 739, 9, 82, 1, 83, 1, 83, 1, 83, 5, 83, 744, 8, 83, 10,
		83, 12, 83, 747, 9, 83, 1, 83, 0, 1, 150, 84, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
		120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148,
		150, 152, 154, 156, 158, 160, 162, 164, 166, 0, 6, 3, 0, 41, 42, 45, 52,
		63, 63, 1, 0, 39, 40, 2, 0, 37, 38, 70, 71, 1, 0, 43, 44, 1, 0, 65, 66,
		1, 0, 67, 69, 776, 0, 168, 1, 0, 0, 0, 2, 180, 1, 0, 0, 0, 4, 184, 1, 0,
		0, 0, 6, 202, 1, 0, 0, 0, 8, 207, 1, 0, 0, 0, 10, 214, 1, 0, 0, 0, 12,
		216, 1, 0, 0, 0, 14, 228, 1, 0, 0, 0, 16, 238, 1, 0, 0, 0, 18, 250, 1,
		0, 0, 0, 20, 259, 1, 0, 0, 0, 22, 271, 1, 0, 0, 0, 24, 277, 1, 0, 0, 0,
		26, 282, 1, 0, 0, 0, 28, 290, 1, 0, 0, 0, 30, 298, 1, 0, 0, 0, 32, 302,
		1, 0, 0, 0, 34, 305, 1, 0, 0, 0, 36, 311, 1, 0, 0, 0, 38, 337, 1, 0, 0,
		0, 40, 342, 1, 0, 0, 0, 42, 344, 1, 0, 0, 0, 44, 353, 1, 0, 0, 0, 46, 355,
		1, 0, 0, 0, 48, 357, 1, 0, 0, 0, 50, 361, 1, 0, 0, 0, 52, 363, 1, 0, 0,
		0, 54, 365, 1, 0, 0, 0, 56, 369, 1, 0, 0, 0, 58, 373, 1, 0, 0, 0, 60, 376,
		1, 0, 0, 0, 62, 380, 1, 0, 0, 0, 64, 383, 1, 0, 0, 0, 66, 387, 1, 0, 0,
		0, 68, 391, 1, 0, 0, 0, 70, 394, 1, 0, 0, 0, 72, 397, 1, 0, 0, 0, 74, 407,
		1, 0, 0, 0, 76, 409, 1, 0, 0, 0, 78, 419, 1, 0, 0, 0, 80, 422, 1, 0, 0,
		0, 82, 435, 1, 0, 0, 0, 84, 439, 1, 0, 0, 0, 86, 441, 1, 0, 0, 0, 88, 448,
		1, 0, 0, 0, 90, 458, 1, 0, 0, 0, 92, 466, 1, 0, 0, 0, 94, 470, 1, 0, 0,
		0, 96, 478, 1, 0, 0, 0, 98, 490, 1, 0, 0, 0, 100, 494, 1, 0, 0, 0, 102,
		501, 1, 0, 0, 0, 104, 511, 1, 0, 0, 0, 106, 513, 1, 0, 0, 0, 108, 518,
		1, 0, 0, 0, 110, 522, 1, 0, 0, 0, 112, 528, 1, 0, 0, 0, 114, 531, 1, 0,
		0, 0, 116, 538, 1, 0, 0, 0, 118, 549, 1, 0, 0, 0, 120, 553, 1, 0, 0, 0,
		122, 560, 1, 0, 0, 0, 124, 571, 1, 0, 0, 0, 126, 574, 1, 0, 0, 0, 128,
		576, 1, 0, 0, 0, 130, 584, 1, 0, 0, 0, 132, 592, 1, 0, 0, 0, 134, 600,
		1, 0, 0, 0, 136, 608, 1, 0, 0, 0, 138, 616, 1, 0, 0, 0, 140, 624, 1, 0,
		0, 0, 142, 632, 1, 0, 0, 0, 144, 640, 1, 0, 0, 0, 146, 648, 1, 0, 0, 0,
		148, 669, 1, 0, 0, 0, 150, 674, 1, 0, 0, 0, 152, 700, 1, 0, 0, 0, 154,
		702, 1, 0, 0, 0, 156, 711, 1, 0, 0, 0, 158, 714, 1, 0, 0, 0, 160, 718,
		1, 0, 0, 0, 162, 724, 1, 0, 0, 0, 164, 732, 1, 0, 0, 0, 166, 740, 1, 0,
		0, 0, 168, 170, 3, 2, 1, 0, 169, 171, 3, 4, 2, 0, 170, 169, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 175, 1, 0, 0, 0, 172, 174, 3, 10, 5, 0, 173,
		172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176,
		1, 0, 0, 0, 176, 178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 5, 0,
		0, 1, 179, 1, 1, 0, 0, 0, 180, 181, 5, 1, 0, 0, 181, 182, 5, 30, 0, 0,
		182, 3, 1, 0, 0, 0, 183, 185, 3, 6, 3, 0, 184, 183, 1, 0, 0, 0, 185, 186,
		1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 5, 1, 0, 0,
		0, 188, 189, 5, 2, 0, 0, 189, 203, 3, 8, 4, 0, 190, 191, 5, 2, 0, 0, 191,
		192, 5, 53, 0, 0, 192, 197, 3, 8, 4, 0, 193, 194, 5, 59, 0, 0, 194, 196,
		3, 8, 4, 0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		200, 201, 5, 54, 0, 0, 201, 203, 1, 0, 0, 0, 202, 188, 1, 0, 0, 0, 202,
		190, 1, 0, 0, 0, 203, 7, 1, 0, 0, 0, 204, 208, 5, 33, 0, 0, 205, 206, 5,
		30, 0, 0, 206, 208, 5, 33, 0, 0, 207, 204, 1, 0, 0, 0, 207, 205, 1, 0,
		0, 0, 208, 9, 1, 0, 0, 0, 209, 215, 3, 12, 6, 0, 210, 215, 3, 16, 8, 0,
		211, 215, 3, 20, 10, 0, 212, 215, 3, 24, 12, 0, 213, 215, 3, 38, 19, 0,
		214, 209, 1, 0, 0, 0, 214, 210, 1, 0, 0, 0, 214, 211, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 11, 1, 0, 0, 0, 216, 217, 5,
		4, 0, 0, 217, 222, 3, 14, 7, 0, 218, 219, 5, 60, 0, 0, 219, 221, 3, 14,
		7, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225,
		227, 5, 60, 0, 0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 13,
		1, 0, 0, 0, 228, 236, 3, 164, 82, 0, 229, 232, 3, 104, 52, 0, 230, 231,
		5, 63, 0, 0, 231, 233, 3, 166, 83, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1,
		0, 0, 0, 233, 237, 1, 0, 0, 0, 234, 235, 5, 63, 0, 0, 235, 237, 3, 166,
		83, 0, 236, 229, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0,
		237, 15, 1, 0, 0, 0, 238, 239, 5, 5, 0, 0, 239, 244, 3, 18, 9, 0, 240,
		241, 5, 60, 0, 0, 241, 243, 3, 18, 9, 0, 242, 240, 1, 0, 0, 0, 243, 246,
		1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 248, 1, 0,
		0, 0, 246, 244, 1, 0, 0, 0, 247, 249, 5, 60, 0, 0, 248, 247, 1, 0, 0, 0,
		248, 249, 1, 0, 0, 0, 249, 17, 1, 0, 0, 0, 250, 257, 3, 164, 82, 0, 251,
		252, 3, 104, 52, 0, 252, 253, 5, 63, 0, 0, 253, 254, 3, 166, 83, 0, 254,
		258, 1, 0, 0, 0, 255, 256, 5, 63, 0, 0, 256, 258, 3, 166, 83, 0, 257, 251,
		1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 19, 1, 0, 0, 0, 259, 260, 5, 6,
		0, 0, 260, 265, 3, 22, 11, 0, 261, 262, 5, 60, 0, 0, 262, 264, 3, 22, 11,
		0, 263, 261, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265,
		266, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 268, 270,
		5, 60, 0, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 21, 1, 0,
		0, 0, 271, 275, 5, 30, 0, 0, 272, 273, 5, 63, 0, 0, 273, 276, 3, 104, 52,
		0, 274, 276, 3, 104, 52, 0, 275, 272, 1, 0, 0, 0, 275, 274, 1, 0, 0, 0,
		276, 23, 1, 0, 0, 0, 277, 278, 5, 3, 0, 0, 278, 279, 5, 30, 0, 0, 279,
		280, 3, 26, 13, 0, 280, 281, 3, 34, 17, 0, 281, 25, 1, 0, 0, 0, 282, 284,
		5, 53, 0, 0, 283, 285, 3, 28, 14, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1,
		0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 288, 5, 54, 0, 0, 287, 289, 3, 32,
		16, 0, 288, 287, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 27, 1, 0, 0, 0,
		290, 295, 3, 30, 15, 0, 291, 292, 5, 59, 0, 0, 292, 294, 3, 30, 15, 0,
		293, 291, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295,
		296, 1, 0, 0, 0, 296, 29, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 298, 299, 5,
		30, 0, 0, 299, 300, 5, 61, 0, 0, 300, 301, 3, 104, 52, 0, 301, 31, 1, 0,
		0, 0, 302, 303, 5, 61, 0, 0, 303, 304, 3, 104, 52, 0, 304, 33, 1, 0, 0,
		0, 305, 307, 5, 55, 0, 0, 306, 308, 3, 36, 18, 0, 307, 306, 1, 0, 0, 0,
		307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 5, 56, 0, 0, 310,
		35, 1, 0, 0, 0, 311, 316, 3, 38, 19, 0, 312, 313, 5, 60, 0, 0, 313, 315,
		3, 38, 19, 0, 314, 312, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1,
		0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0,
		0, 319, 321, 5, 60, 0, 0, 320, 319, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		37, 1, 0, 0, 0, 322, 338, 3, 44, 22, 0, 323, 338, 3, 40, 20, 0, 324, 338,
		3, 72, 36, 0, 325, 338, 3, 76, 38, 0, 326, 338, 3, 86, 43, 0, 327, 338,
		3, 90, 45, 0, 328, 338, 3, 94, 47, 0, 329, 338, 3, 62, 31, 0, 330, 338,
		3, 60, 30, 0, 331, 338, 3, 64, 32, 0, 332, 338, 3, 66, 33, 0, 333, 338,
		3, 68, 34, 0, 334, 338, 3, 70, 35, 0, 335, 338, 3, 42, 21, 0, 336, 338,
		3, 34, 17, 0, 337, 322, 1, 0, 0, 0, 337, 323, 1, 0, 0, 0, 337, 324, 1,
		0, 0, 0, 337, 325, 1, 0, 0, 0, 337, 326, 1, 0, 0, 0, 337, 327, 1, 0, 0,
		0, 337, 328, 1, 0, 0, 0, 337, 329, 1, 0, 0, 0, 337, 330, 1, 0, 0, 0, 337,
		331, 1, 0, 0, 0, 337, 332, 1, 0, 0, 0, 337, 333, 1, 0, 0, 0, 337, 334,
		1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 336, 1, 0, 0, 0, 338, 39, 1, 0,
		0, 0, 339, 343, 3, 12, 6, 0, 340, 343, 3, 16, 8, 0, 341, 343, 3, 20, 10,
		0, 342, 339, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 341, 1, 0, 0, 0, 343,
		41, 1, 0, 0, 0, 344, 345, 5, 30, 0, 0, 345, 346, 5, 61, 0, 0, 346, 347,
		3, 38, 19, 0, 347, 43, 1, 0, 0, 0, 348, 354, 3, 46, 23, 0, 349, 354, 3,
		48, 24, 0, 350, 354, 3, 54, 27, 0, 351, 354, 3, 56, 28, 0, 352, 354, 3,
		58, 29, 0, 353, 348, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 353, 350, 1, 0,
		0, 0, 353, 351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0, 354, 45, 1, 0, 0, 0,
		355, 356, 3, 126, 63, 0, 356, 47, 1, 0, 0, 0, 357, 358, 3, 50, 25, 0, 358,
		359, 3, 52, 26, 0, 359, 360, 3, 166, 83, 0, 360, 49, 1, 0, 0, 0, 361, 362,
		3, 150, 75, 0, 362, 51, 1, 0, 0, 0, 363, 364, 7, 0, 0, 0, 364, 53, 1, 0,
		0, 0, 365, 366, 3, 164, 82, 0, 366, 367, 5, 64, 0, 0, 367, 368, 3, 166,
		83, 0, 368, 55, 1, 0, 0, 0, 369, 370, 3, 126, 63, 0, 370, 371, 5, 36, 0,
		0, 371, 372, 3, 126, 63, 0, 372, 57, 1, 0, 0, 0, 373, 374, 5, 11, 0, 0,
		374, 375, 3, 126, 63, 0, 375, 59, 1, 0, 0, 0, 376, 378, 5, 14, 0, 0, 377,
		379, 3, 166, 83, 0, 378, 377, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 61,
		1, 0, 0, 0, 380, 381, 5, 13, 0, 0, 381, 382, 3, 126, 63, 0, 382, 63, 1,
		0, 0, 0, 383, 385, 5, 22, 0, 0, 384, 386, 5, 30, 0, 0, 385, 384, 1, 0,
		0, 0, 385, 386, 1, 0, 0, 0, 386, 65, 1, 0, 0, 0, 387, 389, 5, 23, 0, 0,
		388, 390, 5, 30, 0, 0, 389, 388, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390,
		67, 1, 0, 0, 0, 391, 392, 5, 24, 0, 0, 392, 393, 5, 30, 0, 0, 393, 69,
		1, 0, 0, 0, 394, 395, 5, 26, 0, 0, 395, 396, 3, 126, 63, 0, 396, 71, 1,
		0, 0, 0, 397, 398, 5, 15, 0, 0, 398, 399, 3, 126, 63, 0, 399, 400, 3, 34,
		17, 0, 400, 401, 3, 74, 37, 0, 401, 73, 1, 0, 0, 0, 402, 403, 5, 16, 0,
		0, 403, 408, 3, 72, 36, 0, 404, 405, 5, 16, 0, 0, 405, 408, 3, 34, 17,
		0, 406, 408, 1, 0, 0, 0, 407, 402, 1, 0, 0, 0, 407, 404, 1, 0, 0, 0, 407,
		406, 1, 0, 0, 0, 408, 75, 1, 0, 0, 0, 409, 411, 5, 17, 0, 0, 410, 412,
		3, 78, 39, 0, 411, 410, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 1,
		0, 0, 0, 413, 415, 5, 55, 0, 0, 414, 416, 3, 80, 40, 0, 415, 414, 1, 0,
		0, 0, 415, 416, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 418, 5, 56, 0, 0,
		418, 77, 1, 0, 0, 0, 419, 420, 3, 126, 63, 0, 420, 79, 1, 0, 0, 0, 421,
		423, 3, 82, 41, 0, 422, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 422,
		1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 81, 1, 0, 0, 0, 426, 427, 5, 18,
		0, 0, 427, 428, 3, 166, 83, 0, 428, 429, 5, 61, 0, 0, 429, 430, 3, 36,
		18, 0, 430, 431, 3, 84, 42, 0, 431, 436, 1, 0, 0, 0, 432, 433, 5, 25, 0,
		0, 433, 434, 5, 61, 0, 0, 434, 436, 3, 36, 18, 0, 435, 426, 1, 0, 0, 0,
		435, 432, 1, 0, 0, 0, 436, 83, 1, 0, 0, 0, 437, 440, 5, 19, 0, 0, 438,
		440, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 438, 1, 0, 0, 0, 440, 85, 1,
		0, 0, 0, 441, 443, 5, 20, 0, 0, 442, 444, 3, 88, 44, 0, 443, 442, 1, 0,
		0, 0, 443, 444, 1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 3, 34, 17,
		0, 446, 87, 1, 0, 0, 0, 447, 449, 3, 44, 22, 0, 448, 447, 1, 0, 0, 0, 448,
		449, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 452, 5, 60, 0, 0, 451, 453,
		3, 126, 63, 0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 1,
		0, 0, 0, 454, 456, 5, 60, 0, 0, 455, 457, 3, 44, 22, 0, 456, 455, 1, 0,
		0, 0, 456, 457, 1, 0, 0, 0, 457, 89, 1, 0, 0, 0, 458, 459, 5, 20, 0, 0,
		459, 460, 5, 21, 0, 0, 460, 461, 3, 92, 46, 0, 461, 462, 3, 34, 17, 0,
		462, 91, 1, 0, 0, 0, 463, 464, 3, 164, 82, 0, 464, 465, 5, 64, 0, 0, 465,
		467, 1, 0, 0, 0, 466, 463, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 468,
		1, 0, 0, 0, 468, 469, 3, 126, 63, 0, 469, 93, 1, 0, 0, 0, 470, 471, 5,
		12, 0, 0, 471, 473, 5, 55, 0, 0, 472, 474, 3, 96, 48, 0, 473, 472, 1, 0,
		0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 5, 56, 0, 0,
		476, 95, 1, 0, 0, 0, 477, 479, 3, 98, 49, 0, 478, 477, 1, 0, 0, 0, 479,
		480, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 97, 1,
		0, 0, 0, 482, 483, 5, 18, 0, 0, 483, 484, 3, 100, 50, 0, 484, 485, 5, 61,
		0, 0, 485, 486, 3, 36, 18, 0, 486, 491, 1, 0, 0, 0, 487, 488, 5, 25, 0,
		0, 488, 489, 5, 61, 0, 0, 489, 491, 3, 36, 18, 0, 490, 482, 1, 0, 0, 0,
		490, 487, 1, 0, 0, 0, 491, 99, 1, 0, 0, 0, 492, 495, 3, 56, 28, 0, 493,
		495, 3, 102, 51, 0, 494, 492, 1, 0, 0, 0, 494, 493, 1, 0, 0, 0, 495, 101,
		1, 0, 0, 0, 496, 497, 3, 166, 83, 0, 497, 498, 5, 64, 0, 0, 498, 499, 3,
		126, 63, 0, 499, 502, 1, 0, 0, 0, 500, 502, 3, 126, 63, 0, 501, 496, 1,
		0, 0, 0, 501, 500, 1, 0, 0, 0, 502, 103, 1, 0, 0, 0, 503, 512, 5, 29, 0,
		0, 504, 512, 3, 106, 53, 0, 505, 512, 3, 108, 54, 0, 506, 512, 3, 110,
		55, 0, 507, 512, 3, 112, 56, 0, 508, 512, 3, 114, 57, 0, 509, 512, 3, 120,
		60, 0, 510, 512, 5, 30, 0, 0, 511, 503, 1, 0, 0, 0, 511, 504, 1, 0, 0,
		0, 511, 505, 1, 0, 0, 0, 511, 506, 1, 0, 0, 0, 511, 507, 1, 0, 0, 0, 511,
		508, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 510, 1, 0, 0, 0, 512, 105,
		1, 0, 0, 0, 513, 514, 5, 57, 0, 0, 514, 515, 5, 31, 0, 0, 515, 516, 5,
		58, 0, 0, 516, 517, 3, 104, 52, 0, 517, 107, 1, 0, 0, 0, 518, 519, 5, 57,
		0, 0, 519, 520, 5, 58, 0, 0, 520, 521, 3, 104, 52, 0, 521, 109, 1, 0, 0,
		0, 522, 523, 5, 9, 0, 0, 523, 524, 5, 57, 0, 0, 524, 525, 3, 104, 52, 0,
		525, 526, 5, 58, 0, 0, 526, 527, 3, 104, 52, 0, 527, 111, 1, 0, 0, 0, 528,
		529, 5, 10, 0, 0, 529, 530, 3, 104, 52, 0, 530, 113, 1, 0, 0, 0, 531, 532,
		5, 7, 0, 0, 532, 534, 5, 55, 0, 0, 533, 535, 3, 116, 58, 0, 534, 533, 1,
		0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 537, 5, 56, 0,
		0, 537, 115, 1, 0, 0, 0, 538, 543, 3, 118, 59, 0, 539, 540, 5, 60, 0, 0,
		540, 542, 3, 118, 59, 0, 541, 539, 1, 0, 0, 0, 542, 545, 1, 0, 0, 0, 543,
		541, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 547, 1, 0, 0, 0, 545, 543,
		1, 0, 0, 0, 546, 548, 5, 60, 0, 0, 547, 546, 1, 0, 0, 0, 547, 548, 1, 0,
		0, 0, 548, 117, 1, 0, 0, 0, 549, 550, 5, 30, 0, 0, 550, 551, 5, 61, 0,
		0, 551, 552, 3, 104, 52, 0, 552, 119, 1, 0, 0, 0, 553, 554, 5, 8, 0, 0,
		554, 556, 5, 55, 0, 0, 555, 557, 3, 122, 61, 0, 556, 555, 1, 0, 0, 0, 556,
		557, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 559, 5, 56, 0, 0, 559, 121,
		1, 0, 0, 0, 560, 565, 3, 124, 62, 0, 561, 562, 5, 60, 0, 0, 562, 564, 3,
		124, 62, 0, 563, 561, 1, 0, 0, 0, 564, 567, 1, 0, 0, 0, 565, 563, 1, 0,
		0, 0, 565, 566, 1, 0, 0, 0, 566, 569, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0,
		568, 570, 5, 60, 0, 0, 569, 568, 1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570,
		123, 1, 0, 0, 0, 571, 572, 5, 30, 0, 0, 572, 573, 3, 26, 13, 0, 573, 125,
		1, 0, 0, 0, 574, 575, 3, 128, 64, 0, 575, 127, 1, 0, 0, 0, 576, 581, 3,
		130, 65, 0, 577, 578, 5, 76, 0, 0, 578, 580, 3, 130, 65, 0, 579, 577, 1,
		0, 0, 0, 580, 583, 1, 0, 0, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0,
		0, 582, 129, 1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 584, 589, 3, 132, 66, 0,
		585, 586, 5, 75, 0, 0, 586, 588, 3, 132, 66, 0, 587, 585, 1, 0, 0, 0, 588,
		591, 1, 0, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 131,
		1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 592, 597, 3, 134, 67, 0, 593, 594, 5,
		73, 0, 0, 594, 596, 3, 134, 67, 0, 595, 593, 1, 0, 0, 0, 596, 599, 1, 0,
		0, 0, 597, 595, 1, 0, 0, 0, 597, 598, 1, 0, 0, 0, 598, 133, 1, 0, 0, 0,
		599, 597, 1, 0, 0, 0, 600, 605, 3, 136, 68, 0, 601, 602, 5, 74, 0, 0, 602,
		604, 3, 136, 68, 0, 603, 601, 1, 0, 0, 0, 604, 607, 1, 0, 0, 0, 605, 603,
		1, 0, 0, 0, 605, 606, 1, 0, 0, 0, 606, 135, 1, 0, 0, 0, 607, 605, 1, 0,
		0, 0, 608, 613, 3, 138, 69, 0, 609, 610, 5, 72, 0, 0, 610, 612, 3, 138,
		69, 0, 611, 609, 1, 0, 0, 0, 612, 615, 1, 0, 0, 0, 613, 611, 1, 0, 0, 0,
		613, 614, 1, 0, 0, 0, 614, 137, 1, 0, 0, 0, 615, 613, 1, 0, 0, 0, 616,
		621, 3, 140, 70, 0, 617, 618, 7, 1, 0, 0, 618, 620, 3, 140, 70, 0, 619,
		617, 1, 0, 0, 0, 620, 623, 1, 0, 0, 0, 621, 619, 1, 0, 0, 0, 621, 622,
		1, 0, 0, 0, 622, 139, 1, 0, 0, 0, 623, 621, 1, 0, 0, 0, 624, 629, 3, 142,
		71, 0, 625, 626, 7, 2, 0, 0, 626, 628, 3, 142, 71, 0, 627, 625, 1, 0, 0,
		0, 628, 631, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 629, 630, 1, 0, 0, 0, 630,
		141, 1, 0, 0, 0, 631, 629, 1, 0, 0, 0, 632, 637, 3, 144, 72, 0, 633, 634,
		7, 3, 0, 0, 634, 636, 3, 144, 72, 0, 635, 633, 1, 0, 0, 0, 636, 639, 1,
		0, 0, 0, 637, 635, 1, 0, 0, 0, 637, 638, 1, 0, 0, 0, 638, 143, 1, 0, 0,
		0, 639, 637, 1, 0, 0, 0, 640, 645, 3, 146, 73, 0, 641, 642, 7, 4, 0, 0,
		642, 644, 3, 146, 73, 0, 643, 641, 1, 0, 0, 0, 644, 647, 1, 0, 0, 0, 645,
		643, 1, 0, 0, 0, 645, 646, 1, 0, 0, 0, 646, 145, 1, 0, 0, 0, 647, 645,
		1, 0, 0, 0, 648, 653, 3, 148, 74, 0, 649, 650, 7, 5, 0, 0, 650, 652, 3,
		148, 74, 0, 651, 649, 1, 0, 0, 0, 652, 655, 1, 0, 0, 0, 653, 651, 1, 0,
		0, 0, 653, 654, 1, 0, 0, 0, 654, 147, 1, 0, 0, 0, 655, 653, 1, 0, 0, 0,
		656, 670, 3, 150, 75, 0, 657, 658, 5, 65, 0, 0, 658, 670, 3, 148, 74, 0,
		659, 660, 5, 66, 0, 0, 660, 670, 3, 148, 74, 0, 661, 662, 5, 77, 0, 0,
		662, 670, 3, 148, 74, 0, 663, 664, 5, 78, 0, 0, 664, 670, 3, 148, 74, 0,
		665, 666, 5, 67, 0, 0, 666, 670, 3, 148, 74, 0, 667, 668, 5, 72, 0, 0,
		668, 670, 3, 148, 74, 0, 669, 656, 1, 0, 0, 0, 669, 657, 1, 0, 0, 0, 669,
		659, 1, 0, 0, 0, 669, 661, 1, 0, 0, 0, 669, 663, 1, 0, 0, 0, 669, 665,
		1, 0, 0, 0, 669, 667, 1, 0, 0, 0, 670, 149, 1, 0, 0, 0, 671, 672, 6, 75,
		-1, 0, 672, 675, 3, 152, 76, 0, 673, 675, 3, 154, 77, 0, 674, 671, 1, 0,
		0, 0, 674, 673, 1, 0, 0, 0, 675, 684, 1, 0, 0, 0, 676, 677, 10, 3, 0, 0,
		677, 683, 3, 156, 78, 0, 678, 679, 10, 2, 0, 0, 679, 683, 3, 158, 79, 0,
		680, 681, 10, 1, 0, 0, 681, 683, 3, 160, 80, 0, 682, 676, 1, 0, 0, 0, 682,
		678, 1, 0, 0, 0, 682, 680, 1, 0, 0, 0, 683, 686, 1, 0, 0, 0, 684, 682,
		1, 0, 0, 0, 684, 685, 1, 0, 0, 0, 685, 151, 1, 0, 0, 0, 686, 684, 1, 0,
		0, 0, 687, 701, 5, 30, 0, 0, 688, 701, 5, 31, 0, 0, 689, 701, 5, 32, 0,
		0, 690, 701, 5, 33, 0, 0, 691, 701, 5, 35, 0, 0, 692, 701, 5, 34, 0, 0,
		693, 694, 5, 53, 0, 0, 694, 695, 3, 126, 63, 0, 695, 696, 5, 54, 0, 0,
		696, 701, 1, 0, 0, 0, 697, 698, 5, 27, 0, 0, 698, 699, 5, 53, 0, 0, 699,
		701, 5, 54, 0, 0, 700, 687, 1, 0, 0, 0, 700, 688, 1, 0, 0, 0, 700, 689,
		1, 0, 0, 0, 700, 690, 1, 0, 0, 0, 700, 691, 1, 0, 0, 0, 700, 692, 1, 0,
		0, 0, 700, 693, 1, 0, 0, 0, 700, 697, 1, 0, 0, 0, 701, 153, 1, 0, 0, 0,
		702, 703, 5, 28, 0, 0, 703, 704, 5, 53, 0, 0, 704, 707, 3, 104, 52, 0,
		705, 706, 5, 59, 0, 0, 706, 708, 3, 166, 83, 0, 707, 705, 1, 0, 0, 0, 707,
		708, 1, 0, 0, 0, 708, 709, 1, 0, 0, 0, 709, 710, 5, 54, 0, 0, 710, 155,
		1, 0, 0, 0, 711, 712, 5, 62, 0, 0, 712, 713, 5, 30, 0, 0, 713, 157, 1,
		0, 0, 0, 714, 715, 5, 57, 0, 0, 715, 716, 3, 126, 63, 0, 716, 717, 5, 58,
		0, 0, 717, 159, 1, 0, 0, 0, 718, 720, 5, 53, 0, 0, 719, 721, 3, 162, 81,
		0, 720, 719, 1, 0, 0, 0, 720, 721, 1, 0, 0, 0, 721, 722, 1, 0, 0, 0, 722,
		723, 5, 54, 0, 0, 723, 161, 1, 0, 0, 0, 724, 729, 3, 126, 63, 0, 725, 726,
		5, 59, 0, 0, 726, 728, 3, 126, 63, 0, 727, 725, 1, 0, 0, 0, 728, 731, 1,
		0, 0, 0, 729, 727, 1, 0, 0, 0, 729, 730, 1, 0, 0, 0, 730, 163, 1, 0, 0,
		0, 731, 729, 1, 0, 0, 0, 732, 737, 5, 30, 0, 0, 733, 734, 5, 59, 0, 0,
		734, 736, 5, 30, 0, 0, 735, 733, 1, 0, 0, 0, 736, 739, 1, 0, 0, 0, 737,
		735, 1, 0, 0, 0, 737, 738, 1, 0, 0, 0, 738, 165, 1, 0, 0, 0, 739, 737,
		1, 0, 0, 0, 740, 745, 3, 126, 63, 0, 741, 742, 5, 59, 0, 0, 742, 744, 3,
		126, 63, 0, 743, 741, 1, 0, 0, 0, 744, 747, 1, 0, 0, 0, 745, 743, 1, 0,
		0, 0, 745, 746, 1, 0, 0, 0, 746, 167, 1, 0, 0, 0, 747, 745, 1, 0, 0, 0,
		72, 170, 175, 186, 197, 202, 207, 214, 222, 226, 232, 236, 244, 248, 257,
		265, 269, 275, 284, 288, 295, 307, 316, 320, 337, 342, 353, 378, 385, 389,
		407, 411, 415, 424, 435, 439, 443, 448, 452, 456, 466, 473, 480, 490, 494,
		501, 511, 534, 543, 547, 556, 565, 569, 581, 589, 597, 605, 613, 621, 629,
		637, 645, 653, 669, 674, 682, 684, 700, 707, 720, 729, 737, 745,
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
	JLangParserKW_MAKE       = 28
	JLangParserTYPE_NAME     = 29
	JLangParserIDENT         = 30
	JLangParserINT_LIT       = 31
	JLangParserFLOAT_LIT     = 32
	JLangParserSTRING_LIT    = 33
	JLangParserRAW_STR       = 34
	JLangParserCHAR_LIT      = 35
	JLangParserCH_SEND       = 36
	JLangParserLE            = 37
	JLangParserGE            = 38
	JLangParserEQ            = 39
	JLangParserNE            = 40
	JLangParserSHLEQ         = 41
	JLangParserSHREQ         = 42
	JLangParserSHL           = 43
	JLangParserSHR           = 44
	JLangParserADDEQ         = 45
	JLangParserSUBEQ         = 46
	JLangParserMULEQ         = 47
	JLangParserDIVEQ         = 48
	JLangParserMODEQ         = 49
	JLangParserANDEQ         = 50
	JLangParserOREQ          = 51
	JLangParserXOREQ         = 52
	JLangParserLPAREN        = 53
	JLangParserRPAREN        = 54
	JLangParserLBRACE        = 55
	JLangParserRBRACE        = 56
	JLangParserLBRACK        = 57
	JLangParserRBRACK        = 58
	JLangParserCOMMA         = 59
	JLangParserSEMI          = 60
	JLangParserCOLON         = 61
	JLangParserDOT           = 62
	JLangParserASSIGN        = 63
	JLangParserDECL          = 64
	JLangParserPLUS          = 65
	JLangParserMINUS         = 66
	JLangParserSTAR          = 67
	JLangParserSLASH         = 68
	JLangParserPERCENT       = 69
	JLangParserLT            = 70
	JLangParserGT            = 71
	JLangParserBAND          = 72
	JLangParserBOR           = 73
	JLangParserBXOR          = 74
	JLangParserANDAND        = 75
	JLangParserOROR          = 76
	JLangParserBANG          = 77
	JLangParserTILDE         = 78
	JLangParserWS            = 79
	JLangParserLINE_COMMENT  = 80
	JLangParserBLOCK_COMMENT = 81
)

// JLangParser rules.
const (
	JLangParserRULE_program            = 0
	JLangParserRULE_pkgDecl            = 1
	JLangParserRULE_impList            = 2
	JLangParserRULE_impDecl            = 3
	JLangParserRULE_importSpec         = 4
	JLangParserRULE_topDecl            = 5
	JLangParserRULE_varDecl            = 6
	JLangParserRULE_varSpec            = 7
	JLangParserRULE_consDecl           = 8
	JLangParserRULE_constSpec          = 9
	JLangParserRULE_typeDecl           = 10
	JLangParserRULE_typeSpec           = 11
	JLangParserRULE_funcDecl           = 12
	JLangParserRULE_signature          = 13
	JLangParserRULE_paramList          = 14
	JLangParserRULE_param              = 15
	JLangParserRULE_resultType         = 16
	JLangParserRULE_block              = 17
	JLangParserRULE_stmtList           = 18
	JLangParserRULE_stmt               = 19
	JLangParserRULE_declStmt           = 20
	JLangParserRULE_labeledStmt        = 21
	JLangParserRULE_simpleStmt         = 22
	JLangParserRULE_exprStmt           = 23
	JLangParserRULE_assignStmt         = 24
	JLangParserRULE_lhs                = 25
	JLangParserRULE_assignOp           = 26
	JLangParserRULE_shortVarDecl       = 27
	JLangParserRULE_sendStmt           = 28
	JLangParserRULE_spawnStmt          = 29
	JLangParserRULE_returnStmt         = 30
	JLangParserRULE_deferStmt          = 31
	JLangParserRULE_breakStmt          = 32
	JLangParserRULE_continueStmt       = 33
	JLangParserRULE_jotoStmt           = 34
	JLangParserRULE_panicStmt          = 35
	JLangParserRULE_ifStmt             = 36
	JLangParserRULE_elseOpt            = 37
	JLangParserRULE_switchStmt         = 38
	JLangParserRULE_switchHead         = 39
	JLangParserRULE_caseClauses        = 40
	JLangParserRULE_caseClause         = 41
	JLangParserRULE_fallOpt            = 42
	JLangParserRULE_forStmt            = 43
	JLangParserRULE_forClause          = 44
	JLangParserRULE_rangeStmt          = 45
	JLangParserRULE_rangeHeader        = 46
	JLangParserRULE_selectStmt         = 47
	JLangParserRULE_commClauses        = 48
	JLangParserRULE_commClause         = 49
	JLangParserRULE_commClauseBody     = 50
	JLangParserRULE_recvStmt           = 51
	JLangParserRULE_type_              = 52
	JLangParserRULE_arrayType          = 53
	JLangParserRULE_sliceType          = 54
	JLangParserRULE_mappingType        = 55
	JLangParserRULE_channelType        = 56
	JLangParserRULE_structType         = 57
	JLangParserRULE_typeFieldList      = 58
	JLangParserRULE_typeField          = 59
	JLangParserRULE_interfaceType      = 60
	JLangParserRULE_methodList         = 61
	JLangParserRULE_methodDecl         = 62
	JLangParserRULE_expr               = 63
	JLangParserRULE_logicalOrExpr      = 64
	JLangParserRULE_logicalAndExpr     = 65
	JLangParserRULE_bitOrExpr          = 66
	JLangParserRULE_bitXorExpr         = 67
	JLangParserRULE_bitAndExpr         = 68
	JLangParserRULE_equalityExpr       = 69
	JLangParserRULE_relationalExpr     = 70
	JLangParserRULE_shiftExpr          = 71
	JLangParserRULE_additiveExpr       = 72
	JLangParserRULE_multiplicativeExpr = 73
	JLangParserRULE_unaryExpr          = 74
	JLangParserRULE_primaryExpr        = 75
	JLangParserRULE_operand            = 76
	JLangParserRULE_makeExpr           = 77
	JLangParserRULE_selector           = 78
	JLangParserRULE_index              = 79
	JLangParserRULE_call               = 80
	JLangParserRULE_argList            = 81
	JLangParserRULE_identList          = 82
	JLangParserRULE_exprList           = 83
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

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.PkgDecl()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_IMP {
		{
			p.SetState(169)
			p.ImpList()
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45036064419805304) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&12423) != 0) {
		{
			p.SetState(172)
			p.TopDecl()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
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

func (s *PkgDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitPkgDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) PkgDecl() (localctx IPkgDeclContext) {
	localctx = NewPkgDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JLangParserRULE_pkgDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(JLangParserKW_PKG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
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

func (s *ImpListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitImpList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ImpList() (localctx IImpListContext) {
	localctx = NewImpListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JLangParserRULE_impList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_IMP {
		{
			p.SetState(183)
			p.ImpDecl()
		}

		p.SetState(186)
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

func (s *ImpDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitImpDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ImpDecl() (localctx IImpDeclContext) {
	localctx = NewImpDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JLangParserRULE_impDecl)
	var _la int

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(JLangParserKW_IMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.ImportSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(JLangParserKW_IMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.ImportSpec()
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == JLangParserCOMMA {
			{
				p.SetState(193)
				p.Match(JLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
				p.ImportSpec()
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(200)
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

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JLangParserRULE_importSpec)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(JLangParserSTRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserIDENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
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

func (s *TopDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitTopDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) TopDecl() (localctx ITopDeclContext) {
	localctx = NewTopDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JLangParserRULE_topDecl)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.VarDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.ConsDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.TypeDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.FuncDecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)
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

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JLangParserRULE_varDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(JLangParserKW_VAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.VarSpec()
	}
	p.SetState(222)
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
				p.SetState(218)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(219)
				p.VarSpec()
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(225)
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

func (s *VarSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitVarSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) VarSpec() (localctx IVarSpecContext) {
	localctx = NewVarSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JLangParserRULE_varSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.IdentList()
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(229)
			p.Type_()
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JLangParserASSIGN {
			{
				p.SetState(230)
				p.Match(JLangParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.ExprList()
			}

		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(234)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
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

func (s *ConsDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitConsDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ConsDecl() (localctx IConsDeclContext) {
	localctx = NewConsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JLangParserRULE_consDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(JLangParserKW_CONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.ConstSpec()
	}
	p.SetState(244)
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
				p.SetState(240)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(241)
				p.ConstSpec()
			}

		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(247)
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

func (s *ConstSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitConstSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ConstSpec() (localctx IConstSpecContext) {
	localctx = NewConstSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JLangParserRULE_constSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.IdentList()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_STRUCT, JLangParserKW_INTERFACE, JLangParserKW_MAPPING, JLangParserKW_CHANNEL, JLangParserTYPE_NAME, JLangParserIDENT, JLangParserLBRACK:
		{
			p.SetState(251)
			p.Type_()
		}
		{
			p.SetState(252)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.ExprList()
		}

	case JLangParserASSIGN:
		{
			p.SetState(255)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
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

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JLangParserRULE_typeDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(JLangParserKW_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.TypeSpec()
	}
	p.SetState(265)
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
				p.SetState(261)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(262)
				p.TypeSpec()
			}

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(268)
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

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JLangParserRULE_typeSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserASSIGN:
		{
			p.SetState(272)
			p.Match(JLangParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Type_()
		}

	case JLangParserKW_STRUCT, JLangParserKW_INTERFACE, JLangParserKW_MAPPING, JLangParserKW_CHANNEL, JLangParserTYPE_NAME, JLangParserIDENT, JLangParserLBRACK:
		{
			p.SetState(274)
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

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JLangParserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(JLangParserKW_DEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Signature()
	}
	{
		p.SetState(280)
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

func (s *SignatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSignature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Signature() (localctx ISignatureContext) {
	localctx = NewSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JLangParserRULE_signature)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(JLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(283)
			p.ParamList()
		}

	}
	{
		p.SetState(286)
		p.Match(JLangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserCOLON {
		{
			p.SetState(287)
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

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JLangParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Param()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(291)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Param()
		}

		p.SetState(297)
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

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JLangParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
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

func (s *ResultTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitResultType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ResultType() (localctx IResultTypeContext) {
	localctx = NewResultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JLangParserRULE_resultType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
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

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JLangParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&45036064419805296) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&12423) != 0) {
		{
			p.SetState(306)
			p.StmtList()
		}

	}
	{
		p.SetState(309)
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

func (s *StmtListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitStmtList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) StmtList() (localctx IStmtListContext) {
	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JLangParserRULE_stmtList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Stmt()
	}
	p.SetState(316)
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
				p.SetState(312)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(313)
				p.Stmt()
			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(319)
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
	PanicStmt() IPanicStmtContext
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

func (s *StmtContext) PanicStmt() IPanicStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPanicStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPanicStmtContext)
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

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JLangParserRULE_stmt)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.SimpleStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.DeclStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.IfStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(325)
			p.SwitchStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(326)
			p.ForStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(327)
			p.RangeStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(328)
			p.SelectStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(329)
			p.DeferStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(330)
			p.ReturnStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(331)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(332)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(333)
			p.JotoStmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(334)
			p.PanicStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(335)
			p.LabeledStmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(336)
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

func (s *DeclStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitDeclStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) DeclStmt() (localctx IDeclStmtContext) {
	localctx = NewDeclStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JLangParserRULE_declStmt)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_VAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.VarDecl()
		}

	case JLangParserKW_CONS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.ConsDecl()
		}

	case JLangParserKW_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
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

func (s *LabeledStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitLabeledStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) LabeledStmt() (localctx ILabeledStmtContext) {
	localctx = NewLabeledStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JLangParserRULE_labeledStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
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

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JLangParserRULE_simpleStmt)
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.ExprStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.AssignStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.ShortVarDecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(351)
			p.SendStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(352)
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

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JLangParserRULE_exprStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Expr()
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

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JLangParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Lhs()
	}
	{
		p.SetState(358)
		p.AssignOp()
	}
	{
		p.SetState(359)
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

func (s *LhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitLhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, JLangParserRULE_lhs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
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

func (s *AssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) AssignOp() (localctx IAssignOpContext) {
	localctx = NewAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, JLangParserRULE_assignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9214393424902356992) != 0) {
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

func (s *ShortVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitShortVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ShortVarDecl() (localctx IShortVarDeclContext) {
	localctx = NewShortVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, JLangParserRULE_shortVarDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.IdentList()
	}
	{
		p.SetState(366)
		p.Match(JLangParserDECL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
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

func (s *SendStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSendStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SendStmt() (localctx ISendStmtContext) {
	localctx = NewSendStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, JLangParserRULE_sendStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Expr()
	}
	{
		p.SetState(370)
		p.Match(JLangParserCH_SEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Expr()
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

func (s *SpawnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSpawnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SpawnStmt() (localctx ISpawnStmtContext) {
	localctx = NewSpawnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, JLangParserRULE_spawnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(JLangParserKW_J)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Expr()
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

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, JLangParserRULE_returnStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(JLangParserKW_RET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(377)
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

func (s *DeferStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitDeferStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) DeferStmt() (localctx IDeferStmtContext) {
	localctx = NewDeferStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, JLangParserRULE_deferStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(JLangParserKW_LATER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Expr()
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

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, JLangParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(JLangParserKW_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(384)
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

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, JLangParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(JLangParserKW_CONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(388)
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

func (s *JotoStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitJotoStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) JotoStmt() (localctx IJotoStmtContext) {
	localctx = NewJotoStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, JLangParserRULE_jotoStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(JLangParserKW_JOTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
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

// IPanicStmtContext is an interface to support dynamic dispatch.
type IPanicStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_PANIC() antlr.TerminalNode
	Expr() IExprContext

	// IsPanicStmtContext differentiates from other interfaces.
	IsPanicStmtContext()
}

type PanicStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPanicStmtContext() *PanicStmtContext {
	var p = new(PanicStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_panicStmt
	return p
}

func InitEmptyPanicStmtContext(p *PanicStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_panicStmt
}

func (*PanicStmtContext) IsPanicStmtContext() {}

func NewPanicStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PanicStmtContext {
	var p = new(PanicStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_panicStmt

	return p
}

func (s *PanicStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PanicStmtContext) KW_PANIC() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_PANIC, 0)
}

func (s *PanicStmtContext) Expr() IExprContext {
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

func (s *PanicStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PanicStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PanicStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterPanicStmt(s)
	}
}

func (s *PanicStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitPanicStmt(s)
	}
}

func (s *PanicStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitPanicStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) PanicStmt() (localctx IPanicStmtContext) {
	localctx = NewPanicStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, JLangParserRULE_panicStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(JLangParserKW_PANIC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Expr()
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

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, JLangParserRULE_ifStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(JLangParserKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Expr()
	}
	{
		p.SetState(399)
		p.Block()
	}
	{
		p.SetState(400)
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

func (s *ElseOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitElseOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ElseOpt() (localctx IElseOptContext) {
	localctx = NewElseOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, JLangParserRULE_elseOpt)
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Match(JLangParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.IfStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(404)
			p.Match(JLangParserKW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
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

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, JLangParserRULE_switchStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(JLangParserKW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&3414808305074683) != 0 {
		{
			p.SetState(410)
			p.SwitchHead()
		}

	}
	{
		p.SetState(413)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(414)
			p.CaseClauses()
		}

	}
	{
		p.SetState(417)
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

func (s *SwitchHeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSwitchHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SwitchHead() (localctx ISwitchHeadContext) {
	localctx = NewSwitchHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, JLangParserRULE_switchHead)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Expr()
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

func (s *CaseClausesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCaseClauses(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) CaseClauses() (localctx ICaseClausesContext) {
	localctx = NewCaseClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, JLangParserRULE_caseClauses)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(421)
			p.CaseClause()
		}

		p.SetState(424)
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

func (s *CaseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCaseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) CaseClause() (localctx ICaseClauseContext) {
	localctx = NewCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, JLangParserRULE_caseClause)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_CASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)
			p.Match(JLangParserKW_CASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.ExprList()
		}
		{
			p.SetState(428)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.StmtList()
		}
		{
			p.SetState(430)
			p.FallOpt()
		}

	case JLangParserKW_DFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(432)
			p.Match(JLangParserKW_DFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
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

func (s *FallOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitFallOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) FallOpt() (localctx IFallOptContext) {
	localctx = NewFallOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, JLangParserRULE_fallOpt)
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_FALL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
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

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, JLangParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(JLangParserKW_FR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1161928771909978112) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&12423) != 0) {
		{
			p.SetState(442)
			p.ForClause()
		}

	}
	{
		p.SetState(445)
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

func (s *ForClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitForClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ForClause() (localctx IForClauseContext) {
	localctx = NewForClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, JLangParserRULE_forClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9007267303131136) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&12423) != 0) {
		{
			p.SetState(447)
			p.SimpleStmt()
		}

	}
	{
		p.SetState(450)
		p.Match(JLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&3414808305074683) != 0 {
		{
			p.SetState(451)
			p.Expr()
		}

	}
	{
		p.SetState(454)
		p.Match(JLangParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9007267303131136) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&12423) != 0) {
		{
			p.SetState(455)
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

func (s *RangeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitRangeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) RangeStmt() (localctx IRangeStmtContext) {
	localctx = NewRangeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, JLangParserRULE_rangeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(JLangParserKW_FR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(459)
		p.Match(JLangParserKW_RANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(460)
		p.RangeHeader()
	}
	{
		p.SetState(461)
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

func (s *RangeHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitRangeHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) RangeHeader() (localctx IRangeHeaderContext) {
	localctx = NewRangeHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, JLangParserRULE_rangeHeader)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(463)
			p.IdentList()
		}
		{
			p.SetState(464)
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
		p.SetState(468)
		p.Expr()
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

func (s *SelectStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSelectStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SelectStmt() (localctx ISelectStmtContext) {
	localctx = NewSelectStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, JLangParserRULE_selectStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		p.Match(JLangParserKW_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(471)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(472)
			p.CommClauses()
		}

	}
	{
		p.SetState(475)
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

func (s *CommClausesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCommClauses(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) CommClauses() (localctx ICommClausesContext) {
	localctx = NewCommClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, JLangParserRULE_commClauses)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JLangParserKW_CASE || _la == JLangParserKW_DFT {
		{
			p.SetState(477)
			p.CommClause()
		}

		p.SetState(480)
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

func (s *CommClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCommClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) CommClause() (localctx ICommClauseContext) {
	localctx = NewCommClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, JLangParserRULE_commClause)
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_CASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(482)
			p.Match(JLangParserKW_CASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.CommClauseBody()
		}
		{
			p.SetState(484)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.StmtList()
		}

	case JLangParserKW_DFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.Match(JLangParserKW_DFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(488)
			p.Match(JLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(489)
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

func (s *CommClauseBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCommClauseBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) CommClauseBody() (localctx ICommClauseBodyContext) {
	localctx = NewCommClauseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, JLangParserRULE_commClauseBody)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(492)
			p.SendStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
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

func (s *RecvStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitRecvStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) RecvStmt() (localctx IRecvStmtContext) {
	localctx = NewRecvStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, JLangParserRULE_recvStmt)
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.ExprList()
		}
		{
			p.SetState(497)
			p.Match(JLangParserDECL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(500)
			p.Expr()
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

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, JLangParserRULE_type_)
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.Match(JLangParserTYPE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(504)
			p.ArrayType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(505)
			p.SliceType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(506)
			p.MappingType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(507)
			p.ChannelType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(508)
			p.StructType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(509)
			p.InterfaceType()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(510)
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

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, JLangParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(JLangParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(516)
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

func (s *SliceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSliceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) SliceType() (localctx ISliceTypeContext) {
	localctx = NewSliceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, JLangParserRULE_sliceType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
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

func (s *MappingTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitMappingType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) MappingType() (localctx IMappingTypeContext) {
	localctx = NewMappingTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, JLangParserRULE_mappingType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(JLangParserKW_MAPPING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
		p.Type_()
	}
	{
		p.SetState(525)
		p.Match(JLangParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(526)
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

func (s *ChannelTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitChannelType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ChannelType() (localctx IChannelTypeContext) {
	localctx = NewChannelTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, JLangParserRULE_channelType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(JLangParserKW_CHANNEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
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

func (s *StructTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitStructType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) StructType() (localctx IStructTypeContext) {
	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, JLangParserRULE_structType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.Match(JLangParserKW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(532)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(533)
			p.TypeFieldList()
		}

	}
	{
		p.SetState(536)
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

func (s *TypeFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitTypeFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) TypeFieldList() (localctx ITypeFieldListContext) {
	localctx = NewTypeFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, JLangParserRULE_typeFieldList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
		p.TypeField()
	}
	p.SetState(543)
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
				p.SetState(539)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(540)
				p.TypeField()
			}

		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(546)
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

func (s *TypeFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitTypeField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) TypeField() (localctx ITypeFieldContext) {
	localctx = NewTypeFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, JLangParserRULE_typeField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)
		p.Match(JLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(551)
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

func (s *InterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) InterfaceType() (localctx IInterfaceTypeContext) {
	localctx = NewInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, JLangParserRULE_interfaceType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		p.Match(JLangParserKW_INTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(554)
		p.Match(JLangParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserIDENT {
		{
			p.SetState(555)
			p.MethodList()
		}

	}
	{
		p.SetState(558)
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

func (s *MethodListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitMethodList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) MethodList() (localctx IMethodListContext) {
	localctx = NewMethodListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, JLangParserRULE_methodList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)
		p.MethodDecl()
	}
	p.SetState(565)
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
				p.SetState(561)
				p.Match(JLangParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(562)
				p.MethodDecl()
			}

		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserSEMI {
		{
			p.SetState(568)
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

func (s *MethodDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitMethodDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) MethodDecl() (localctx IMethodDeclContext) {
	localctx = NewMethodDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, JLangParserRULE_methodDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(572)
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
	LogicalOrExpr() ILogicalOrExprContext

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

func (s *ExprContext) LogicalOrExpr() ILogicalOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExprContext)
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

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, JLangParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(574)
		p.LogicalOrExpr()
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

// ILogicalOrExprContext is an interface to support dynamic dispatch.
type ILogicalOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpr() []ILogicalAndExprContext
	LogicalAndExpr(i int) ILogicalAndExprContext
	AllOROR() []antlr.TerminalNode
	OROR(i int) antlr.TerminalNode

	// IsLogicalOrExprContext differentiates from other interfaces.
	IsLogicalOrExprContext()
}

type LogicalOrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExprContext() *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_logicalOrExpr
	return p
}

func InitEmptyLogicalOrExprContext(p *LogicalOrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_logicalOrExpr
}

func (*LogicalOrExprContext) IsLogicalOrExprContext() {}

func NewLogicalOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_logicalOrExpr

	return p
}

func (s *LogicalOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExprContext) AllLogicalAndExpr() []ILogicalAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExprContext); ok {
			tst[i] = t.(ILogicalAndExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExprContext) LogicalAndExpr(i int) ILogicalAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
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

	return t.(ILogicalAndExprContext)
}

func (s *LogicalOrExprContext) AllOROR() []antlr.TerminalNode {
	return s.GetTokens(JLangParserOROR)
}

func (s *LogicalOrExprContext) OROR(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserOROR, i)
}

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterLogicalOrExpr(s)
	}
}

func (s *LogicalOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitLogicalOrExpr(s)
	}
}

func (s *LogicalOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitLogicalOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) LogicalOrExpr() (localctx ILogicalOrExprContext) {
	localctx = NewLogicalOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, JLangParserRULE_logicalOrExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.LogicalAndExpr()
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserOROR {
		{
			p.SetState(577)
			p.Match(JLangParserOROR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(578)
			p.LogicalAndExpr()
		}

		p.SetState(583)
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

// ILogicalAndExprContext is an interface to support dynamic dispatch.
type ILogicalAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitOrExpr() []IBitOrExprContext
	BitOrExpr(i int) IBitOrExprContext
	AllANDAND() []antlr.TerminalNode
	ANDAND(i int) antlr.TerminalNode

	// IsLogicalAndExprContext differentiates from other interfaces.
	IsLogicalAndExprContext()
}

type LogicalAndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExprContext() *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_logicalAndExpr
	return p
}

func InitEmptyLogicalAndExprContext(p *LogicalAndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_logicalAndExpr
}

func (*LogicalAndExprContext) IsLogicalAndExprContext() {}

func NewLogicalAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_logicalAndExpr

	return p
}

func (s *LogicalAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExprContext) AllBitOrExpr() []IBitOrExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBitOrExprContext); ok {
			len++
		}
	}

	tst := make([]IBitOrExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBitOrExprContext); ok {
			tst[i] = t.(IBitOrExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExprContext) BitOrExpr(i int) IBitOrExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitOrExprContext); ok {
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

	return t.(IBitOrExprContext)
}

func (s *LogicalAndExprContext) AllANDAND() []antlr.TerminalNode {
	return s.GetTokens(JLangParserANDAND)
}

func (s *LogicalAndExprContext) ANDAND(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserANDAND, i)
}

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterLogicalAndExpr(s)
	}
}

func (s *LogicalAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitLogicalAndExpr(s)
	}
}

func (s *LogicalAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitLogicalAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) LogicalAndExpr() (localctx ILogicalAndExprContext) {
	localctx = NewLogicalAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, JLangParserRULE_logicalAndExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.BitOrExpr()
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserANDAND {
		{
			p.SetState(585)
			p.Match(JLangParserANDAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
			p.BitOrExpr()
		}

		p.SetState(591)
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

// IBitOrExprContext is an interface to support dynamic dispatch.
type IBitOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitXorExpr() []IBitXorExprContext
	BitXorExpr(i int) IBitXorExprContext
	AllBOR() []antlr.TerminalNode
	BOR(i int) antlr.TerminalNode

	// IsBitOrExprContext differentiates from other interfaces.
	IsBitOrExprContext()
}

type BitOrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitOrExprContext() *BitOrExprContext {
	var p = new(BitOrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitOrExpr
	return p
}

func InitEmptyBitOrExprContext(p *BitOrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitOrExpr
}

func (*BitOrExprContext) IsBitOrExprContext() {}

func NewBitOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitOrExprContext {
	var p = new(BitOrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_bitOrExpr

	return p
}

func (s *BitOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitOrExprContext) AllBitXorExpr() []IBitXorExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBitXorExprContext); ok {
			len++
		}
	}

	tst := make([]IBitXorExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBitXorExprContext); ok {
			tst[i] = t.(IBitXorExprContext)
			i++
		}
	}

	return tst
}

func (s *BitOrExprContext) BitXorExpr(i int) IBitXorExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitXorExprContext); ok {
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

	return t.(IBitXorExprContext)
}

func (s *BitOrExprContext) AllBOR() []antlr.TerminalNode {
	return s.GetTokens(JLangParserBOR)
}

func (s *BitOrExprContext) BOR(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserBOR, i)
}

func (s *BitOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterBitOrExpr(s)
	}
}

func (s *BitOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitBitOrExpr(s)
	}
}

func (s *BitOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitBitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) BitOrExpr() (localctx IBitOrExprContext) {
	localctx = NewBitOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, JLangParserRULE_bitOrExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(592)
		p.BitXorExpr()
	}
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserBOR {
		{
			p.SetState(593)
			p.Match(JLangParserBOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)
			p.BitXorExpr()
		}

		p.SetState(599)
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

// IBitXorExprContext is an interface to support dynamic dispatch.
type IBitXorExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBitAndExpr() []IBitAndExprContext
	BitAndExpr(i int) IBitAndExprContext
	AllBXOR() []antlr.TerminalNode
	BXOR(i int) antlr.TerminalNode

	// IsBitXorExprContext differentiates from other interfaces.
	IsBitXorExprContext()
}

type BitXorExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitXorExprContext() *BitXorExprContext {
	var p = new(BitXorExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitXorExpr
	return p
}

func InitEmptyBitXorExprContext(p *BitXorExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitXorExpr
}

func (*BitXorExprContext) IsBitXorExprContext() {}

func NewBitXorExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitXorExprContext {
	var p = new(BitXorExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_bitXorExpr

	return p
}

func (s *BitXorExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitXorExprContext) AllBitAndExpr() []IBitAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBitAndExprContext); ok {
			len++
		}
	}

	tst := make([]IBitAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBitAndExprContext); ok {
			tst[i] = t.(IBitAndExprContext)
			i++
		}
	}

	return tst
}

func (s *BitXorExprContext) BitAndExpr(i int) IBitAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitAndExprContext); ok {
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

	return t.(IBitAndExprContext)
}

func (s *BitXorExprContext) AllBXOR() []antlr.TerminalNode {
	return s.GetTokens(JLangParserBXOR)
}

func (s *BitXorExprContext) BXOR(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserBXOR, i)
}

func (s *BitXorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitXorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterBitXorExpr(s)
	}
}

func (s *BitXorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitBitXorExpr(s)
	}
}

func (s *BitXorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitBitXorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) BitXorExpr() (localctx IBitXorExprContext) {
	localctx = NewBitXorExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, JLangParserRULE_bitXorExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(600)
		p.BitAndExpr()
	}
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserBXOR {
		{
			p.SetState(601)
			p.Match(JLangParserBXOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.BitAndExpr()
		}

		p.SetState(607)
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

// IBitAndExprContext is an interface to support dynamic dispatch.
type IBitAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpr() []IEqualityExprContext
	EqualityExpr(i int) IEqualityExprContext
	AllBAND() []antlr.TerminalNode
	BAND(i int) antlr.TerminalNode

	// IsBitAndExprContext differentiates from other interfaces.
	IsBitAndExprContext()
}

type BitAndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitAndExprContext() *BitAndExprContext {
	var p = new(BitAndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitAndExpr
	return p
}

func InitEmptyBitAndExprContext(p *BitAndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_bitAndExpr
}

func (*BitAndExprContext) IsBitAndExprContext() {}

func NewBitAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitAndExprContext {
	var p = new(BitAndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_bitAndExpr

	return p
}

func (s *BitAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BitAndExprContext) AllEqualityExpr() []IEqualityExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExprContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExprContext); ok {
			tst[i] = t.(IEqualityExprContext)
			i++
		}
	}

	return tst
}

func (s *BitAndExprContext) EqualityExpr(i int) IEqualityExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExprContext); ok {
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

	return t.(IEqualityExprContext)
}

func (s *BitAndExprContext) AllBAND() []antlr.TerminalNode {
	return s.GetTokens(JLangParserBAND)
}

func (s *BitAndExprContext) BAND(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserBAND, i)
}

func (s *BitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterBitAndExpr(s)
	}
}

func (s *BitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitBitAndExpr(s)
	}
}

func (s *BitAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitBitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) BitAndExpr() (localctx IBitAndExprContext) {
	localctx = NewBitAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, JLangParserRULE_bitAndExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		p.EqualityExpr()
	}
	p.SetState(613)
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
			{
				p.SetState(609)
				p.Match(JLangParserBAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(610)
				p.EqualityExpr()
			}

		}
		p.SetState(615)
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpr() []IRelationalExprContext
	RelationalExpr(i int) IRelationalExprContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNE() []antlr.TerminalNode
	NE(i int) antlr.TerminalNode

	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_equalityExpr
	return p
}

func InitEmptyEqualityExprContext(p *EqualityExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_equalityExpr
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) AllRelationalExpr() []IRelationalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExprContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExprContext); ok {
			tst[i] = t.(IRelationalExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExprContext) RelationalExpr(i int) IRelationalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExprContext); ok {
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

	return t.(IRelationalExprContext)
}

func (s *EqualityExprContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(JLangParserEQ)
}

func (s *EqualityExprContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserEQ, i)
}

func (s *EqualityExprContext) AllNE() []antlr.TerminalNode {
	return s.GetTokens(JLangParserNE)
}

func (s *EqualityExprContext) NE(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserNE, i)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, JLangParserRULE_equalityExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(616)
		p.RelationalExpr()
	}
	p.SetState(621)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserEQ || _la == JLangParserNE {
		{
			p.SetState(617)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JLangParserEQ || _la == JLangParserNE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(618)
			p.RelationalExpr()
		}

		p.SetState(623)
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

// IRelationalExprContext is an interface to support dynamic dispatch.
type IRelationalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShiftExpr() []IShiftExprContext
	ShiftExpr(i int) IShiftExprContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode

	// IsRelationalExprContext differentiates from other interfaces.
	IsRelationalExprContext()
}

type RelationalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExprContext() *RelationalExprContext {
	var p = new(RelationalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_relationalExpr
	return p
}

func InitEmptyRelationalExprContext(p *RelationalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_relationalExpr
}

func (*RelationalExprContext) IsRelationalExprContext() {}

func NewRelationalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExprContext {
	var p = new(RelationalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_relationalExpr

	return p
}

func (s *RelationalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExprContext) AllShiftExpr() []IShiftExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShiftExprContext); ok {
			len++
		}
	}

	tst := make([]IShiftExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShiftExprContext); ok {
			tst[i] = t.(IShiftExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExprContext) ShiftExpr(i int) IShiftExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExprContext); ok {
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

	return t.(IShiftExprContext)
}

func (s *RelationalExprContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(JLangParserLT)
}

func (s *RelationalExprContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserLT, i)
}

func (s *RelationalExprContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(JLangParserLE)
}

func (s *RelationalExprContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserLE, i)
}

func (s *RelationalExprContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(JLangParserGT)
}

func (s *RelationalExprContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserGT, i)
}

func (s *RelationalExprContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(JLangParserGE)
}

func (s *RelationalExprContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserGE, i)
}

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) RelationalExpr() (localctx IRelationalExprContext) {
	localctx = NewRelationalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, JLangParserRULE_relationalExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(624)
		p.ShiftExpr()
	}
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-37)) & ^0x3f) == 0 && ((int64(1)<<(_la-37))&25769803779) != 0 {
		{
			p.SetState(625)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-37)) & ^0x3f) == 0 && ((int64(1)<<(_la-37))&25769803779) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(626)
			p.ShiftExpr()
		}

		p.SetState(631)
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

// IShiftExprContext is an interface to support dynamic dispatch.
type IShiftExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpr() []IAdditiveExprContext
	AdditiveExpr(i int) IAdditiveExprContext
	AllSHL() []antlr.TerminalNode
	SHL(i int) antlr.TerminalNode
	AllSHR() []antlr.TerminalNode
	SHR(i int) antlr.TerminalNode

	// IsShiftExprContext differentiates from other interfaces.
	IsShiftExprContext()
}

type ShiftExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExprContext() *ShiftExprContext {
	var p = new(ShiftExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_shiftExpr
	return p
}

func InitEmptyShiftExprContext(p *ShiftExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_shiftExpr
}

func (*ShiftExprContext) IsShiftExprContext() {}

func NewShiftExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExprContext {
	var p = new(ShiftExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_shiftExpr

	return p
}

func (s *ShiftExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *ShiftExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
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

	return t.(IAdditiveExprContext)
}

func (s *ShiftExprContext) AllSHL() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSHL)
}

func (s *ShiftExprContext) SHL(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSHL, i)
}

func (s *ShiftExprContext) AllSHR() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSHR)
}

func (s *ShiftExprContext) SHR(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSHR, i)
}

func (s *ShiftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterShiftExpr(s)
	}
}

func (s *ShiftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitShiftExpr(s)
	}
}

func (s *ShiftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitShiftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ShiftExpr() (localctx IShiftExprContext) {
	localctx = NewShiftExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, JLangParserRULE_shiftExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(632)
		p.AdditiveExpr()
	}
	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserSHL || _la == JLangParserSHR {
		{
			p.SetState(633)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JLangParserSHL || _la == JLangParserSHR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(634)
			p.AdditiveExpr()
		}

		p.SetState(639)
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

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpr() []IMultiplicativeExprContext
	MultiplicativeExpr(i int) IMultiplicativeExprContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_additiveExpr
	return p
}

func InitEmptyAdditiveExprContext(p *AdditiveExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_additiveExpr
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
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

	return t.(IMultiplicativeExprContext)
}

func (s *AdditiveExprContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(JLangParserPLUS)
}

func (s *AdditiveExprContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserPLUS, i)
}

func (s *AdditiveExprContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(JLangParserMINUS)
}

func (s *AdditiveExprContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserMINUS, i)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, JLangParserRULE_additiveExpr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)
		p.MultiplicativeExpr()
	}
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(641)
				_la = p.GetTokenStream().LA(1)

				if !(_la == JLangParserPLUS || _la == JLangParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(642)
				p.MultiplicativeExpr()
			}

		}
		p.SetState(647)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpr() []IUnaryExprContext
	UnaryExpr(i int) IUnaryExprContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllPERCENT() []antlr.TerminalNode
	PERCENT(i int) antlr.TerminalNode

	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_multiplicativeExpr
	return p
}

func InitEmptyMultiplicativeExprContext(p *MultiplicativeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_multiplicativeExpr
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) AllUnaryExpr() []IUnaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExprContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExprContext); ok {
			tst[i] = t.(IUnaryExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExprContext) UnaryExpr(i int) IUnaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
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

	return t.(IUnaryExprContext)
}

func (s *MultiplicativeExprContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSTAR)
}

func (s *MultiplicativeExprContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSTAR, i)
}

func (s *MultiplicativeExprContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(JLangParserSLASH)
}

func (s *MultiplicativeExprContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserSLASH, i)
}

func (s *MultiplicativeExprContext) AllPERCENT() []antlr.TerminalNode {
	return s.GetTokens(JLangParserPERCENT)
}

func (s *MultiplicativeExprContext) PERCENT(i int) antlr.TerminalNode {
	return s.GetToken(JLangParserPERCENT, i)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitMultiplicativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, JLangParserRULE_multiplicativeExpr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(648)
		p.UnaryExpr()
	}
	p.SetState(653)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(649)
				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&7) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(650)
				p.UnaryExpr()
			}

		}
		p.SetState(655)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
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
	p.ExitRule()
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

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, JLangParserRULE_unaryExpr)
	p.SetState(669)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_RECOVER, JLangParserKW_MAKE, JLangParserIDENT, JLangParserINT_LIT, JLangParserFLOAT_LIT, JLangParserSTRING_LIT, JLangParserRAW_STR, JLangParserCHAR_LIT, JLangParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(656)
			p.primaryExpr(0)
		}

	case JLangParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(657)
			p.Match(JLangParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(658)
			p.UnaryExpr()
		}

	case JLangParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(659)
			p.Match(JLangParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(660)
			p.UnaryExpr()
		}

	case JLangParserBANG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(661)
			p.Match(JLangParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(662)
			p.UnaryExpr()
		}

	case JLangParserTILDE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(663)
			p.Match(JLangParserTILDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(664)
			p.UnaryExpr()
		}

	case JLangParserSTAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(665)
			p.Match(JLangParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(666)
			p.UnaryExpr()
		}

	case JLangParserBAND:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(667)
			p.Match(JLangParserBAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(668)
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
	MakeExpr() IMakeExprContext
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

func (s *PrimaryExprContext) MakeExpr() IMakeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMakeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMakeExprContext)
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

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
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
	_startState := 150
	p.EnterRecursionRule(localctx, 150, JLangParserRULE_primaryExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserKW_RECOVER, JLangParserIDENT, JLangParserINT_LIT, JLangParserFLOAT_LIT, JLangParserSTRING_LIT, JLangParserRAW_STR, JLangParserCHAR_LIT, JLangParserLPAREN:
		{
			p.SetState(672)
			p.Operand()
		}

	case JLangParserKW_MAKE:
		{
			p.SetState(673)
			p.MakeExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(684)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(682)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(676)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(677)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(678)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(679)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JLangParserRULE_primaryExpr)
				p.SetState(680)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(681)
					p.Call()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(686)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
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

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, JLangParserRULE_operand)
	p.SetState(700)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JLangParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(687)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserINT_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(688)
			p.Match(JLangParserINT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(689)
			p.Match(JLangParserFLOAT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(690)
			p.Match(JLangParserSTRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserCHAR_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(691)
			p.Match(JLangParserCHAR_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserRAW_STR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(692)
			p.Match(JLangParserRAW_STR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserLPAREN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(693)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(694)
			p.Expr()
		}
		{
			p.SetState(695)
			p.Match(JLangParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JLangParserKW_RECOVER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(697)
			p.Match(JLangParserKW_RECOVER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(698)
			p.Match(JLangParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(699)
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

// IMakeExprContext is an interface to support dynamic dispatch.
type IMakeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_MAKE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Type_() IType_Context
	RPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ExprList() IExprListContext

	// IsMakeExprContext differentiates from other interfaces.
	IsMakeExprContext()
}

type MakeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMakeExprContext() *MakeExprContext {
	var p = new(MakeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_makeExpr
	return p
}

func InitEmptyMakeExprContext(p *MakeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JLangParserRULE_makeExpr
}

func (*MakeExprContext) IsMakeExprContext() {}

func NewMakeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MakeExprContext {
	var p = new(MakeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JLangParserRULE_makeExpr

	return p
}

func (s *MakeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MakeExprContext) KW_MAKE() antlr.TerminalNode {
	return s.GetToken(JLangParserKW_MAKE, 0)
}

func (s *MakeExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserLPAREN, 0)
}

func (s *MakeExprContext) Type_() IType_Context {
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

func (s *MakeExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JLangParserRPAREN, 0)
}

func (s *MakeExprContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JLangParserCOMMA, 0)
}

func (s *MakeExprContext) ExprList() IExprListContext {
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

func (s *MakeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MakeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MakeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.EnterMakeExpr(s)
	}
}

func (s *MakeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JLangListener); ok {
		listenerT.ExitMakeExpr(s)
	}
}

func (s *MakeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitMakeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) MakeExpr() (localctx IMakeExprContext) {
	localctx = NewMakeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, JLangParserRULE_makeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(702)
		p.Match(JLangParserKW_MAKE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(703)
		p.Match(JLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(704)
		p.Type_()
	}
	p.SetState(707)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JLangParserCOMMA {
		{
			p.SetState(705)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(706)
			p.ExprList()
		}

	}
	{
		p.SetState(709)
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

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, JLangParserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(711)
		p.Match(JLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(712)
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

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, JLangParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(714)
		p.Match(JLangParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(715)
		p.Expr()
	}
	{
		p.SetState(716)
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

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, JLangParserRULE_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(718)
		p.Match(JLangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(720)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&3414808305074683) != 0 {
		{
			p.SetState(719)
			p.ArgList()
		}

	}
	{
		p.SetState(722)
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

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, JLangParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(724)
		p.Expr()
	}
	p.SetState(729)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(725)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(726)
			p.Expr()
		}

		p.SetState(731)
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

func (s *IdentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitIdentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) IdentList() (localctx IIdentListContext) {
	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, JLangParserRULE_identList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(732)
		p.Match(JLangParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(737)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(733)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(734)
			p.Match(JLangParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(739)
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

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JLangVisitor:
		return t.VisitExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JLangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, JLangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(740)
		p.Expr()
	}
	p.SetState(745)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JLangParserCOMMA {
		{
			p.SetState(741)
			p.Match(JLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(742)
			p.Expr()
		}

		p.SetState(747)
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
	case 75:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JLangParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
