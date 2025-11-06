// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // JLang
import "github.com/antlr4-go/antlr/v4"

// BaseJLangListener is a complete listener for a parse tree produced by JLangParser.
type BaseJLangListener struct{}

var _ JLangListener = &BaseJLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseJLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseJLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterPkgDecl is called when production pkgDecl is entered.
func (s *BaseJLangListener) EnterPkgDecl(ctx *PkgDeclContext) {}

// ExitPkgDecl is called when production pkgDecl is exited.
func (s *BaseJLangListener) ExitPkgDecl(ctx *PkgDeclContext) {}

// EnterImpList is called when production impList is entered.
func (s *BaseJLangListener) EnterImpList(ctx *ImpListContext) {}

// ExitImpList is called when production impList is exited.
func (s *BaseJLangListener) ExitImpList(ctx *ImpListContext) {}

// EnterImpDecl is called when production impDecl is entered.
func (s *BaseJLangListener) EnterImpDecl(ctx *ImpDeclContext) {}

// ExitImpDecl is called when production impDecl is exited.
func (s *BaseJLangListener) ExitImpDecl(ctx *ImpDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseJLangListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseJLangListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterTopDecl is called when production topDecl is entered.
func (s *BaseJLangListener) EnterTopDecl(ctx *TopDeclContext) {}

// ExitTopDecl is called when production topDecl is exited.
func (s *BaseJLangListener) ExitTopDecl(ctx *TopDeclContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseJLangListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseJLangListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarSpec is called when production varSpec is entered.
func (s *BaseJLangListener) EnterVarSpec(ctx *VarSpecContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *BaseJLangListener) ExitVarSpec(ctx *VarSpecContext) {}

// EnterConsDecl is called when production consDecl is entered.
func (s *BaseJLangListener) EnterConsDecl(ctx *ConsDeclContext) {}

// ExitConsDecl is called when production consDecl is exited.
func (s *BaseJLangListener) ExitConsDecl(ctx *ConsDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *BaseJLangListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *BaseJLangListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseJLangListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseJLangListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseJLangListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseJLangListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseJLangListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseJLangListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterSignature is called when production signature is entered.
func (s *BaseJLangListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *BaseJLangListener) ExitSignature(ctx *SignatureContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseJLangListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseJLangListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseJLangListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseJLangListener) ExitParam(ctx *ParamContext) {}

// EnterResultType is called when production resultType is entered.
func (s *BaseJLangListener) EnterResultType(ctx *ResultTypeContext) {}

// ExitResultType is called when production resultType is exited.
func (s *BaseJLangListener) ExitResultType(ctx *ResultTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJLangListener) ExitBlock(ctx *BlockContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BaseJLangListener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BaseJLangListener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseJLangListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseJLangListener) ExitStmt(ctx *StmtContext) {}

// EnterDeclStmt is called when production declStmt is entered.
func (s *BaseJLangListener) EnterDeclStmt(ctx *DeclStmtContext) {}

// ExitDeclStmt is called when production declStmt is exited.
func (s *BaseJLangListener) ExitDeclStmt(ctx *DeclStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *BaseJLangListener) EnterLabeledStmt(ctx *LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *BaseJLangListener) ExitLabeledStmt(ctx *LabeledStmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseJLangListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseJLangListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseJLangListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseJLangListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseJLangListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseJLangListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BaseJLangListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BaseJLangListener) ExitLhs(ctx *LhsContext) {}

// EnterAssignOp is called when production assignOp is entered.
func (s *BaseJLangListener) EnterAssignOp(ctx *AssignOpContext) {}

// ExitAssignOp is called when production assignOp is exited.
func (s *BaseJLangListener) ExitAssignOp(ctx *AssignOpContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BaseJLangListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BaseJLangListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *BaseJLangListener) EnterSendStmt(ctx *SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *BaseJLangListener) ExitSendStmt(ctx *SendStmtContext) {}

// EnterSpawnStmt is called when production spawnStmt is entered.
func (s *BaseJLangListener) EnterSpawnStmt(ctx *SpawnStmtContext) {}

// ExitSpawnStmt is called when production spawnStmt is exited.
func (s *BaseJLangListener) ExitSpawnStmt(ctx *SpawnStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseJLangListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseJLangListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseJLangListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseJLangListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseJLangListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseJLangListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseJLangListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseJLangListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterJotoStmt is called when production jotoStmt is entered.
func (s *BaseJLangListener) EnterJotoStmt(ctx *JotoStmtContext) {}

// ExitJotoStmt is called when production jotoStmt is exited.
func (s *BaseJLangListener) ExitJotoStmt(ctx *JotoStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseJLangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseJLangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterElseOpt is called when production elseOpt is entered.
func (s *BaseJLangListener) EnterElseOpt(ctx *ElseOptContext) {}

// ExitElseOpt is called when production elseOpt is exited.
func (s *BaseJLangListener) ExitElseOpt(ctx *ElseOptContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseJLangListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseJLangListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterSwitchHead is called when production switchHead is entered.
func (s *BaseJLangListener) EnterSwitchHead(ctx *SwitchHeadContext) {}

// ExitSwitchHead is called when production switchHead is exited.
func (s *BaseJLangListener) ExitSwitchHead(ctx *SwitchHeadContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseJLangListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseJLangListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseJLangListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseJLangListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterFallOpt is called when production fallOpt is entered.
func (s *BaseJLangListener) EnterFallOpt(ctx *FallOptContext) {}

// ExitFallOpt is called when production fallOpt is exited.
func (s *BaseJLangListener) ExitFallOpt(ctx *FallOptContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseJLangListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseJLangListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseJLangListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseJLangListener) ExitForClause(ctx *ForClauseContext) {}

// EnterRangeStmt is called when production rangeStmt is entered.
func (s *BaseJLangListener) EnterRangeStmt(ctx *RangeStmtContext) {}

// ExitRangeStmt is called when production rangeStmt is exited.
func (s *BaseJLangListener) ExitRangeStmt(ctx *RangeStmtContext) {}

// EnterRangeHeader is called when production rangeHeader is entered.
func (s *BaseJLangListener) EnterRangeHeader(ctx *RangeHeaderContext) {}

// ExitRangeHeader is called when production rangeHeader is exited.
func (s *BaseJLangListener) ExitRangeHeader(ctx *RangeHeaderContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseJLangListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseJLangListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterCommClauses is called when production commClauses is entered.
func (s *BaseJLangListener) EnterCommClauses(ctx *CommClausesContext) {}

// ExitCommClauses is called when production commClauses is exited.
func (s *BaseJLangListener) ExitCommClauses(ctx *CommClausesContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *BaseJLangListener) EnterCommClause(ctx *CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *BaseJLangListener) ExitCommClause(ctx *CommClauseContext) {}

// EnterCommClauseBody is called when production commClauseBody is entered.
func (s *BaseJLangListener) EnterCommClauseBody(ctx *CommClauseBodyContext) {}

// ExitCommClauseBody is called when production commClauseBody is exited.
func (s *BaseJLangListener) ExitCommClauseBody(ctx *CommClauseBodyContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *BaseJLangListener) EnterRecvStmt(ctx *RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *BaseJLangListener) ExitRecvStmt(ctx *RecvStmtContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseJLangListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseJLangListener) ExitType_(ctx *Type_Context) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseJLangListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseJLangListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BaseJLangListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BaseJLangListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterMappingType is called when production mappingType is entered.
func (s *BaseJLangListener) EnterMappingType(ctx *MappingTypeContext) {}

// ExitMappingType is called when production mappingType is exited.
func (s *BaseJLangListener) ExitMappingType(ctx *MappingTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *BaseJLangListener) EnterChannelType(ctx *ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *BaseJLangListener) ExitChannelType(ctx *ChannelTypeContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseJLangListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseJLangListener) ExitStructType(ctx *StructTypeContext) {}

// EnterTypeFieldList is called when production typeFieldList is entered.
func (s *BaseJLangListener) EnterTypeFieldList(ctx *TypeFieldListContext) {}

// ExitTypeFieldList is called when production typeFieldList is exited.
func (s *BaseJLangListener) ExitTypeFieldList(ctx *TypeFieldListContext) {}

// EnterTypeField is called when production typeField is entered.
func (s *BaseJLangListener) EnterTypeField(ctx *TypeFieldContext) {}

// ExitTypeField is called when production typeField is exited.
func (s *BaseJLangListener) ExitTypeField(ctx *TypeFieldContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseJLangListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseJLangListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterMethodList is called when production methodList is entered.
func (s *BaseJLangListener) EnterMethodList(ctx *MethodListContext) {}

// ExitMethodList is called when production methodList is exited.
func (s *BaseJLangListener) ExitMethodList(ctx *MethodListContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseJLangListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseJLangListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseJLangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseJLangListener) ExitExpr(ctx *ExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseJLangListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseJLangListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseJLangListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseJLangListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseJLangListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseJLangListener) ExitOperand(ctx *OperandContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseJLangListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseJLangListener) ExitSelector(ctx *SelectorContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseJLangListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseJLangListener) ExitIndex(ctx *IndexContext) {}

// EnterCall is called when production call is entered.
func (s *BaseJLangListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseJLangListener) ExitCall(ctx *CallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseJLangListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseJLangListener) ExitArgList(ctx *ArgListContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseJLangListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseJLangListener) ExitIdentList(ctx *IdentListContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseJLangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseJLangListener) ExitExprList(ctx *ExprListContext) {}
