// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser_src // JLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by JLangParser.
type JLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by JLangParser#pkgDecl.
	VisitPkgDecl(ctx *PkgDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#impList.
	VisitImpList(ctx *ImpListContext) interface{}

	// Visit a parse tree produced by JLangParser#impDecl.
	VisitImpDecl(ctx *ImpDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by JLangParser#topDecl.
	VisitTopDecl(ctx *TopDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#varSpec.
	VisitVarSpec(ctx *VarSpecContext) interface{}

	// Visit a parse tree produced by JLangParser#consDecl.
	VisitConsDecl(ctx *ConsDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#constSpec.
	VisitConstSpec(ctx *ConstSpecContext) interface{}

	// Visit a parse tree produced by JLangParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by JLangParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by JLangParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by JLangParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by JLangParser#resultType.
	VisitResultType(ctx *ResultTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by JLangParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by JLangParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by JLangParser#declStmt.
	VisitDeclStmt(ctx *DeclStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#labeledStmt.
	VisitLabeledStmt(ctx *LabeledStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#lhs.
	VisitLhs(ctx *LhsContext) interface{}

	// Visit a parse tree produced by JLangParser#assignOp.
	VisitAssignOp(ctx *AssignOpContext) interface{}

	// Visit a parse tree produced by JLangParser#shortVarDecl.
	VisitShortVarDecl(ctx *ShortVarDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#sendStmt.
	VisitSendStmt(ctx *SendStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#spawnStmt.
	VisitSpawnStmt(ctx *SpawnStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#deferStmt.
	VisitDeferStmt(ctx *DeferStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#jotoStmt.
	VisitJotoStmt(ctx *JotoStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#panicStmt.
	VisitPanicStmt(ctx *PanicStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#elseOpt.
	VisitElseOpt(ctx *ElseOptContext) interface{}

	// Visit a parse tree produced by JLangParser#switchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#switchHead.
	VisitSwitchHead(ctx *SwitchHeadContext) interface{}

	// Visit a parse tree produced by JLangParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by JLangParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by JLangParser#fallOpt.
	VisitFallOpt(ctx *FallOptContext) interface{}

	// Visit a parse tree produced by JLangParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by JLangParser#rangeStmt.
	VisitRangeStmt(ctx *RangeStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#rangeHeader.
	VisitRangeHeader(ctx *RangeHeaderContext) interface{}

	// Visit a parse tree produced by JLangParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#commClauses.
	VisitCommClauses(ctx *CommClausesContext) interface{}

	// Visit a parse tree produced by JLangParser#commClause.
	VisitCommClause(ctx *CommClauseContext) interface{}

	// Visit a parse tree produced by JLangParser#commClauseBody.
	VisitCommClauseBody(ctx *CommClauseBodyContext) interface{}

	// Visit a parse tree produced by JLangParser#recvStmt.
	VisitRecvStmt(ctx *RecvStmtContext) interface{}

	// Visit a parse tree produced by JLangParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by JLangParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#sliceType.
	VisitSliceType(ctx *SliceTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#mappingType.
	VisitMappingType(ctx *MappingTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#channelType.
	VisitChannelType(ctx *ChannelTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#typeFieldList.
	VisitTypeFieldList(ctx *TypeFieldListContext) interface{}

	// Visit a parse tree produced by JLangParser#typeField.
	VisitTypeField(ctx *TypeFieldContext) interface{}

	// Visit a parse tree produced by JLangParser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by JLangParser#methodList.
	VisitMethodList(ctx *MethodListContext) interface{}

	// Visit a parse tree produced by JLangParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by JLangParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by JLangParser#logicalOrExpr.
	VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{}

	// Visit a parse tree produced by JLangParser#logicalAndExpr.
	VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{}

	// Visit a parse tree produced by JLangParser#bitOrExpr.
	VisitBitOrExpr(ctx *BitOrExprContext) interface{}

	// Visit a parse tree produced by JLangParser#bitXorExpr.
	VisitBitXorExpr(ctx *BitXorExprContext) interface{}

	// Visit a parse tree produced by JLangParser#bitAndExpr.
	VisitBitAndExpr(ctx *BitAndExprContext) interface{}

	// Visit a parse tree produced by JLangParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by JLangParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by JLangParser#shiftExpr.
	VisitShiftExpr(ctx *ShiftExprContext) interface{}

	// Visit a parse tree produced by JLangParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by JLangParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by JLangParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by JLangParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by JLangParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by JLangParser#makeExpr.
	VisitMakeExpr(ctx *MakeExprContext) interface{}

	// Visit a parse tree produced by JLangParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by JLangParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by JLangParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by JLangParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by JLangParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

	// Visit a parse tree produced by JLangParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}
}
