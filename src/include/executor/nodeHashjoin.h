/*-------------------------------------------------------------------------
 *
 * nodeHashjoin.h
 *	  prototypes for nodeHashjoin.c
 *
 *
 * Portions Copyright (c) 1996-2008, PostgreSQL Global Development Group
 * Portions Copyright (c) 1994, Regents of the University of California
 *
 * $PostgreSQL: pgsql/src/include/executor/nodeHashjoin.h,v 1.37 2008/01/01 19:45:57 momjian Exp $
 *
 *-------------------------------------------------------------------------
 */
#ifndef NODEHASHJOIN_H
#define NODEHASHJOIN_H

#include "nodes/execnodes.h"
#include "storage/buffile.h"

struct HashJoinBatchSide;               /* #include "executor/hashjoin.h" */
struct Instrumentation;                 /* #include "executor/instrument.h" */

extern int	ExecCountSlotsHashJoin(HashJoin *node);
extern HashJoinState *ExecInitHashJoin(HashJoin *node, EState *estate, int eflags);
extern TupleTableSlot *ExecHashJoin(HashJoinState *node);
extern void ExecEndHashJoin(HashJoinState *node);
extern void ExecReScanHashJoin(HashJoinState *node, ExprContext *exprCtxt);

extern void ExecHashJoinSaveTuple(PlanState *ps, MemTuple tuple, uint32 hashvalue,
								  HashJoinTable hashtable, struct HashJoinBatchSide *batchside,
								  MemoryContext bfCxt);
extern void ExecEagerFreeHashJoin(HashJoinState *node);

#endif   /* NODEHASHJOIN_H */
