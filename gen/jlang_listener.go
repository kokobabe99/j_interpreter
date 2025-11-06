// Code generated from JLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // JLang
import "github.com/antlr4-go/antlr/v4"

// JLangListener is a complete listener for a parse tree produced by JLangParser.
type JLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterPkgDecl is called when entering the pkgDecl production.
	EnterPkgDecl(c *PkgDeclContext)

	// EnterImpList is called when entering the impList production.
	EnterImpList(c *ImpListContext)

	// EnterImpDecl is called when entering the impDecl production.
	EnterImpDecl(c *ImpDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterTopDecl is called when entering the topDecl production.
	EnterTopDecl(c *TopDeclContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterVarSpec is called when entering the varSpec production.
	EnterVarSpec(c *VarSpecContext)

	// EnterConsDecl is called when entering the consDecl production.
	EnterConsDecl(c *ConsDeclContext)

	// EnterConstSpec is called when entering the constSpec production.
	EnterConstSpec(c *ConstSpecContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterSignature is called when entering the signature production.
	EnterSignature(c *SignatureContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterResultType is called when entering the resultType production.
	EnterResultType(c *ResultTypeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclStmt is called when entering the declStmt production.
	EnterDeclStmt(c *DeclStmtContext)

	// EnterLabeledStmt is called when entering the labeledStmt production.
	EnterLabeledStmt(c *LabeledStmtContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterAssignOp is called when entering the assignOp production.
	EnterAssignOp(c *AssignOpContext)

	// EnterShortVarDecl is called when entering the shortVarDecl production.
	EnterShortVarDecl(c *ShortVarDeclContext)

	// EnterSendStmt is called when entering the sendStmt production.
	EnterSendStmt(c *SendStmtContext)

	// EnterSpawnStmt is called when entering the spawnStmt production.
	EnterSpawnStmt(c *SpawnStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterDeferStmt is called when entering the deferStmt production.
	EnterDeferStmt(c *DeferStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterJotoStmt is called when entering the jotoStmt production.
	EnterJotoStmt(c *JotoStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterElseOpt is called when entering the elseOpt production.
	EnterElseOpt(c *ElseOptContext)

	// EnterSwitchStmt is called when entering the switchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchHead is called when entering the switchHead production.
	EnterSwitchHead(c *SwitchHeadContext)

	// EnterCaseClauses is called when entering the caseClauses production.
	EnterCaseClauses(c *CaseClausesContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterFallOpt is called when entering the fallOpt production.
	EnterFallOpt(c *FallOptContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterForClause is called when entering the forClause production.
	EnterForClause(c *ForClauseContext)

	// EnterRangeStmt is called when entering the rangeStmt production.
	EnterRangeStmt(c *RangeStmtContext)

	// EnterRangeHeader is called when entering the rangeHeader production.
	EnterRangeHeader(c *RangeHeaderContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterCommClauses is called when entering the commClauses production.
	EnterCommClauses(c *CommClausesContext)

	// EnterCommClause is called when entering the commClause production.
	EnterCommClause(c *CommClauseContext)

	// EnterCommClauseBody is called when entering the commClauseBody production.
	EnterCommClauseBody(c *CommClauseBodyContext)

	// EnterRecvStmt is called when entering the recvStmt production.
	EnterRecvStmt(c *RecvStmtContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterSliceType is called when entering the sliceType production.
	EnterSliceType(c *SliceTypeContext)

	// EnterMappingType is called when entering the mappingType production.
	EnterMappingType(c *MappingTypeContext)

	// EnterChannelType is called when entering the channelType production.
	EnterChannelType(c *ChannelTypeContext)

	// EnterStructType is called when entering the structType production.
	EnterStructType(c *StructTypeContext)

	// EnterTypeFieldList is called when entering the typeFieldList production.
	EnterTypeFieldList(c *TypeFieldListContext)

	// EnterTypeField is called when entering the typeField production.
	EnterTypeField(c *TypeFieldContext)

	// EnterInterfaceType is called when entering the interfaceType production.
	EnterInterfaceType(c *InterfaceTypeContext)

	// EnterMethodList is called when entering the methodList production.
	EnterMethodList(c *MethodListContext)

	// EnterMethodDecl is called when entering the methodDecl production.
	EnterMethodDecl(c *MethodDeclContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitPkgDecl is called when exiting the pkgDecl production.
	ExitPkgDecl(c *PkgDeclContext)

	// ExitImpList is called when exiting the impList production.
	ExitImpList(c *ImpListContext)

	// ExitImpDecl is called when exiting the impDecl production.
	ExitImpDecl(c *ImpDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitTopDecl is called when exiting the topDecl production.
	ExitTopDecl(c *TopDeclContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitVarSpec is called when exiting the varSpec production.
	ExitVarSpec(c *VarSpecContext)

	// ExitConsDecl is called when exiting the consDecl production.
	ExitConsDecl(c *ConsDeclContext)

	// ExitConstSpec is called when exiting the constSpec production.
	ExitConstSpec(c *ConstSpecContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitSignature is called when exiting the signature production.
	ExitSignature(c *SignatureContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitResultType is called when exiting the resultType production.
	ExitResultType(c *ResultTypeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclStmt is called when exiting the declStmt production.
	ExitDeclStmt(c *DeclStmtContext)

	// ExitLabeledStmt is called when exiting the labeledStmt production.
	ExitLabeledStmt(c *LabeledStmtContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitAssignOp is called when exiting the assignOp production.
	ExitAssignOp(c *AssignOpContext)

	// ExitShortVarDecl is called when exiting the shortVarDecl production.
	ExitShortVarDecl(c *ShortVarDeclContext)

	// ExitSendStmt is called when exiting the sendStmt production.
	ExitSendStmt(c *SendStmtContext)

	// ExitSpawnStmt is called when exiting the spawnStmt production.
	ExitSpawnStmt(c *SpawnStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitDeferStmt is called when exiting the deferStmt production.
	ExitDeferStmt(c *DeferStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitJotoStmt is called when exiting the jotoStmt production.
	ExitJotoStmt(c *JotoStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitElseOpt is called when exiting the elseOpt production.
	ExitElseOpt(c *ElseOptContext)

	// ExitSwitchStmt is called when exiting the switchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchHead is called when exiting the switchHead production.
	ExitSwitchHead(c *SwitchHeadContext)

	// ExitCaseClauses is called when exiting the caseClauses production.
	ExitCaseClauses(c *CaseClausesContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitFallOpt is called when exiting the fallOpt production.
	ExitFallOpt(c *FallOptContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitForClause is called when exiting the forClause production.
	ExitForClause(c *ForClauseContext)

	// ExitRangeStmt is called when exiting the rangeStmt production.
	ExitRangeStmt(c *RangeStmtContext)

	// ExitRangeHeader is called when exiting the rangeHeader production.
	ExitRangeHeader(c *RangeHeaderContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitCommClauses is called when exiting the commClauses production.
	ExitCommClauses(c *CommClausesContext)

	// ExitCommClause is called when exiting the commClause production.
	ExitCommClause(c *CommClauseContext)

	// ExitCommClauseBody is called when exiting the commClauseBody production.
	ExitCommClauseBody(c *CommClauseBodyContext)

	// ExitRecvStmt is called when exiting the recvStmt production.
	ExitRecvStmt(c *RecvStmtContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitSliceType is called when exiting the sliceType production.
	ExitSliceType(c *SliceTypeContext)

	// ExitMappingType is called when exiting the mappingType production.
	ExitMappingType(c *MappingTypeContext)

	// ExitChannelType is called when exiting the channelType production.
	ExitChannelType(c *ChannelTypeContext)

	// ExitStructType is called when exiting the structType production.
	ExitStructType(c *StructTypeContext)

	// ExitTypeFieldList is called when exiting the typeFieldList production.
	ExitTypeFieldList(c *TypeFieldListContext)

	// ExitTypeField is called when exiting the typeField production.
	ExitTypeField(c *TypeFieldContext)

	// ExitInterfaceType is called when exiting the interfaceType production.
	ExitInterfaceType(c *InterfaceTypeContext)

	// ExitMethodList is called when exiting the methodList production.
	ExitMethodList(c *MethodListContext)

	// ExitMethodDecl is called when exiting the methodDecl production.
	ExitMethodDecl(c *MethodDeclContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)
}
