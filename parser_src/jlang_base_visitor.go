// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser_src // JLang
import "github.com/antlr4-go/antlr/v4"

type BaseJLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitPkgDecl(ctx *PkgDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitImpList(ctx *ImpListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitImpDecl(ctx *ImpDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitTopDecl(ctx *TopDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitVarSpec(ctx *VarSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitConsDecl(ctx *ConsDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitConstSpec(ctx *ConstSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSignature(ctx *SignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitResultType(ctx *ResultTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitStmtList(ctx *StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitDeclStmt(ctx *DeclStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitLabeledStmt(ctx *LabeledStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitLhs(ctx *LhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitAssignOp(ctx *AssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitShortVarDecl(ctx *ShortVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSendStmt(ctx *SendStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSpawnStmt(ctx *SpawnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitJotoStmt(ctx *JotoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitPanicStmt(ctx *PanicStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitElseOpt(ctx *ElseOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSwitchHead(ctx *SwitchHeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitFallOpt(ctx *FallOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitRangeStmt(ctx *RangeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitRangeHeader(ctx *RangeHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCommClauses(ctx *CommClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCommClause(ctx *CommClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCommClauseBody(ctx *CommClauseBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitRecvStmt(ctx *RecvStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSliceType(ctx *SliceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitMappingType(ctx *MappingTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitChannelType(ctx *ChannelTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitTypeFieldList(ctx *TypeFieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitTypeField(ctx *TypeFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitMethodList(ctx *MethodListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitMethodDecl(ctx *MethodDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitLogicalOrExpr(ctx *LogicalOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitLogicalAndExpr(ctx *LogicalAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitBitOrExpr(ctx *BitOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitBitXorExpr(ctx *BitXorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitBitAndExpr(ctx *BitAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitShiftExpr(ctx *ShiftExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitMakeExpr(ctx *MakeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJLangVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}
