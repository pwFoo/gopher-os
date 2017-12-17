package entity

import "testing"

func TestOpcodeToString(t *testing.T) {
	opcodeList := []AMLOpcode{
		OpZero,
		OpOne,
		OpAlias,
		OpName,
		OpBytePrefix,
		OpWordPrefix,
		OpDwordPrefix,
		OpStringPrefix,
		OpQwordPrefix,
		OpScope,
		OpBuffer,
		OpPackage,
		OpVarPackage,
		OpMethod,
		OpExternal,
		OpLocal0,
		OpLocal1,
		OpLocal2,
		OpLocal3,
		OpLocal4,
		OpLocal5,
		OpLocal6,
		OpLocal7,
		OpArg0,
		OpArg1,
		OpArg2,
		OpArg3,
		OpArg4,
		OpArg5,
		OpArg6,
		OpStore,
		OpRefOf,
		OpAdd,
		OpConcat,
		OpSubtract,
		OpIncrement,
		OpDecrement,
		OpMultiply,
		OpDivide,
		OpShiftLeft,
		OpShiftRight,
		OpAnd,
		OpNand,
		OpOr,
		OpNor,
		OpXor,
		OpNot,
		OpFindSetLeftBit,
		OpFindSetRightBit,
		OpDerefOf,
		OpConcatRes,
		OpMod,
		OpNotify,
		OpSizeOf,
		OpIndex,
		OpMatch,
		OpCreateDWordField,
		OpCreateWordField,
		OpCreateByteField,
		OpCreateBitField,
		OpObjectType,
		OpCreateQWordField,
		OpLand,
		OpLor,
		OpLnot,
		OpLEqual,
		OpLGreater,
		OpLLess,
		OpToBuffer,
		OpToDecimalString,
		OpToHexString,
		OpToInteger,
		OpToString,
		OpCopyObject,
		OpMid,
		OpContinue,
		OpIf,
		OpElse,
		OpWhile,
		OpNoop,
		OpReturn,
		OpBreak,
		OpBreakPoint,
		OpOnes,
		OpMutex,
		OpEvent,
		OpCondRefOf,
		OpCreateField,
		OpLoadTable,
		OpLoad,
		OpStall,
		OpSleep,
		OpAcquire,
		OpSignal,
		OpWait,
		OpReset,
		OpRelease,
		OpFromBCD,
		OpToBCD,
		OpUnload,
		OpRevision,
		OpDebug,
		OpFatal,
		OpTimer,
		OpOpRegion,
		OpField,
		OpDevice,
		OpProcessor,
		OpPowerRes,
		OpThermalZone,
		OpIndexField,
		OpBankField,
		OpDataRegion,
	}

	for specIndex, op := range opcodeList {
		if op.String() == "unknown" {
			t.Errorf("[spec %d] op 0x%x String() returned \"unknown\"", specIndex, op)
		}
	}

	// Also test invalid opcode
	if got := AMLOpcode(0xffff).String(); got != "unknown" {
		t.Fatalf("expected String() for invalid opcode to return \"unknown\"; got: %q", got)
	}
}

func TestOpcodeIsX(t *testing.T) {
	specs := []struct {
		op     AMLOpcode
		testFn func(AMLOpcode) bool
		want   bool
	}{
		// OpIsLocalArg
		{OpLocal0, OpIsLocalArg, true},
		{OpLocal1, OpIsLocalArg, true},
		{OpLocal2, OpIsLocalArg, true},
		{OpLocal3, OpIsLocalArg, true},
		{OpLocal4, OpIsLocalArg, true},
		{OpLocal5, OpIsLocalArg, true},
		{OpLocal6, OpIsLocalArg, true},
		{OpLocal7, OpIsLocalArg, true},
		{OpArg0, OpIsLocalArg, false},
		{OpDivide, OpIsLocalArg, false},
		// OpIsMethodArg
		{OpArg0, OpIsMethodArg, true},
		{OpArg1, OpIsMethodArg, true},
		{OpArg2, OpIsMethodArg, true},
		{OpArg3, OpIsMethodArg, true},
		{OpArg4, OpIsMethodArg, true},
		{OpArg5, OpIsMethodArg, true},
		{OpArg6, OpIsMethodArg, true},
		{OpLocal7, OpIsMethodArg, false},
		{OpIf, OpIsMethodArg, false},
		// OpIsArg
		{OpLocal5, OpIsArg, true},
		{OpArg1, OpIsArg, true},
		{OpDivide, OpIsArg, false},
		// OpIsType2
		{OpAcquire, OpIsType2, true},
		{OpAdd, OpIsType2, true},
		{OpAnd, OpIsType2, true},
		{OpBuffer, OpIsType2, true},
		{OpConcat, OpIsType2, true},
		{OpConcatRes, OpIsType2, true},
		{OpCondRefOf, OpIsType2, true},
		{OpCopyObject, OpIsType2, true},
		{OpDecrement, OpIsType2, true},
		{OpDerefOf, OpIsType2, true},
		{OpDivide, OpIsType2, true},
		{OpFindSetLeftBit, OpIsType2, true},
		{OpFindSetRightBit, OpIsType2, true},
		{OpFromBCD, OpIsType2, true},
		{OpIncrement, OpIsType2, true},
		{OpIndex, OpIsType2, true},
		{OpLand, OpIsType2, true},
		{OpLEqual, OpIsType2, true},
		{OpLGreater, OpIsType2, true},
		{OpLLess, OpIsType2, true},
		{OpMid, OpIsType2, true},
		{OpLnot, OpIsType2, true},
		{OpLoadTable, OpIsType2, true},
		{OpLor, OpIsType2, true},
		{OpMatch, OpIsType2, true},
		{OpMod, OpIsType2, true},
		{OpMultiply, OpIsType2, true},
		{OpNand, OpIsType2, true},
		{OpNor, OpIsType2, true},
		{OpNot, OpIsType2, true},
		{OpObjectType, OpIsType2, true},
		{OpOr, OpIsType2, true},
		{OpPackage, OpIsType2, true},
		{OpVarPackage, OpIsType2, true},
		{OpRefOf, OpIsType2, true},
		{OpShiftLeft, OpIsType2, true},
		{OpShiftRight, OpIsType2, true},
		{OpSizeOf, OpIsType2, true},
		{OpStore, OpIsType2, true},
		{OpSubtract, OpIsType2, true},
		{OpTimer, OpIsType2, true},
		{OpToBCD, OpIsType2, true},
		{OpToBuffer, OpIsType2, true},
		{OpToDecimalString, OpIsType2, true},
		{OpToHexString, OpIsType2, true},
		{OpToInteger, OpIsType2, true},
		{OpToString, OpIsType2, true},
		{OpWait, OpIsType2, true},
		{OpXor, OpIsType2, true},
		{OpBytePrefix, OpIsType2, false},
		// OpIsDataObject
		{OpBytePrefix, OpIsDataObject, true},
		{OpWordPrefix, OpIsDataObject, true},
		{OpDwordPrefix, OpIsDataObject, true},
		{OpQwordPrefix, OpIsDataObject, true},
		{OpStringPrefix, OpIsDataObject, true},
		{OpZero, OpIsDataObject, true},
		{OpOne, OpIsDataObject, true},
		{OpOnes, OpIsDataObject, true},
		{OpRevision, OpIsDataObject, true},
		{OpBuffer, OpIsDataObject, true},
		{OpPackage, OpIsDataObject, true},
		{OpVarPackage, OpIsDataObject, true},
		{OpLor, OpIsDataObject, false},
		// OpIsBufferField
		{OpCreateField, OpIsBufferField, true},
		{OpCreateBitField, OpIsBufferField, true},
		{OpCreateByteField, OpIsBufferField, true},
		{OpCreateWordField, OpIsBufferField, true},
		{OpCreateDWordField, OpIsBufferField, true},
		{OpCreateQWordField, OpIsBufferField, true},
		{OpRevision, OpIsBufferField, false},
	}

	for specIndex, spec := range specs {
		if got := spec.testFn(spec.op); got != spec.want {
			t.Errorf("[spec %d] opcode %q: expected to get %t; got %t", specIndex, spec.op, spec.want, got)
		}
	}
}